package productlisting

type Repository interface {
	GetProducts() []Product
}

type Service interface {
	GetProducts() []Product
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetProducts() []Product {
	return s.r.GetProducts()
}
