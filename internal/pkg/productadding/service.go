package productadding

import (
	"productmanagement/internal/pkg/productlisting"
)

type Repository interface {
	AddProduct(p Product)
	GetProducts() []productlisting.Product
	// GetProduct() productlisting.Product
}

type Service interface {
	AddProduct(p Product) productlisting.Product
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddProduct(p Product) productlisting.Product {
	s.r.AddProduct(p)
	// return s.r.GetProduct()
	return productlisting.Product{}
}
