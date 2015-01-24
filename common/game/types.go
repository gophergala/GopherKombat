package game

type GopherData struct {
	X      int  `json:"x"`
	Y      int  `json:"y"`
	Friend bool `json:"friend"`
}

type State struct {
	GopherId int          `json:"gid"`
	Health   int          `json:"health"`
	Ammo     int          `json:"ammo"`
	Me       GopherData   `json:"me"`
	Nearby   []GopherData `json:"nearby"`

	Test string
}

type Action struct {
	Test string
}
