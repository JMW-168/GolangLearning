package model

type Task struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	VendorID int    `json:"vendor_id"`
	IsLimit  int    `json:"is_limit"`
}
