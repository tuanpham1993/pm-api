package supplierlisting

type Repository interface {
	GetSuppliers() []Supplier
}

type Service interface {
	GetSuppliers() []Supplier
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetSuppliers() []Supplier {
	return s.r.GetSuppliers()
}
