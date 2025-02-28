package models

type Product struct {
	ID    string
	Brand string
	Name  string
}

type ProductRecord struct {
	CostPrice    string
	SellingPrice string
	ExpiryDate   int64
	SupplierID   string
}
