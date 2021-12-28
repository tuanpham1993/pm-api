package sellorderlisting

type Repository interface {
	GetSellOrders() []SellOrder
}

type Service interface {
	GetSellOrders() []SellOrder
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetSellOrders() []SellOrder {
	return s.r.GetSellOrders()
}
