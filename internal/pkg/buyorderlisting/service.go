package buyorderlisting

type Repository interface {
	GetBuyOrders() []BuyOrder
}

type Service interface {
	GetBuyOrders() []BuyOrder
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetBuyOrders() []BuyOrder {
	return s.r.GetBuyOrders()
}
