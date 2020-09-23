package dumper

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"fetch-servant/processer"
)

type Dumper struct {
	results *processer.Results
}

func New(results *processer.Results) *Dumper {
	return &Dumper{results: results}
}

func (d *Dumper) Dump() error {
	var wg sync.WaitGroup
	for _, image := range d.results.Images {
		d.downloadImage(&wg, image.Url, image.Filepath)
		// TODO: download class icons
		// TODO: download command card icons
	}
	for _, data := range d.results.Data {
		if err := d.writeToFile(data.Filepath, bytes.NewBuffer(data.Text)); err != nil {
			log.Printf("error: %s\n", err)
		}
	}
	wg.Wait()
	return nil
}

func (d *Dumper) downloadImage(wg *sync.WaitGroup, url, filepath string) {
	wg.Add(1)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if err := d.writeToFile(filepath, resp.Body); err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		wg.Done()
	}()
}

func (d *Dumper) writeToFile(filepath string, r io.Reader) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, r)
	if err != nil {
		return err
	}
	return nil
}
