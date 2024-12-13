package models

type Edge struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	X1   float64 `json:"x1"`
	Y1   float64 `json:"y1"`
	X2   float64 `json:"x2"`
	Y2   float64 `json:"y2"`
}