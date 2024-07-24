package core

import (
	"fmt"
	"polaris/db"
	"polaris/log"
	"polaris/pkg/metadata"
	"polaris/pkg/torznab"
	"polaris/pkg/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func SearchSeasonPackage(db1 *db.Client, seriesId, seasonNum int, checkResolution bool) ([]torznab.Result, error) {
	return SearchEpisode(db1, seriesId, seasonNum, -1, checkResolution)
}

func SearchEpisode(db1 *db.Client, seriesId, seasonNum, episodeNum int, checkResolution bool) ([]torznab.Result, error) {
	series := db1.GetMediaDetails(seriesId)
	if series == nil {
		return nil, fmt.Errorf("no tv series of id %v", seriesId)
	}

	res := searchWithTorznab(db1, series.NameEn)
	resCn := searchWithTorznab(db1, series.NameCn)
	res = append(res, resCn...)

	var filtered []torznab.Result
	for _, r := range res {
		meta := metadata.ParseTv(r.Name)
		if meta.Season != seasonNum {
			continue
		}
		if episodeNum != -1 && meta.Episode != episodeNum { //not season pack
			continue
		}
		if checkResolution && meta.Resolution != series.Resolution.String() {
			continue
		}
		if !utils.IsNameAcceptable(meta.NameEn, series.NameEn) && !utils.IsNameAcceptable(meta.NameCn, series.NameCn) {
			continue
		}
		filtered = append(filtered, r)
	}
	if len(filtered) == 0 {
		return nil, errors.New("no resource found")
	}

	return filtered, nil

}

func SearchMovie(db1 *db.Client, movieId int, checkResolution bool) ([]torznab.Result, error) {
	movieDetail := db1.GetMediaDetails(movieId)
	if movieDetail == nil {
		return nil, errors.New("no media found of id")
	}

	res := searchWithTorznab(db1, movieDetail.NameEn)

	res1 := searchWithTorznab(db1, movieDetail.NameCn)
	res = append(res, res1...)

	if len(res) == 0 {
		return nil, fmt.Errorf("no resource found")
	}
	var filtered []torznab.Result
	for _, r := range res {
		meta := metadata.ParseMovie(r.Name)
		if !utils.IsNameAcceptable(meta.NameEn, movieDetail.NameEn) {
			continue
		}
		if checkResolution && meta.Resolution != movieDetail.Resolution.String() {
			continue
		}
		ss := strings.Split(movieDetail.AirDate, "-")[0]
		year, _ := strconv.Atoi(ss)
		if meta.Year != year && meta.Year != year-1 && meta.Year != year+1 { //year not match
			continue
		}

		filtered = append(filtered, r)

	}
	if len(filtered) == 0 {
		return nil, errors.New("no resource found")
	}

	return filtered, nil

}

func searchWithTorznab(db *db.Client, q string) []torznab.Result {

	var res []torznab.Result
	allTorznab := db.GetAllTorznabInfo()
	for _, tor := range allTorznab {
		resp, err := torznab.Search(tor.URL, tor.ApiKey, q)
		if err != nil {
			log.Errorf("search %s error: %v", tor.Name, err)
			continue
		}
		res = append(res, resp...)
	}
	sort.Slice(res, func(i, j int) bool {
		var s1 = res[i]
		var s2 = res[j]
		return s1.Seeders > s2.Seeders
	})

	return res
}
