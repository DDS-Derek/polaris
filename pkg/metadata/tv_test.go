package metadata

import (
	"polaris/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseTV1(t *testing.T) {
	s1 := "Twinkle Love 2024 S04 Complete 2160p WEB-DL HEVC AAC-QHstudIo"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 4)
	assert.Equal(t, m.IsSeasonPack, true)
	assert.Equal(t, m.Resolution, "2160p")
}

func Test_ParseTV2(t *testing.T) {
	s1 := "Cowboy Cartel S01E04 Photo Finish 1080p ATVP WEB-DL DDP5 1 Atmos H 264-FLUX [eztv] "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	assert.Equal(t, m.StartEpisode, 4)
	assert.Equal(t, m.IsSeasonPack, false)
	assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV3(t *testing.T) {
	s1 := "The.Bold.and.the.Beautiful.S37E219.XviD-AFG "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 37)
	assert.Equal(t, m.StartEpisode, 219)
	assert.Equal(t, m.IsSeasonPack, false)
	//assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV4(t *testing.T) {
	s1 := "Limitless Wrestling 2021 01 06 The Road Season 2 Episode 12 XviD-AFG [eztv] "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 2)
	assert.Equal(t, m.StartEpisode, 12)
	assert.Equal(t, m.IsSeasonPack, false)
	//assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV5(t *testing.T) {
	s1 := "[Breeze] One Punch Man S​01 S​02 [1080​p BD AV​1][dual audio]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	//assert.Equal(t, m.Episode, 219)
	assert.Equal(t, m.IsSeasonPack, true)
	//assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV6(t *testing.T) {
	s1 := "[千夏字幕组][小市民系列_Shoushimin Series][第03话][1080p_HEVC][简繁内封][招募新人]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	assert.Equal(t, m.StartEpisode, 3)
	assert.Equal(t, m.IsSeasonPack, false)
	assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV7(t *testing.T) {
	s1 := " [OPFans楓雪動漫][ONE PIECE 海賊王][第1113話][周日版][1080p][MP4][簡體]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	assert.Equal(t, m.StartEpisode, 1113)
	assert.Equal(t, m.IsSeasonPack, false)
	assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV8(t *testing.T) {
	s1 := "[桜都字幕组] 亦叶亦花 / Nanare Hananare [04][1080p][简体内嵌] "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	assert.Equal(t, m.StartEpisode, 4)
	assert.Equal(t, m.IsSeasonPack, false)
	assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV9(t *testing.T) {
	s1 := "[ANi] 戰國妖狐 千魔混沌篇 - 16 [1080P][Baha][WEB-DL][AAC AVC][CHT][MP4]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, m.Season, 1)
	assert.Equal(t, m.StartEpisode, 16)
	assert.Equal(t, m.IsSeasonPack, false)
	assert.Equal(t, m.Resolution, "1080p")
}

func Test_ParseTV10(t *testing.T) {
	s1 := " [桜都字幕组][一拳超人 第2季/One Punch Man 2nd Season][01-12 END][BIG5][720P]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 2, m.Season)
	//assert.Equal(t, 01, m.Episode)
	assert.Equal(t, true, m.IsSeasonPack)
	assert.Equal(t, "720p", m.Resolution)
}

func Test_ParseTV11(t *testing.T) {
	s1 := " [ANi] 這是妳與我的最後戰場，或是開創世界的聖戰 第二季 - 04 [1080P][Baha][WEB-DL][AAC AVC][CHT][MP4] "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 2, m.Season)
	assert.Equal(t, 4, m.StartEpisode)
	assert.Equal(t, false, m.IsSeasonPack)
	assert.Equal(t, "1080p", m.Resolution)
}

func Test_ParseTV12(t *testing.T) {
	s1 := " 牛仔Cowboy Cartel S02E04 Photo Finish 1080p ATVP WEB-DL DDP5 1 Atmos H 264-FLUX [eztv] "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 2, m.Season)
	assert.Equal(t, 4, m.StartEpisode)
	assert.Equal(t, false, m.IsSeasonPack)
	assert.Equal(t, "1080p", m.Resolution)
}

func Test_ParseTV13(t *testing.T) {
	s1 := "House of Dragon 2024 1080p S02E08 Leaked HQCAM NOT COMPLETE English Audio x264 ESub BOTHD"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 2, m.Season)
	assert.Equal(t, 8, m.StartEpisode)
	assert.Equal(t, false, m.IsSeasonPack)
	assert.Equal(t, "1080p", m.Resolution)
}

func Test_ParseTV14(t *testing.T) {
	s1 := "[GM-Team][国漫][斗破苍穹 第5季][Fights Break Sphere Ⅴ][2022][113][HEVC][GB][4K]"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 5, m.Season)
	assert.Equal(t, 113, m.StartEpisode)
	assert.Equal(t, false, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

//

func Test_ParseTV15(t *testing.T) {
	s1 := "[7³ACG] 剧场版 回转企鹅罐 [前篇] 你的列车是生存战略/RE:cycle of the PENGUINDRUM  Zenpen [简繁字幕] BDrip 1080p x265 FLAC"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	b := m.IsAcceptable("The Penguin")
	assert.False(t, b)
	//assert.Equal(t, 1, m.Season)
	//assert.Equal(t, 113, m.Episode)
	//assert.Equal(t, false, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

func Test_ParseTV16(t *testing.T) {
	s1 := "Romance in the Alley 2024 S01 E24-E25 1080p WEB-DL H.264 AAC-PTerWEB"
	m := ParseTv(s1)
	log.Infof("results: %+v", m)

	assert.Equal(t, 1, m.Season)
	assert.Equal(t, 24, m.StartEpisode)
	assert.Equal(t, 25, m.EndEpisode)
	//assert.Equal(t, false, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

func Test_ParseTV17(t *testing.T) {
	s1 := "小巷人家/Romance in the Alley 第24-25集 "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)

	assert.Equal(t, 1, m.Season)
	assert.Equal(t, 24, m.StartEpisode)
	assert.Equal(t, 25, m.EndEpisode)
	//assert.Equal(t, false, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

// Romance in the Alley 2024 S01E01-S01E21 2160p WEB-DL HEVC AAC-UBWEB
func Test_ParseTV18(t *testing.T) {
	s1 := "Romance in the Alley 2024 S01E01-S01E21 2160p WEB-DL HEVC AAC-UBWEB "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)

	assert.Equal(t, 1, m.Season)
	assert.Equal(t, 1, m.StartEpisode)
	assert.Equal(t, 21, m.EndEpisode)
	//assert.Equal(t, false, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

//The Count of Monte Cristo 2024 S01 1080p WEB-DL DD 5.1 H.264-playWEB
func Test_ParseTV20(t *testing.T) {
	s1 := "The Count of Monte Cristo 2024 S01 1080p WEB-DL DD 5.1 H.264-playWEB "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 1, m.Season)
	assert.Equal(t, true, m.IsSeasonPack)
}


func Test_ParseTV21(t *testing.T) {
	s1 := "【东京不够热】基督山伯爵-华丽的复仇-【01~09】【1280x720】【简中/日双语字幕】【2018春季日剧】【合集】 "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)
	assert.Equal(t, 1, m.Season)
	assert.Equal(t, 2018, m.Year)
	assert.Equal(t, true, m.IsSeasonPack)
}

// The Day of the Jackal (Season 1) WEB-DL 1080​p
func Test_ParseTV19(t *testing.T) {
	s1 := "The Day of the Jackal (Season 1) WEB-DL 1080​p "
	m := ParseTv(s1)
	log.Infof("results: %+v", m)

	assert.Equal(t, 1, m.Season)
	assert.Equal(t, true, m.IsSeasonPack)
	//assert.Equal(t, "720p", m.Resolution)
}

func Test_Name(t *testing.T) {
	m := Info{NameEn: "word кибердеревня новый год 2023 webrip 1080p"}
	b := m.IsAcceptable("word")
	assert.True(t, b)
}
