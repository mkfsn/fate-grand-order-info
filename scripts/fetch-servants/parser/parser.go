package parser

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseurl = "https://grandorder.wiki"
)

type Parser struct {
	reader io.Reader
}

func New(b []byte) *Parser {
	return &Parser{reader: bytes.NewBuffer(b)}
}

func (p *Parser) Parse() ([]*Servant, CommandMap, error) {
	doc, err := goquery.NewDocumentFromReader(p.reader)
	if err != nil {
		return nil, nil, err
	}

	commands := make(CommandMap)
	servants := make([]*Servant, 0)
	doc.Find("table.wikitable.sortable tbody tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		servant := p.parseServant(selection.Find("td"))
		servants = append(servants, servant)
		commands[servant.NoblePhantasm.Name] = servant.NoblePhantasm.Image
	})

	return servants, commands, nil
}

func (p *Parser) parseServant(td *goquery.Selection) *Servant {
	var servant Servant
	p.parseServantId(td, &servant)
	p.parseServantIcon(td, &servant)
	p.parseServantName(td, &servant)
	p.parseServantClass(td, &servant)
	p.parseServantRarity(td, &servant)
	p.parseServantAttack(td, &servant)
	p.parseServantHP(td, &servant)
	p.parseServantCommandCards(td, &servant)
	p.parseServantNoblePhantasm(td, &servant)
	return &servant
}

func (p *Parser) parseServantId(td *goquery.Selection, servant *Servant) {
	servant.ID = strings.Trim(td.Eq(0).Text(), "\n")
}

func (p *Parser) parseServantIcon(td *goquery.Selection, servant *Servant) {
	img := td.Eq(1).Find("img")
	src := img.AttrOr("src", "")
	servant.Icon.Src_1_0x = baseurl + src

	srcset := img.AttrOr("srcset", "")
	pairs := strings.Split(srcset, ", ")
	servant.Icon.Src_1_5x = baseurl + strings.Split(pairs[0], " ")[0]
	servant.Icon.Src_2_0x = baseurl + strings.Split(pairs[1], " ")[0]
}

func (p *Parser) parseServantName(td *goquery.Selection, servant *Servant) {
	servant.Name.EN = td.Eq(2).Find("a").Eq(0).AttrOr("title", "")
	servant.Name.JP = strings.Trim(td.Eq(2).Contents().Not("a").Text(), " \n")
}

func (p *Parser) parseServantClass(td *goquery.Selection, servant *Servant) {
	img := td.Eq(3).Find("img")
	src := img.AttrOr("src", "")
	servant.Class.Src_1_0x = baseurl + src

	srcset := img.AttrOr("srcset", "")
	pairs := strings.Split(srcset, ", ")
	servant.Class.Src_1_5x = baseurl + strings.Split(pairs[0], " ")[0]
	servant.Class.Src_2_0x = baseurl + strings.Split(pairs[1], " ")[0]

	alt := img.AttrOr("alt", "")
	servant.Class.Name = strings.Split(alt, " ")[2]
}

func (p *Parser) parseServantRarity(td *goquery.Selection, servant *Servant) {
	stars := td.Eq(4).Text()
	servant.Rarity, _ = strconv.Atoi(stars[:1])
}

func (p *Parser) parseServantAttack(td *goquery.Selection, servant *Servant) {
	servant.Attack.Min, _ = strconv.Atoi(strings.TrimSpace(td.Eq(5).Text()))
	servant.Attack.Max, _ = strconv.Atoi(strings.TrimSpace(td.Eq(6).Text()))
}

func (p *Parser) parseServantHP(td *goquery.Selection, servant *Servant) {
	servant.HP.Min, _ = strconv.Atoi(strings.TrimSpace(td.Eq(7).Text()))
	servant.HP.Max, _ = strconv.Atoi(strings.TrimSpace(td.Eq(8).Text()))
}

func (p *Parser) parseServantCommandCards(td *goquery.Selection, servant *Servant) {
	png := td.Eq(9).Find("img").AttrOr("alt", "")
	for i, command := range png[3:8] {
		servant.CommandCards[i].Name = string(command)
	}
}

func (p *Parser) parseServantNoblePhantasm(td *goquery.Selection, servant *Servant) {
	img := td.Eq(10).Find("img")
	src := img.AttrOr("src", "")
	servant.NoblePhantasm.Src_1_0x = baseurl + src

	srcset := img.AttrOr("srcset", "")
	pairs := strings.Split(srcset, ", ")
	servant.NoblePhantasm.Src_1_5x = baseurl + strings.Split(pairs[0], " ")[0]
	servant.NoblePhantasm.Src_2_0x = baseurl + strings.Split(pairs[1], " ")[0]

	alt := img.AttrOr("alt", "")
	servant.NoblePhantasm.Name = string(strings.Split(alt, " ")[2][0])
}
