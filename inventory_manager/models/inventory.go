package models

type Inventory struct {
	ID       string
	ShopID   string
	Products []ProductRecord
}
