package productadding

type Product struct {
	Name       string `json:"name"`
	Unit       string `json:"unit"`
	Quantity   int64  `json:"quantity"`
	InputPrice int64  `json:"inputPrice"`
}
