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
	Src_1_0x string `json:"src-1x"`
	Src_1_5x string `json:"src-1.5x"`
	Src_2_0x string `json:"src-2x"`
	Title    string `json:"title"`
}

type Servant struct {
	ID            string      `json:"id"`
	Icon          picture     `json:"icon"`
	Name          i18nName    `json:"name"`
	Class         picture     `json:"class"`
	Rarity        int         `json:"rarity"`
	Attack        numberRange `json:"attack"`
	HP            numberRange `json:"hp"`
	CommandCards  [5]picture  `json:"command_cards"`
	NoblePhantasm picture     `json:"noble_phantasm"`
}
