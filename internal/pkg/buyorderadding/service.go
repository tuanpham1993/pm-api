package buyorderadding

import (
	"productmanagement/internal/pkg/productlisting"
	"productmanagement/internal/pkg/productupdating"
	"productmanagement/internal/pkg/supplierlisting"
	"productmanagement/internal/pkg/supplierupdating"
)

type Repository interface {
	AddBuyOrder(so BuyOrder)
	UpdateProduct(p productupdating.Product)
	GetProduct(id string) productlisting.Product
	UpdateSupplier(s supplierupdating.Supplier)
	GetSupplier(id string) supplierlisting.Supplier
}

type Service interface {
	AddBuyOrder(so BuyOrder)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddBuyOrder(so BuyOrder) {
	s.r.AddBuyOrder(so)

	var totalCost int64

	for _, item := range so.Items {
		product := s.r.GetProduct(item.ProductID)
		s.r.UpdateProduct(productupdating.Product{
			ID:       item.ProductID,
			Quantity: product.Quantity + item.Quantity,
		})

		totalCost += item.Quantity * item.Price
	}

	supplier := s.r.GetSupplier(so.SupplierID)
	s.r.UpdateSupplier(supplierupdating.Supplier{
		ID:   so.SupplierID,
		Debt: supplier.Debt + totalCost,
	})
}
