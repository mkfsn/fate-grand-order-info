package processer

type numberRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type command = string

type i18nName struct {
	EN string `json:"en"`
	JP string `json:"jp"`
}

type picture struct {
	Src70px  string `json:"src_70px"`
	Src105px string `json:"src_105px"`
	Src140px string `json:"src_140px"`
	Title    string `json:"title"`
}

type Servant struct {
	ID            string      `json:"id"`
	Icon          picture     `json:"icon"`
	Name          i18nName    `json:"name"`
	Class         string      `json:"class"`
	Rarity        int         `json:"rarity"`
	Attack        numberRange `json:"attack"`
	HP            numberRange `json:"hp"`
	CommandCards  [5]command  `json:"command_cards"`
	NoblePhantasm command     `json:"noble_phantasm"`
}
