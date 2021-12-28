package customeradding

import (
	"productmanagement/internal/pkg/customerlisting"
)

type Repository interface {
	AddCustomer(c Customer)
	GetCustomers() []customerlisting.Customer
}

type Service interface {
	AddCustomer(p Customer) customerlisting.Customer
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCustomer(c Customer) customerlisting.Customer {
	s.r.AddCustomer(c)
	return customerlisting.Customer{}
}
