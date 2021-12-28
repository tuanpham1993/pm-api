package pg

import (
	"os"
	"productmanagement/internal/pkg/buyorderadding"
	"productmanagement/internal/pkg/buyorderlisting"
	"productmanagement/internal/pkg/customeradding"
	"productmanagement/internal/pkg/customerlisting"
	"productmanagement/internal/pkg/customerupdating"
	"productmanagement/internal/pkg/debtcollectionadding"
	"productmanagement/internal/pkg/debtcollectionlisting"
	"productmanagement/internal/pkg/paymentadding"
	"productmanagement/internal/pkg/paymentlisting"
	"productmanagement/internal/pkg/productadding"
	"productmanagement/internal/pkg/productlisting"
	"productmanagement/internal/pkg/productupdating"
	"productmanagement/internal/pkg/sellorderadding"
	"productmanagement/internal/pkg/sellorderlisting"
	"productmanagement/internal/pkg/supplieradding"
	"productmanagement/internal/pkg/supplierlisting"
	"productmanagement/internal/pkg/supplierupdating"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	var err error

	stg := new(Storage)

	dsn := "host=" + os.Getenv("PGHOST") + " user=" + os.Getenv("PGUSER") + " password=" + os.Getenv("PGPASSWORD") + " dbname=" + os.Getenv("PGDATABASE") + " port=" + os.Getenv("PGPORT") + " sslmode=disable"
	stg.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "product-management.",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	return stg, nil
}

func (s *Storage) GetProducts() []productlisting.Product {
	list := []Product{}

	s.db.Find(&list)

	products := []productlisting.Product{}
	for _, item := range list {
		products = append(products, productlisting.Product{
			ID:         item.ID,
			Name:       item.Name,
			Unit:       item.Unit,
			Quantity:   item.Quantity,
			InputPrice: item.InputPrice,
		})
	}

	return products
}

func (s *Storage) GetProduct(id string) productlisting.Product {
	product := Product{ID: id}

	s.db.First(&product)

	return productlisting.Product{
		ID:         product.ID,
		Name:       product.Name,
		Quantity:   product.Quantity,
		InputPrice: product.InputPrice,
	}
}

func (s *Storage) AddProduct(p productadding.Product) {
	product := Product{Name: p.Name, Unit: p.Unit, InputPrice: p.InputPrice, Quantity: p.Quantity}

	s.db.Create(&product)
}

func (s *Storage) DeleteProduct(id string) {
	s.db.Delete(&Product{ID: id})
}

func (s *Storage) UpdateProduct(p productupdating.Product) {
	product := Product{ID: p.ID}
	s.db.Model(&product).Updates(p)
}

func (s *Storage) GetCustomers() []customerlisting.Customer {
	list := []Customer{}

	s.db.Find(&list)

	customers := []customerlisting.Customer{}
	for _, item := range list {
		customers = append(customers, customerlisting.Customer{
			ID:   item.ID,
			Name: item.Name,
			Debt: item.Debt,
		})
	}

	return customers
}

func (s *Storage) GetCustomer(id string) customerlisting.Customer {
	customer := Customer{ID: id}

	s.db.First(&customer)

	return customerlisting.Customer{
		ID:   customer.ID,
		Name: customer.Name,
		Debt: customer.Debt,
	}
}

func (s *Storage) AddCustomer(c customeradding.Customer) {
	customer := Customer{Name: c.Name, Debt: c.Debt}

	s.db.Create(&customer)
}

func (s *Storage) DeleteCustomer(id string) {
	s.db.Delete(&Customer{ID: id})
}

func (s *Storage) UpdateCustomer(p customerupdating.Customer) {
	customer := Customer{ID: p.ID}
	s.db.Model(&customer).Updates(p)
}

func (s *Storage) GetSuppliers() []supplierlisting.Supplier {
	list := []Supplier{}

	s.db.Find(&list)

	suppliers := []supplierlisting.Supplier{}
	for _, item := range list {
		suppliers = append(suppliers, supplierlisting.Supplier{
			ID:   item.ID,
			Name: item.Name,
			Debt: item.Debt,
		})
	}

	return suppliers
}

func (s *Storage) GetSupplier(id string) supplierlisting.Supplier {
	supplier := Supplier{ID: id}

	s.db.First(&supplier)

	return supplierlisting.Supplier{
		ID:   supplier.ID,
		Name: supplier.Name,
		Debt: supplier.Debt,
	}
}

func (s *Storage) AddSupplier(c supplieradding.Supplier) {
	supplier := Supplier{Name: c.Name, Debt: c.Debt}

	s.db.Create(&supplier)
}

func (s *Storage) DeleteSupplier(id string) {
	s.db.Delete(&Supplier{ID: id})
}

func (s *Storage) UpdateSupplier(p supplierupdating.Supplier) {
	supplier := Supplier{ID: p.ID}
	s.db.Model(&supplier).Updates(p)
}

func (s *Storage) GetSellOrders() []sellorderlisting.SellOrder {
	list := []SellOrder{}

	s.db.Preload("Customer").Preload("Items").Preload("Items.Product").Find(&list)

	sellOrders := []sellorderlisting.SellOrder{}
	for _, order := range list {
		sellOrderItems := []sellorderlisting.SellOrderItem{}

		var totalCost int64
		for _, item := range order.Items {
			sellOrderItems = append(sellOrderItems, sellorderlisting.SellOrderItem{
				ID:          item.ID,
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				Quantity:    item.Quantity,
				Price:       item.Price,
				Cost:        item.Price * item.Quantity,
			})

			totalCost += item.Price * item.Quantity
		}

		so := sellorderlisting.SellOrder{
			ID: order.ID,
			// CustomerID:   *order.CustomerID,
			// CustomerName: order.Customer.Name,
			CreatedAt: formatDate(order.CreatedAt),
			Items:     sellOrderItems,
			TotalCost: totalCost,
		}

		if order.CustomerID != nil {
			so.CustomerID = *order.CustomerID
			so.CustomerName = *&order.Customer.Name
		}

		sellOrders = append(sellOrders, so)
	}

	return sellOrders
}

func (s *Storage) AddSellOrder(so sellorderadding.SellOrder) {
	items := []SellOrderItem{}

	for _, item := range so.Items {
		items = append(items, SellOrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	sellOrder := SellOrder{CustomerID: so.CustomerID, Items: items}
	s.db.Create(&sellOrder)
}

func (s *Storage) GetBuyOrders() []buyorderlisting.BuyOrder {
	list := []BuyOrder{}

	s.db.Preload("Supplier").Preload("Items").Preload("Items.Product").Find(&list)

	buyOrders := []buyorderlisting.BuyOrder{}
	for _, order := range list {
		buyOrderItems := []buyorderlisting.BuyOrderItem{}

		var totalCost int64
		for _, item := range order.Items {
			buyOrderItems = append(buyOrderItems, buyorderlisting.BuyOrderItem{
				ID:          item.ID,
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				Quantity:    item.Quantity,
				Price:       item.Price,
				Cost:        item.Price * item.Quantity,
			})

			totalCost += item.Price * item.Quantity
		}

		buyOrders = append(buyOrders, buyorderlisting.BuyOrder{
			SupplierID:   order.SupplierID,
			SupplierName: order.Supplier.Name,
			Items:        buyOrderItems,
			TotalCost:    totalCost,
			CreatedAt:    formatDate(order.CreatedAt),
		})
	}

	return buyOrders
}

func (s *Storage) AddBuyOrder(so buyorderadding.BuyOrder) {
	items := []BuyOrderItem{}

	for _, item := range so.Items {
		items = append(items, BuyOrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	buyOrder := BuyOrder{SupplierID: so.SupplierID, Items: items}
	s.db.Create(&buyOrder)
}

func (s *Storage) GetPayments() []paymentlisting.Payment {
	list := []Payment{}

	s.db.Preload("Supplier").Find(&list)

	payments := []paymentlisting.Payment{}
	for _, payment := range list {
		payments = append(payments, paymentlisting.Payment{
			SupplierID:   payment.SupplierID,
			SupplierName: payment.Supplier.Name,
			Amount:       payment.Amount,
			CreatedAt:    formatDate(payment.CreatedAt),
		})
	}

	return payments
}

func (s *Storage) AddPayment(p paymentadding.Payment) {
	payment := Payment{SupplierID: p.SupplierID, Amount: p.Amount}
	s.db.Create(&payment)
}

func (s *Storage) GetDebtCollections() []debtcollectionlisting.DebtCollection {
	list := []DebtCollection{}

	s.db.Preload("Customer").Find(&list)

	debtCollections := []debtcollectionlisting.DebtCollection{}
	for _, debtCollection := range list {
		debtCollections = append(debtCollections, debtcollectionlisting.DebtCollection{
			CustomerID:   debtCollection.CustomerID,
			CustomerName: debtCollection.Customer.Name,
			Amount:       debtCollection.Amount,
			CreatedAt:    formatDate(debtCollection.CreatedAt),
		})
	}

	return debtCollections
}

func (s *Storage) AddDebtCollection(p debtcollectionadding.DebtCollection) {
	debtcollection := DebtCollection{CustomerID: p.CustomerID, Amount: p.Amount}
	s.db.Create(&debtcollection)
}

func formatDate(t time.Time) string {
	layout := "02/01/2006"
	return t.Format(layout)
}
