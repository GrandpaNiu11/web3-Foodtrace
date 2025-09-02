package models

type FoodItems struct {
	ID                uint
	FoodName          string
	OutboundPrice     float64
	ShippingWarehouse string
	OutboundDate      string
	Code              string
	Description       string
	BlockchainAddress string
	TransactionHash   string
}

func (f FoodItems) TableName() string {
	return "food_items"
}
