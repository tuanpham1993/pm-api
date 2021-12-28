package debtcollectionlisting

type DebtCollection struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customerId"`
	CustomerName string `json:"customerName"`
	Amount       int64  `json:"amount"`
	CreatedAt    string `json:"createdAt"`
}
