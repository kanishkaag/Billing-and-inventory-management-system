package models

type Bill struct {
	ID          string
	CustomerID  string
	OutletID    string
	TotalAmount float64
	TimeStamp   int64
}
