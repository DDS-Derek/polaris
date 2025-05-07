package qbittorrent

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"polaris/pkg"
	"polaris/pkg/go-qbittorrent/qbt"
	"polaris/pkg/utils"
	"strings"

	"github.com/pkg/errors"
)

const btCategory = "polaris"

type Info struct {
	URL      string
	User     string
	Password string
}

type Client struct {
	c        *qbt.Client
	category string
	Info
}

func NewClient(url, user, pass string) (*Client, error) {
	// connect to qbittorrent client
	qb := qbt.NewClient(url)

	// login to the client
	loginOpts := qbt.LoginOptions{
		Username: user,
		Password: pass,
	}
	err := qb.Login(loginOpts)
	if err != nil {
		return nil, err
	}

	return &Client{c: qb, category: btCategory, Info: Info{URL: url, User: user, Password: pass}}, nil
}

func (c *Client) GetAll() ([]pkg.Torrent, error) {
	tt, err := c.c.Torrents(qbt.TorrentsOptions{Category: &c.category})
	if err != nil {
		return nil, errors.Wrap(err, "get torrents")
	}
	var res []pkg.Torrent
	for _, t := range tt {
		t1 := &Torrent{
			c:    c.c,
			hash: t.Hash,
			//Info: c.Info,
		}
		res = append(res, t1)
	}
	return res, nil
}

func (c *Client) GetListenPort() (int, error) {
	pref, err := c.c.Preferences()
	if err != nil {
		return 0, errors.Wrap(err, "get preferences")
	}
	
	return pref.ListenPort, nil
}

func (c *Client) SetListenPort(port int) error {
	ok, err := c.c.SetPreferences(map[string]any{"listen_port": port})
	if !ok || err != nil {
		return errors.Wrap(err, "set preferences")
	}
	return nil
}

func (c *Client) Download(link, hash, dir string) (pkg.Torrent, error) {
	err := c.c.DownloadLinks([]string{link}, qbt.DownloadOptions{Savepath: &dir, Category: &c.category})
	if err != nil {
		return nil, errors.Wrap(err, "qbt download")
	}
	return &Torrent{hash: hash, c: c.c}, nil

}

func NewTorrentHash(info Info, hash string) (*Torrent, error) {
	c, err := NewClient(info.URL, info.User, info.Password)
	if err != nil {
		return nil, err
	}

	t := &Torrent{
		c:    c.c,
		hash: hash,
	}
	if !t.Exists() {
		return nil, errors.Errorf("torrent not exist: %v", hash)
	}
	return t, nil
}

func NewTorrent(info Info, link string) (*Torrent, error) {
	magnet, err := utils.Link2Magnet(link)
	if err != nil {
		return nil, errors.Errorf("converting link to magnet error, link: %v, error: %v", link, err)
	}

	hash, err := utils.MagnetHash(magnet)
	if err != nil {
		return nil, err
	}

	return NewTorrentHash(info, hash)
}

type Torrent struct {
	c    *qbt.Client
	hash string
	//info Info
}

func (t *Torrent) GetHash() string {
	return t.hash
}

func (t *Torrent) getTorrent() (*qbt.TorrentInfo, error) {
	all, err := t.c.Torrents(qbt.TorrentsOptions{Hashes: []string{t.hash}})
	if err != nil {
		return nil, err
	}
	if len(all) == 0 {
		return nil, fmt.Errorf("no such torrent: %v", t.hash)
	}
	return &all[0], nil
}

func (t *Torrent) Name() (string, error) {
	dir, err := t.getTorrentBaseNameOrDir()
	if err != nil { //use torrent name
		qb, err := t.getTorrent()
		if err != nil {
			return "", err
		}
		return qb.Name, nil
	}
	return dir, nil
}

// https://github.com/qbittorrent/qBittorrent/issues/13572
func (t *Torrent) getTorrentBaseNameOrDir() (string, error) {
	files, err := t.c.TorrentFiles(t.hash)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", errors.Wrap(err, "no file")
	}
	name := files[0].Name
	dir := strings.Split(name, string(os.PathSeparator))[0]
	return dir, nil
}

func (t *Torrent) Progress() (int, error) {
	qb, err := t.getTorrent()
	if err != nil {
		return 0, err
	}
	p := qb.Progress * 100
	if p >= 100 {
		return 100, nil
	}
	if int(p) == 100 {
		return 99, nil
	}

	return int(p), nil
}

func (t *Torrent) Stop() error {
	return t.c.Pause([]string{t.hash})
}

func (t *Torrent) Start() error {
	ok, err := t.c.Resume([]string{t.hash})
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("status not 200")
	}
	return nil
}

func (t *Torrent) Remove() error {
	ok, err := t.c.Delete([]string{t.hash}, true)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("status not 200")
	}
	return nil
}

func (t *Torrent) Exists() bool {
	_, err := t.getTorrent()
	return err == nil
}

func (t *Torrent) SeedRatio() (float64, error) {
	qb, err := t.getTorrent()
	if err != nil {
		return 0, err
	}
	return qb.Ratio, nil
}

func (t *Torrent) Walk(f func(string) error) error {
	files, err := t.c.TorrentFiles(t.hash)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := f(file.Name); err != nil {
			return errors.Errorf("proccess file (%s) error: %v", file.Name, err)
		}
	}
	return nil
}

func (t *Torrent) WalkFunc() func(fn func(path string, info fs.FileInfo) error) error {
	files, err := t.c.TorrentFiles(t.hash)
	if err != nil {
		return func(fn func(path string, info fs.FileInfo) error) error {
			return err
		}
	}
	path, err := t.c.DefaultSavePath()
	if err != nil {
		return func(fn func(path string, info fs.FileInfo) error) error {
			return err
		}
	}
	
	return func(fn func(path string, info fs.FileInfo) error) error {
		for _, file := range files {
			name := filepath.Join(path, file.Name)
			info, err := os.Stat(name)
			if err != nil {
				return err
			}

			if err := fn(name, info); err != nil {
				return errors.Errorf("proccess file (%s) error: %v", file.Name, err)
			}
		}
		return nil

	}

}
