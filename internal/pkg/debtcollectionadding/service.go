package debtcollectionadding

import (
	"productmanagement/internal/pkg/customerlisting"
	"productmanagement/internal/pkg/customerupdating"
)

type Repository interface {
	AddDebtCollection(p DebtCollection)
	UpdateCustomer(s customerupdating.Customer)
	GetCustomer(id string) customerlisting.Customer
}

type Service interface {
	AddDebtCollection(p DebtCollection)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddDebtCollection(dc DebtCollection) {
	s.r.AddDebtCollection(dc)

	customer := s.r.GetCustomer(dc.CustomerID)
	s.r.UpdateCustomer(customerupdating.Customer{
		ID:   dc.CustomerID,
		Debt: customer.Debt + dc.Amount,
	})
}
