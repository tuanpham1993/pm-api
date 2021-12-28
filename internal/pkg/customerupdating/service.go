package customerupdating

type Repository interface {
	UpdateCustomer(c Customer)
}

type Service interface {
	UpdateCustomer(c Customer)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateCustomer(c Customer) {
	s.r.UpdateCustomer(c)
}
