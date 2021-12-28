package supplierupdating

type Supplier struct {
	ID   string
	Name string `json:"name"`
	Debt int64  `json:"debt"`
}
