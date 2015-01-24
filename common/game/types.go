package game

type Gopher struct {
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Friend bool `json:"friend"`
}

type State struct {
	GopherId int      `json:"gid"`
	Health   int      `json:"health"`
	Ammo     int      `json:"ammo"`
	Me       Gopher   `json:"me"`
	Nearby   []Gopher `json:"nearby"`

	Test string
}

type Action struct {
	Test string
}
