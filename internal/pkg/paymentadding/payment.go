package paymentadding

type Payment struct {
	SupplierID   string `json:"supplierId"`
	SupplierName string `json:"supplierName"`
	Amount       int64  `json:"amount"`
}
