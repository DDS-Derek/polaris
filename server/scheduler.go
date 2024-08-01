package server

import (
	"fmt"
	"path/filepath"
	"polaris/ent"
	"polaris/ent/episode"
	"polaris/ent/history"
	"polaris/ent/media"
	"polaris/log"
	"polaris/pkg"
	"polaris/pkg/notifier/message"
	"polaris/pkg/utils"
	"polaris/server/core"
	"time"

	"github.com/pkg/errors"
)

func (s *Server) scheduler() {
	s.mustAddCron("@every 1m", s.checkTasks)
	s.mustAddCron("0 0 * * * *", func() {
		s.downloadTvSeries()
		s.downloadMovie()
	})
	s.mustAddCron("0 0 */12 * * *", s.checkAllSeriesNewSeason)
	s.cron.Start()
}

func (s *Server) mustAddCron(spec string, cmd func()) {
	if err := s.cron.AddFunc(spec, cmd); err != nil {
		log.Errorf("add func error: %v", err)
		panic(err)
	}
}

func (s *Server) checkTasks() {
	log.Debug("begin check tasks...")
	for id, t := range s.tasks {
		if !t.Exists() {
			log.Infof("task no longer exists: %v", id)
			delete(s.tasks, id)
			continue
		}
		log.Infof("task (%s) percentage done: %d%%", t.Name(), t.Progress())
		if t.Progress() == 100 {
			r := s.db.GetHistory(id)
			if r.Status == history.StatusSuccess {
				//task already success, check seed ratio
				torrent := s.tasks[id]
				ok, err := s.isSeedRatioLimitReached(r.IndexerID, torrent)
				if err != nil {
					log.Warnf("getting torrent seed ratio : %v", err)
					ok = false
				}
				if ok {
					log.Infof("torrent file seed ratio reached, remove: %v", torrent.Name())
					torrent.Remove()
					delete(s.tasks, id)
				} else {
					log.Infof("torrent file still sedding: %v", torrent.Name())
				}
				continue
			}
			log.Infof("task is done: %v", t.Name())
			s.sendMsg(fmt.Sprintf(message.DownloadComplete, t.Name()))
			go func() {
				if err := s.moveCompletedTask(id); err != nil {
					log.Infof("post tasks for id %v fail: %v", id, err)
				}
			}()
		}
	}
}

func (s *Server) moveCompletedTask(id int) (err1 error) {
	torrent := s.tasks[id]
	r := s.db.GetHistory(id)
	if r.Status == history.StatusUploading {
		log.Infof("task %d is already uploading, skip", id)
		return nil
	}
	s.db.SetHistoryStatus(r.ID, history.StatusUploading)
	seasonNum, err := utils.SeasonId(r.TargetDir)
	if err != nil {
		log.Errorf("no season id: %v", r.TargetDir)
		seasonNum = -1
	}
	downloadclient, err := s.db.GetDownloadClient(r.DownloadClientID)
	if err != nil {
		log.Errorf("get task download client error: %v, use default one", err)
		downloadclient = &ent.DownloadClients{RemoveCompletedDownloads: true, RemoveFailedDownloads: true}
	}
	torrentName := torrent.Name()

	defer func() {

		if err1 != nil {
			s.db.SetHistoryStatus(r.ID, history.StatusFail)
			if r.EpisodeID != 0 {
				s.db.SetEpisodeStatus(r.EpisodeID, episode.StatusMissing)
			} else {
				s.db.SetSeasonAllEpisodeStatus(r.MediaID, seasonNum, episode.StatusMissing)
			}
			s.sendMsg(fmt.Sprintf(message.ProcessingFailed, err))
			if downloadclient.RemoveFailedDownloads {
				log.Debugf("task failed, remove failed torrent and files related")
				delete(s.tasks, r.ID)
				torrent.Remove()
			}
		}
	}()

	series := s.db.GetMediaDetails(r.MediaID)
	if series == nil {
		return nil
	}
	st := s.db.GetStorage(series.StorageID)
	log.Infof("move task files to target dir: %v", r.TargetDir)
	stImpl, err := s.getStorage(st.ID, series.MediaType)
	if err != nil {
		return err
	}

	//如果种子是路径，则会把路径展开，只移动文件，类似 move dir/* dir2/, 如果种子是文件，则会直接移动文件，类似 move file dir/
	if err := stImpl.Copy(filepath.Join(s.db.GetDownloadDir(), torrentName), r.TargetDir); err != nil {
		return errors.Wrap(err, "move file")
	}

	// .plexmatch file
	if err := s.writePlexmatch(r.MediaID, r.EpisodeID, r.TargetDir, torrentName); err != nil {
		log.Errorf("create .plexmatch file error: %v", err)
	}

	s.db.SetHistoryStatus(r.ID, history.StatusSuccess)
	if r.EpisodeID != 0 {
		s.db.SetEpisodeStatus(r.EpisodeID, episode.StatusDownloaded)
	} else {
		s.db.SetSeasonAllEpisodeStatus(r.MediaID, seasonNum, episode.StatusDownloaded)
	}
	s.sendMsg(fmt.Sprintf(message.ProcessingComplete, torrentName))

	//判断是否需要删除本地文件
	ok, err := s.isSeedRatioLimitReached(r.IndexerID, torrent) 
	if err != nil {
		log.Warnf("getting torrent seed ratio %s: %v", torrent.Name(), err)
		ok = false
	}
	if downloadclient.RemoveCompletedDownloads && ok {
		log.Debugf("download complete,remove torrent and files related")
		delete(s.tasks, r.ID)
		torrent.Remove()
	}

	log.Infof("move downloaded files to target dir success, file: %v, target dir: %v", torrentName, r.TargetDir)
	return nil
}

func (s *Server) checkDownloadedSeriesFiles(m *ent.Media) error {
	if m.MediaType != media.MediaTypeTv {
		return nil
	}
	log.Infof("check files in directory: %s", m.TargetDir)

	var storageImpl, err = s.getStorage(m.StorageID, media.MediaTypeTv)
	if err != nil {
		return err
	}

	files, err := storageImpl.ReadDir(m.TargetDir)
	if err != nil {
		return errors.Wrapf(err, "read dir %s", m.TargetDir)
	}

	for _, in := range files {
		if !in.IsDir() { //season dir, ignore file
			continue
		}
		dir := filepath.Join(m.TargetDir, in.Name())
		epFiles, err := storageImpl.ReadDir(dir)
		if err != nil {
			log.Errorf("read dir %s error: %v", dir, err)
			continue
		}
		for _, ep := range epFiles {
			log.Infof("found file: %v", ep.Name())
			seNum, epNum, err := utils.FindSeasonEpisodeNum(ep.Name())
			if err != nil {
				log.Errorf("find season episode num error: %v", err)
				continue
			}
			log.Infof("found match, season num %d, episode num %d", seNum, epNum)
			ep, err := s.db.GetEpisode(m.ID, seNum, epNum)
			if err != nil {
				log.Error("update episode: %v", err)
				continue
			}
			err = s.db.SetEpisodeStatus(ep.ID, episode.StatusDownloaded)
			if err != nil {
				log.Error("update episode: %v", err)
				continue
			}

		}
	}
	return nil

}

type Task struct {
	//Processing bool
	pkg.Torrent
}

func (s *Server) downloadTvSeries() {
	log.Infof("begin check all tv series resources")
	allSeries := s.db.GetMediaWatchlist(media.MediaTypeTv)
	for _, series := range allSeries {
		tvDetail := s.db.GetMediaDetails(series.ID)
		for _, ep := range tvDetail.Episodes {
			if !series.DownloadHistoryEpisodes { //设置不下载历史已播出剧集，只下载将来剧集
				t, err := time.Parse("2006-01-02", ep.AirDate)
				if err != nil {
					log.Error("air date not known, skip: %v", ep.Title)
					continue
				}
				if series.CreatedAt.Sub(t) > 24*time.Hour { //剧集在加入watchlist之前，不去下载
					continue
				}
			}

			if ep.Status != episode.StatusMissing { //已经下载的不去下载
				continue
			}
			name, err := s.searchAndDownload(series.ID, ep.SeasonNumber, ep.EpisodeNumber)
			if err != nil {
				log.Infof("cannot find resource to download for %s: %v", ep.Title, err)
			} else {
				log.Infof("begin download torrent resource: %v", name)
			}

		}

	}
}

func (s *Server) downloadMovie() {
	log.Infof("begin check all movie resources")
	allSeries := s.db.GetMediaWatchlist(media.MediaTypeMovie)

	for _, series := range allSeries {
		detail := s.db.GetMediaDetails(series.ID)
		if len(detail.Episodes) == 0 {
			log.Errorf("no related dummy episode: %v", detail.NameEn)
			continue
		}
		ep := detail.Episodes[0]
		if ep.Status == episode.StatusDownloaded {
			continue
		}

		if err := s.downloadMovieSingleEpisode(ep); err != nil {
			log.Errorf("download movie error: %v", err)
		}
	}
}

func (s *Server) downloadMovieSingleEpisode(ep *ent.Episode) error {
	trc, dlc, err := s.getDownloadClient()
	if err != nil {
		return errors.Wrap(err, "connect transmission")
	}

	res, err := core.SearchMovie(s.db, ep.MediaID, true)
	if err != nil {

		return errors.Wrap(err, "search movie")
	}
	r1 := res[0]
	log.Infof("begin download torrent resource: %v", r1.Name)
	torrent, err := trc.Download(r1.Link, s.db.GetDownloadDir())
	if err != nil {
		return errors.Wrap(err, "downloading")
	}
	torrent.Start()

	history, err := s.db.SaveHistoryRecord(ent.History{
		MediaID:          ep.MediaID,
		EpisodeID:        ep.ID,
		SourceTitle:      r1.Name,
		TargetDir:        "./",
		Status:           history.StatusRunning,
		Size:             r1.Size,
		Saved:            torrent.Save(),
		DownloadClientID: dlc.ID,
	})
	if err != nil {
		log.Errorf("save history error: %v", err)
	}

	s.tasks[history.ID] = &Task{Torrent: torrent}

	s.db.SetEpisodeStatus(ep.ID, episode.StatusDownloading)
	return nil
}

func (s *Server) checkAllSeriesNewSeason() {
	log.Infof("begin checking series all new season")
	allSeries := s.db.GetMediaWatchlist(media.MediaTypeTv)
	for _, series := range allSeries {
		err := s.checkSeiesNewSeason(series)
		if err != nil {
			log.Errorf("check series new season error: series name %v, error: %v", series.NameEn, err)
		}
	}
}

func (s *Server) checkSeiesNewSeason(media *ent.Media) error {
	d, err := s.MustTMDB().GetTvDetails(media.TmdbID, s.language)
	if err != nil {
		return errors.Wrap(err, "tmdb")
	}
	lastsSason := d.NumberOfSeasons
	seasonDetail, err := s.MustTMDB().GetSeasonDetails(media.TmdbID, lastsSason, s.language)
	if err != nil {
		return errors.Wrap(err, "tmdb season")
	}

	for _, ep := range seasonDetail.Episodes {
		epDb, err := s.db.GetEpisode(media.ID, ep.SeasonNumber, ep.EpisodeNumber)
		if err != nil {
			if ent.IsNotFound(err) {
				log.Infof("add new episode: %+v", ep)
				episode := &ent.Episode{
					MediaID:       media.ID,
					SeasonNumber:  ep.SeasonNumber,
					EpisodeNumber: ep.EpisodeNumber,
					Title:         ep.Name,
					Overview:      ep.Overview,
					AirDate:       ep.AirDate,
					Status:        episode.StatusMissing,
				}
				s.db.SaveEposideDetail2(episode)
			}
		} else { //update episode
			if ep.Name != epDb.Title || ep.Overview != epDb.Overview || ep.AirDate != epDb.AirDate {
				log.Infof("update new episode: %+v", ep)
				s.db.UpdateEpiode2(epDb.ID, ep.Name, ep.Overview, ep.AirDate)
			}
		}
	}
	return nil
}

func (s *Server) isSeedRatioLimitReached(indexId int, t pkg.Torrent) (bool, error) {
	indexer, err := s.db.GetIndexer(indexId)
	if err != nil {
		return false, err
	}
	currentRatio := t.SeedRatio()
	if currentRatio == nil {
		return false, errors.New("seed ratio not exist")
	}
	return *currentRatio >= float64(indexer.SeedRatio), nil
}
