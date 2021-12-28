package paymentadding

import (
	"productmanagement/internal/pkg/supplierlisting"
	"productmanagement/internal/pkg/supplierupdating"
)

type Repository interface {
	AddPayment(p Payment)
	UpdateSupplier(s supplierupdating.Supplier)
	GetSupplier(id string) supplierlisting.Supplier
}

type Service interface {
	AddPayment(p Payment)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddPayment(p Payment) {
	s.r.AddPayment(p)

	supplier := s.r.GetSupplier(p.SupplierID)
	s.r.UpdateSupplier(supplierupdating.Supplier{
		ID:   p.SupplierID,
		Debt: supplier.Debt - p.Amount,
	})
}
