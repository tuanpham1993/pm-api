package debtcollectionadding

type DebtCollection struct {
	CustomerID   string `json:"customerId"`
	CustomerName string `json:"customerName"`
	Amount       int64  `json:"amount"`
}
