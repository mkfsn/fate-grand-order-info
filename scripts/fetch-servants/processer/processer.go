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

	var servants []*Servant
	for _, servant := range p.servants {
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src70px))
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src105px))
		results.Images = append(results.Images, newImageFromSrc(servant.Icon.Src140px))

		servants = append(servants, &Servant{
			ID: servant.ID,
			Icon: picture{
				Src70px:  "/images/" + getFilename(servant.Icon.Src70px),
				Src105px: "/images/" + getFilename(servant.Icon.Src105px),
				Src140px: "/images/" + getFilename(servant.Icon.Src140px),
				Title:    servant.Name.JP,
			},
			Name: i18nName{
				EN: servant.Name.EN,
				JP: servant.Name.JP,
			},
			Class:  servant.Class,
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
