package torznab

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"polaris/ent"
	"polaris/log"
	"slices"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Response struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Torznab string   `xml:"torznab,attr"`
	Channel struct {
		Text string `xml:",chardata"`
		Link struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Category    string `xml:"category"`
		Item        []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Text           string `xml:",chardata"`
	Title          string `xml:"title"`
	Guid           string `xml:"guid"`
	Jackettindexer struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"jackettindexer"`
	Type        string   `xml:"type"`
	Comments    string   `xml:"comments"`
	PubDate     string   `xml:"pubDate"`
	Size        string   `xml:"size"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	Category    []string `xml:"category"`
	Enclosure   struct {
		Text   string `xml:",chardata"`
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure"`
	Attr []struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name,attr"`
		Value string `xml:"value,attr"`
	} `xml:"attr"`
}

func (i *Item) GetAttr(key string) string {
	for _, a := range i.Attr {
		if a.Name == key {
			return a.Value
		}
	}
	return ""
}
func (r *Response) ToResults(indexer *ent.Indexers) []Result {
	var res []Result
	for _, item := range r.Channel.Item {
		if slices.Contains(item.Category, "3000") { //exclude audio files
			continue
		}
		// link, err := utils.Link2Magnet(item.Link) //TODO time consuming operation
		// if err != nil {
		// 	log.Warnf("converting link to magnet error, error: %v, link: %v", err, item.Link)
		// 	continue
		// }
		imdb := ""
		if item.GetAttr("imdbid") != "" {
			imdb = item.GetAttr("imdbid")
		} else if item.GetAttr("imdb") != "" {
			imdb = item.GetAttr("imdb")
		}
		r := Result{
			Name:                 item.Title,
			Description:          item.Description,
			Link:                 item.Link,
			Size:                 mustAtoI(item.Size),
			Seeders:              int(mustAtoI(item.GetAttr("seeders"))),
			Peers:                int(mustAtoI(item.GetAttr("peers"))),
			Category:             int(mustAtoI(item.GetAttr("category"))),
			ImdbId:               imdb,
			DownloadVolumeFactor: tryParseFloat(item.GetAttr("downloadvolumefactor")),
			UploadVolumeFactor:   tryParseFloat(item.GetAttr("uploadvolumefactor")),
			Source:               indexer.Name,
			IndexerId:            indexer.ID,
			Priority:             indexer.Priority,
			IsPrivate:            item.Type == "private",
		}
		res = append(res, r)
	}
	return res
}

func mustAtoI(key string) int64 {
	i, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Errorf("must atoi error: %v", err)
		panic(err)
	}
	return i
}

func tryParseFloat(s string) float32 {
	r, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Warnf("parse float error: %v", err)
		return 0
	}
	return float32(r)
}

func Search(indexer *ent.Indexers, keyWord string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, indexer.URL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}
	var q = url.Values{}
	q.Add("apikey", indexer.APIKey)
	q.Add("t", "search")
	q.Add("q", keyWord)
	req.URL.RawQuery = q.Encode()
	key := fmt.Sprintf("%s: %s", indexer.Name, keyWord)

	cacheRes, ok := cc.Get(key)
	if !ok {
		log.Debugf("not found in cache, need query again: %v", key)
		res, err := doRequest(req)
		if err != nil {
			//cc.Set(key, nil)
			return nil, errors.Wrap(err, "do http request")
		}
		cacheRes = res.ToResults(indexer)
		cc.Set(key, cacheRes)
	} else {
		log.Debugf("found cache match for key: %v", key)
	}
	return cacheRes, nil
}

func doRequest(req *http.Request) (*Response, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do http")
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read http body")
	}
	var res Response
	err = xml.Unmarshal(data, &res)
	if err != nil {
		return nil, errors.Wrapf(err, "xml unmarshal data: %v", string(data))
	}
	return &res, nil
}

type Result struct {
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	Link                 string  `json:"link"`
	Size                 int64   `json:"size"`
	Seeders              int     `json:"seeders"`
	Peers                int     `json:"peers"`
	Category             int     `json:"category"`
	Source               string  `json:"source"`
	DownloadVolumeFactor float32 `json:"download_volume_factor"`
	UploadVolumeFactor   float32 `json:"upload_volume_factor"`
	IndexerId            int     `json:"indexer_id"`
	Priority             int     `json:"priority"`
	IsPrivate            bool    `json:"is_private"`
	ImdbId               string  `json:"imdb_id"`
}
