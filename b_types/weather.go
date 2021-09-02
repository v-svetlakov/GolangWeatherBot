package b_types

type Weathers struct {
	Base        string        `json:"base"`
	Coord       Coord         `json:"coord"`
	Visibility  float64       `json:"visibility"`
	Main        Main          `json:"main"`
	Wind        Wind          `json:"wind"`
}
