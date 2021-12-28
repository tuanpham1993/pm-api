package supplieradding

import (
	"productmanagement/internal/pkg/supplierlisting"
)

type Repository interface {
	AddSupplier(c Supplier)
	GetSuppliers() []supplierlisting.Supplier
}

type Service interface {
	AddSupplier(p Supplier) supplierlisting.Supplier
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddSupplier(c Supplier) supplierlisting.Supplier {
	s.r.AddSupplier(c)
	return supplierlisting.Supplier{}
}
