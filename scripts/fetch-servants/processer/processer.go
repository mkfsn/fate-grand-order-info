package processer

import (
	"encoding/json"
	"strings"

	"fetch-servant/parser"
)

type Data struct {
	Text     []byte
	Filepath string
}

type Image struct {
	Url      string
	Filepath string
}

func NewImageFromSrc(src string) *Image {
	return &Image{
		Url:      src,
		Filepath: "../../static/images/" + getFilename(src),
	}
}

type Results struct {
	Data   []*Data
	Images []*Image
}

type Processer struct {
	servants []*parser.Servant
	commands parser.CommandMap
}

func New(servants []*parser.Servant, commands parser.CommandMap) *Processer {
	return &Processer{servants: servants, commands: commands}
}

func (p *Processer) Process() (*Results, error) {
	var results Results

	iconSet := make(map[string]bool)

	var servants []*Servant
	for _, servant := range p.servants {
		results.Images = append(results.Images, NewImageFromSrc(servant.Icon.Src_1_0x))
		results.Images = append(results.Images, NewImageFromSrc(servant.Icon.Src_1_5x))
		results.Images = append(results.Images, NewImageFromSrc(servant.Icon.Src_2_0x))
		iconSet[servant.Class.Src_1_0x] = true
		iconSet[servant.Class.Src_1_5x] = true
		iconSet[servant.Class.Src_2_0x] = true
		iconSet[servant.NoblePhantasm.Src_1_0x] = true
		iconSet[servant.NoblePhantasm.Src_1_5x] = true
		iconSet[servant.NoblePhantasm.Src_2_0x] = true

		var s Servant
		s.ID = servant.ID
		s.Icon = picture{
			Src_1_0x: "/images/" + getFilename(servant.Icon.Src_1_0x),
			Src_1_5x: "/images/" + getFilename(servant.Icon.Src_1_5x),
			Src_2_0x: "/images/" + getFilename(servant.Icon.Src_2_0x),
			Title:    servant.Name.JP,
		}
		s.Name = i18nName{
			EN: servant.Name.EN,
			JP: servant.Name.JP,
		}
		s.Class = picture{
			Src_1_0x: "/images/" + getFilename(servant.Class.Src_1_0x),
			Src_1_5x: "/images/" + getFilename(servant.Class.Src_1_5x),
			Src_2_0x: "/images/" + getFilename(servant.Class.Src_2_0x),
			Title:    servant.Class.Name,
		}
		s.Rarity = servant.Rarity
		s.Attack = numberRange{Min: servant.Attack.Min, Max: servant.Attack.Max}
		s.HP = numberRange{Min: servant.HP.Min, Max: servant.HP.Max}
		for i, command := range servant.CommandCards {
			s.CommandCards[i].Title = command.Name
			s.CommandCards[i].Src_1_0x = "/images/" + getFilename(p.commands[command.Name].Src_1_0x)
			s.CommandCards[i].Src_1_5x = "/images/" + getFilename(p.commands[command.Name].Src_1_5x)
			s.CommandCards[i].Src_2_0x = "/images/" + getFilename(p.commands[command.Name].Src_2_0x)
		}
		s.NoblePhantasm = picture{
			Src_1_0x: "/images/" + getFilename(servant.NoblePhantasm.Src_1_0x),
			Src_1_5x: "/images/" + getFilename(servant.NoblePhantasm.Src_1_5x),
			Src_2_0x: "/images/" + getFilename(servant.NoblePhantasm.Src_2_0x),
			Title:    servant.NoblePhantasm.Name,
		}

		servants = append(servants, &s)
	}

	for icon := range iconSet {
		results.Images = append(results.Images, NewImageFromSrc(icon))
	}

	b, err := json.MarshalIndent(servants, "", "\t")
	if err != nil {
		return nil, err
	}
	results.Data = append(results.Data, &Data{
		Filepath: "../../static/data/servants.json",
		Text:     b,
	})

	return &results, nil
}

func getFilename(src string) string {
	items := strings.Split(src, "/")
	return items[len(items)-1]
}
