package productdeleting

type Repository interface {
	DeleteProduct(id string)
}

type Service interface {
	DeleteProduct(id string)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteProduct(id string) {
	s.r.DeleteProduct(id)
}
