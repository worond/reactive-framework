package entities

type Product struct {
	ID          int64
	Quantity    int32
	Description string
}

type Products map[int64]Product
