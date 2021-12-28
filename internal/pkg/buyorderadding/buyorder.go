package buyorderadding

type BuyOrder struct {
	SupplierID string         `json:"supplierId"`
	Items      []BuyOrderItem `json:"items"`
}
