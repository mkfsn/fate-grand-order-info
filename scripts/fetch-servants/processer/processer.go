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

func newImageFromSrc(src string) *Image {
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
}

func New(servants []*parser.Servant) *Processer {
	return &Processer{servants: servants}
}

func (p *Processer) Process() (*Results, error) {
	var results Results

	iconSet := make(map[string]bool)

	var servants []*Servant
	for _, servant := range p.servants {
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src_1_0x))
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src_1_5x))
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src_2_0x))
		iconSet[servant.Class.Src_1_0x] = true
		iconSet[servant.Class.Src_1_5x] = true
		iconSet[servant.Class.Src_2_0x] = true

		servants = append(servants, &Servant{
			ID: servant.ID,
			Icon: picture{
				Src_1_0x: "/images/" + getFilename(servant.Icon.Src_1_0x),
				Src_1_5x: "/images/" + getFilename(servant.Icon.Src_1_5x),
				Src_2_0x: "/images/" + getFilename(servant.Icon.Src_2_0x),
				Title:    servant.Name.JP,
			},
			Name: i18nName{
				EN: servant.Name.EN,
				JP: servant.Name.JP,
			},
			Class: picture{
				Src_1_0x: "/images/" + getFilename(servant.Class.Src_1_0x),
				Src_1_5x: "/images/" + getFilename(servant.Class.Src_1_5x),
				Src_2_0x: "/images/" + getFilename(servant.Class.Src_2_0x),
				Title:    servant.Class.Name,
			},
			Rarity: servant.Rarity,
			Attack: numberRange{
				Min: servant.Attack.Min,
				Max: servant.Attack.Max,
			},
			HP: numberRange{
				Min: servant.HP.Min,
				Max: servant.HP.Max,
			},
			CommandCards:  servant.CommandCards,
			NoblePhantasm: servant.NoblePhantasm,
		})
	}

	for icon := range iconSet {
		results.Images = append(results.Images, newImageFromSrc(icon))
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
