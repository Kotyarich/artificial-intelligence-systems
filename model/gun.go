package model

type Gun struct {
	ID      uint
	Model   string `json:"model"`
	Company string
	Country string
	Year    uint

	Type          string
	LoadType 	  string
	// For break-action rifles only
	BreakType	  string
	IsBullPup	  bool
	IsSliced      bool
	BarrelsNumber uint
	Caliber		  string

	Length       uint
	BarrelLength string
	Weight       float32
}
