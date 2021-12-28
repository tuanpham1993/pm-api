package productupdating

type Repository interface {
	UpdateProduct(p Product)
}

type Service interface {
	UpdateProduct(p Product)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateProduct(p Product) {
	s.r.UpdateProduct(p)
}
