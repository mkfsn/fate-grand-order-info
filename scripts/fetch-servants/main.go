package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type IntegerRange struct {
	Min int
	Max int
}

type CommandCard = string

type Servant struct {
	ID   string
	Icon struct {
		Src string
	}
	Name struct {
		EN string
		JP string
	}
	Class         string
	Rarity        int
	Attack        IntegerRange
	HP            IntegerRange
	CommandCards  [5]CommandCard
	NoblePhantasm CommandCard
}

func main() {
	b, err := fetch()
	if err != nil {
		log.Fatalln(err)
	}

	if err := parse(b); err != nil {
		log.Fatalln(err)
	}
}

func fetch() ([]byte, error) {
	url := "https://grandorder.wiki/Servant_List"
	resp, err := http.Get(url)
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

func parse(b []byte) error {
	buf := bytes.NewBuffer(b)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	servants := make([]Servant, 0)
	doc.Find("table.wikitable.sortable tbody tr").Each(func(i int, selection *goquery.Selection) {
		var servant Servant
		// ID
		id := strings.Trim(selection.Find("td").Eq(0).Text(), "\n")
		if id == "" {
			return
		}
		servant.ID = id
		// Icon
		src := selection.Find("td").Eq(1).Find("img").AttrOr("src", "")
		items := strings.Split(src, "/")
		filename := items[len(items)-1]
		downloadImage(&wg, "https://grandorder.wiki"+src, filename)
		servant.Icon.Src = "/images/" + filename
		// Name
		servant.Name.EN = selection.Find("td").Eq(2).Find("a").Eq(0).AttrOr("title", "")
		servant.Name.JP = strings.Trim(selection.Find("td").Eq(2).Contents().Not("a").Text(), "\n")
		// Class
		png := selection.Find("td").Eq(3).Find("img").AttrOr("alt", "")
		servant.Class = strings.Split(png, " ")[2]
		// Rarity
		stars := selection.Find("td").Eq(4).Text()
		servant.Rarity, _ = strconv.Atoi(stars[:1])
		// Attack
		servant.Attack.Min, _ = strconv.Atoi(strings.TrimSpace(selection.Find("td").Eq(5).Text()))
		servant.Attack.Max, _ = strconv.Atoi(strings.TrimSpace(selection.Find("td").Eq(6).Text()))
		// HP
		servant.HP.Min, _ = strconv.Atoi(strings.TrimSpace(selection.Find("td").Eq(7).Text()))
		servant.HP.Max, _ = strconv.Atoi(strings.TrimSpace(selection.Find("td").Eq(8).Text()))
		// Command Cards
		png = selection.Find("td").Eq(9).Find("img").AttrOr("alt", "")
		for i, command := range png[3:8] {
			servant.CommandCards[i] = string(command)
		}

		// Noble Phantasm
		png = selection.Find("td").Eq(10).Find("img").AttrOr("alt", "")
		servant.NoblePhantasm = string(strings.Split(png, " ")[2][0])

		servants = append(servants, servant)
	})
	wg.Wait()

	b, _ = json.MarshalIndent(servants, "", "\t")
	fmt.Printf("%s\n", b)

	file, err := os.Create(fmt.Sprintf("../../static/data/servants.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return nil
}

func downloadImage(wg *sync.WaitGroup, url, filename string) {
	wg.Add(1)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		defer resp.Body.Close()

		file, err := os.Create(fmt.Sprintf("../../static/images/%s", filename))
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Printf("error: %s\n", err)
			return
		}
		wg.Done()
	}()
}
