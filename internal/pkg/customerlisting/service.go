package customerlisting

type Repository interface {
	GetCustomers() []Customer
}

type Service interface {
	GetCustomers() []Customer
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetCustomers() []Customer {
	return s.r.GetCustomers()
}
