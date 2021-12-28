package sellorderadding

type SellOrderItem struct {
	SellOrderID string `json:"sellOrderId"`
	ProductID   string `json:"productId"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
}
