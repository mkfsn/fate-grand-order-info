package parser

type IntegerRange struct {
	Min int
	Max int
}

type CommandCard = string

type Servant struct {
	ID   string
	Icon struct {
		Src70px  string
		Src105px string
		Src140px string
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
