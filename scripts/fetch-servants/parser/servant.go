package parser

type IntegerRange struct {
	Min int
	Max int
}

type CommandCard struct {
	Image
	Name string
}

type Image struct {
	// FIXME: Don't know how to name variable for 1.5x image, and
	// to make it consistent, use snake_case for all fields.
	Src_1_0x string
	Src_1_5x string
	Src_2_0x string
}

type Servant struct {
	ID   string
	Icon Image
	Name struct {
		EN string
		JP string
	}
	Class struct {
		Image
		Name string
	}
	Rarity        int
	Attack        IntegerRange
	HP            IntegerRange
	CommandCards  [5]CommandCard
	NoblePhantasm CommandCard
}
