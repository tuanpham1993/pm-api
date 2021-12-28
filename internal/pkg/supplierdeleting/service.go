package supplierdeleting

type Repository interface {
	DeleteSupplier(id string)
}

type Service interface {
	DeleteSupplier(id string)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteSupplier(id string) {
	s.r.DeleteSupplier(id)
}
