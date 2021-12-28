package sellorderadding

type SellOrder struct {
	CustomerID *string         `json:"customerId"`
	Items      []SellOrderItem `json:"items"`
}
