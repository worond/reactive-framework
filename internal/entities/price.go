package entities

type Price struct {
	Original   float64
	Discounted float64
}

type Prices map[int64]Price
