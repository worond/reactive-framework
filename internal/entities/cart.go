package entities

type Cart struct {
	UserID int64
	Items  []CartItem
}

type CartItem struct {
	ID          int64
	Quantity    int32
	Description string
	Price       Price
	Label       string
}
