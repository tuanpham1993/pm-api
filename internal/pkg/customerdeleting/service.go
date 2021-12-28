package customerdeleting

type Repository interface {
	DeleteCustomer(id string)
}

type Service interface {
	DeleteCustomer(id string)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteCustomer(id string) {
	s.r.DeleteCustomer(id)
}
