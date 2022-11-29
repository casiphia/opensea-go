package model

type Owner struct {
	Owner       *Account `opensea:"owner" json:"owner"`
	Quantity    string   `opensea:"quantity" json:"quantity"`
	CreatedDate string   `opensea:"created_date" json:"created_date"`
}
