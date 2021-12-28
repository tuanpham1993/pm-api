package sellorderlisting

type SellOrder struct {
	ID           string          `json:"id"`
	CustomerID   string          `json:"customerId"`
	CustomerName string          `json:"customerName"`
	Items        []SellOrderItem `json:"items"`
	TotalCost    int64           `json:"totalCost"`
	CreatedAt    string          `json:"createdAt"`
}

type SellOrderItem struct {
	ID          string `gorm:"type:uuid;primary_key;" json:"id"`
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Cost        int64  `json:"cost"`
}
