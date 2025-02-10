package models

type Receipt struct {
	retailer     string
	purchaseDate string
	purchaseItem string
	items        []Item
	total        string
}

type Item struct {
	shortDescription string
	price            string
}
