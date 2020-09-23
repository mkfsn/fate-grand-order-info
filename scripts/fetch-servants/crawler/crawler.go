package crawler

import (
	"io/ioutil"
	"net/http"
)

type Crawler struct {
	url string
}

func New(url string) *Crawler {
	return &Crawler{url: url}
}

func (c *Crawler) Crawl() ([]byte, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
