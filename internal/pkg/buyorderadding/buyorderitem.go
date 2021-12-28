package buyorderadding

type BuyOrderItem struct {
	BuyOrderID string `json:"buyOrderId"`
	ProductID  string `json:"productId"`
	Price      int64  `json:"price"`
	Quantity   int64  `json:"quantity"`
}
