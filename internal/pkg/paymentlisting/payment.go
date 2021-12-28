package paymentlisting

type Payment struct {
	ID           string `json:"id"`
	SupplierID   string `json:"supplierId"`
	SupplierName string `json:"supplierName"`
	Amount       int64  `json:"amount"`
	CreatedAt    string `json:"createdAt"`
}
