package types

type Invoice struct {
	ObuId         int     `json:"obuId"`
	TotalDistance float64 `json:"totalDistance"`
	TotalAmount   float64 `json:"totalAmount"`
}
