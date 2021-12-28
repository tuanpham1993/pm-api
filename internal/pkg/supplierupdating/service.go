package supplierupdating

type Repository interface {
	UpdateSupplier(s Supplier)
}

type Service interface {
	UpdateSupplier(s Supplier)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateSupplier(sp Supplier) {
	s.r.UpdateSupplier(sp)
}
