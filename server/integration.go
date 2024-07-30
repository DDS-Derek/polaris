package server

import (
	"bytes"
	"fmt"
	"path/filepath"
	"polaris/db"
	"polaris/ent/media"
	"polaris/log"

	"github.com/pkg/errors"
)

func (s *Server) writePlexmatch(seriesId int, episodeNum int, targetDir, name string) error {

	if !s.plexmatchEnabled() {
		return nil
	}
	series, err := s.db.GetMedia(seriesId)
	if err != nil {
		return err
	}
	if series.MediaType != media.MediaTypeTv {
		return nil
	}
	st, err := s.getStorage(series.StorageID, media.MediaTypeTv)
	if err != nil {
		return errors.Wrap(err, "get storage")
	}

	_, err = st.ReadFile(filepath.Join(series.TargetDir, ".plexmatch"))
	if err != nil { 
		//create new
		log.Warnf(".plexmatch file not found, create new one: %s", series.NameEn)
		return st.WriteFile(filepath.Join(series.TargetDir, ".plexmatch"), []byte(fmt.Sprintf("tmdbid: %d\n",series.TmdbID)))
	} 

	if episodeNum == 0 {
		return nil
	}
	buff := bytes.Buffer{}
	seasonPlex := filepath.Join(targetDir, ".plexmatch")
	data, err := st.ReadFile(seasonPlex)
	if err != nil {
		log.Infof("read season plexmatch: %v", err)
	} else {
		buff.Write(data)
	}
	buff.WriteString(fmt.Sprintf("ep: %d: %s\n", episodeNum, name))
	log.Infof("write season plexmatch file content: %s", buff.String())
	return st.WriteFile(seasonPlex, buff.Bytes())
}

func (s *Server) plexmatchEnabled() bool {
	return s.db.GetSetting(db.SettingPlexMatchEnabled) == "true"
}
