package controllers

type TransactionList struct {
	Type           string `json:"type"`
	Amount         string `json:"amount"`
	Date           string `json:"date"`
	Memo           string `json:"memo"`
	Destination    string `json:"destination"`
	Source         string `json:"source"`
	Hash           string `json:"hash"`
	Asset          string `json:"asset"`
	TransactionUrl string `json:"url"`
}

type TransactionHistory struct {
	TransactionList []TransactionList `json:"transactionList"`
}
