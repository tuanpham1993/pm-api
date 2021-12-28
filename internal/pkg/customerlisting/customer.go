package customerlisting

type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Debt int64  `json:"debt"`
}
