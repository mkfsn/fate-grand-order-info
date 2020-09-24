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

type job struct {
	wg       *sync.WaitGroup
	url      string
	filepath string
}

type Dumper struct {
	results *processer.Results
	jobs    chan job
}

func New(results *processer.Results, setters ...Option) *Dumper {
	var opts options
	for _, setter := range setters {
		setter(&opts)
	}
	for _, image := range opts.imageURLs {
		results.Images = append(results.Images, processer.NewImageFromSrc(image))
	}
	d := &Dumper{
		results: results,
		jobs:    make(chan job, 1024),
	}
	for i := 0; i < 128; i++ {
		go d.downloader()
	}
	return d
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
	d.jobs <- job{
		wg:       wg,
		url:      url,
		filepath: filepath,
	}
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

func (d *Dumper) downloader() {

	handler := func(wg *sync.WaitGroup, url, filepath string) {
		defer wg.Done()

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		defer resp.Body.Close()

		if err := d.writeToFile(filepath, resp.Body); err != nil {
			log.Printf("error: %s\n", err)
		}
	}

	for {
		select {
		case job := <-d.jobs:
			handler(job.wg, job.url, job.filepath)
		}
	}
}
