package productupdating

type Product struct {
	ID string
	Name       string `json:"name"`
	Unit       string `json:"unit"`
	Quantity   int64  `json:"quantity"`
	InputPrice int64  `json:"inputPrice"`
}
