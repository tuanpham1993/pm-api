package buyorderlisting

type BuyOrder struct {
	ID           string         `json:"id"`
	SupplierID   string         `json:"supplierId"`
	SupplierName string         `json:"supplierName"`
	Items        []BuyOrderItem `json:"items"`
	TotalCost    int64          `json:"totalCost"`
	CreatedAt    string         `json:"createdAt"`
}

type BuyOrderItem struct {
	ID          string `gorm:"type:uuid;primary_key;" json:"id"`
	ProductID   string `json:"productId"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Cost        int64  `json:"cost"`
}
