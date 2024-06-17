package types

type ObuData struct {
	ObuId int `json:"obuId"`
	Geo   Geo `json:"geo"`
}

type Distance struct {
	Value float64 `json:"value"`
	ObuId int     `json:"obuId"`
	Unix  int64   `json:"unix/"`
}
