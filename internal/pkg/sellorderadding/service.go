package sellorderadding

import (
	"productmanagement/internal/pkg/customerlisting"
	"productmanagement/internal/pkg/customerupdating"
	"productmanagement/internal/pkg/productlisting"
	"productmanagement/internal/pkg/productupdating"
)

type Repository interface {
	AddSellOrder(so SellOrder)
	UpdateProduct(p productupdating.Product)
	GetProduct(id string) productlisting.Product
	UpdateCustomer(s customerupdating.Customer)
	GetCustomer(id string) customerlisting.Customer
}

type Service interface {
	AddSellOrder(so SellOrder)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddSellOrder(so SellOrder) {
	s.r.AddSellOrder(so)

	var totalCost int64

	for _, item := range so.Items {
		product := s.r.GetProduct(item.ProductID)
		s.r.UpdateProduct(productupdating.Product{
			ID:       item.ProductID,
			Quantity: product.Quantity - item.Quantity,
		})

		totalCost += item.Quantity * item.Price

	}

	if so.CustomerID != nil {
		customer := s.r.GetCustomer(*so.CustomerID)
		s.r.UpdateCustomer(customerupdating.Customer{
			ID:   *so.CustomerID,
			Debt: customer.Debt + totalCost,
		})
	}
}
