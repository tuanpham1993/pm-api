package customerupdating

type Customer struct {
	ID   string
	Name string `json:"name"`
	Debt int64  `json:"debt"`
}
