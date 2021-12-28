package rest

import (
	"productmanagement/internal/pkg/buyorderadding"
	"productmanagement/internal/pkg/buyorderlisting"
	"productmanagement/internal/pkg/customeradding"
	"productmanagement/internal/pkg/customerdeleting"
	"productmanagement/internal/pkg/customerlisting"
	"productmanagement/internal/pkg/customerupdating"
	"productmanagement/internal/pkg/debtcollectionadding"
	"productmanagement/internal/pkg/debtcollectionlisting"
	"productmanagement/internal/pkg/paymentadding"
	"productmanagement/internal/pkg/paymentlisting"
	"productmanagement/internal/pkg/productadding"
	"productmanagement/internal/pkg/productdeleting"
	"productmanagement/internal/pkg/productlisting"
	"productmanagement/internal/pkg/productupdating"
	"productmanagement/internal/pkg/sellorderadding"
	"productmanagement/internal/pkg/sellorderlisting"
	"productmanagement/internal/pkg/supplieradding"
	"productmanagement/internal/pkg/supplierdeleting"
	"productmanagement/internal/pkg/supplierlisting"
	"productmanagement/internal/pkg/supplierupdating"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func Handler(
	pl productlisting.Service,
	pa productadding.Service,
	pd productdeleting.Service,
	pu productupdating.Service,
	cl customerlisting.Service,
	ca customeradding.Service,
	cd customerdeleting.Service,
	cu customerupdating.Service,
	sl supplierlisting.Service,
	sa supplieradding.Service,
	sd supplierdeleting.Service,
	su supplierupdating.Service,
	soa sellorderadding.Service,
	sol sellorderlisting.Service,
	boa buyorderadding.Service,
	bol buyorderlisting.Service,
	paa paymentadding.Service,
	pal paymentlisting.Service,
	dca debtcollectionadding.Service,
	dcl debtcollectionlisting.Service,
) *gin.Engine {

	r := gin.New()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/ping"},
	}))

	// r.Use(gin.Recovery())

	r.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://app.mailent.com", "https://mailent.com"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PATCH", "DELETE"},
		Debug:            false,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products", getProducts(pl))
	r.POST("/products", addProduct(pa))
	r.DELETE("/products/:id", deleteProduct(pd))
	r.PATCH("/products/:id", updateProduct(pu))

	r.GET("/customers", getCustomers(cl))
	r.POST("/customers", addCustomer(ca))
	r.DELETE("/customers/:id", deleteCustomer(cd))
	r.PATCH("/customers/:id", updateCustomer(cu))

	r.GET("/suppliers", getSuppliers(sl))
	r.POST("/suppliers", addSupplier(sa))
	r.DELETE("/suppliers/:id", deleteSupplier(sd))
	r.PATCH("/suppliers/:id", updateSupplier(su))

	r.POST("/sell-orders", addSellOrder(soa))
	r.GET("/sell-orders", getSellOrders(sol))

	r.POST("/buy-orders", addBuyOrder(boa))
	r.GET("/buy-orders", getBuyOrders(bol))

	r.POST("/payments", addPayment(paa))
	r.GET("/payments", getPayments(pal))

	r.POST("/debt-collections", addDebtCollection(dca))
	r.GET("/debt-collections", getDebtCollections(dcl))

	return r
}

func getProducts(pl productlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, pl.GetProducts())
	}
}

func addProduct(pa productadding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p productadding.Product
		c.BindJSON(&p)
		pa.AddProduct(p)
		c.Status(201)
	}
}

func deleteProduct(pd productdeleting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		pd.DeleteProduct(id)
		c.Status(204)
	}
}

func updateProduct(pu productupdating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p productupdating.Product
		c.BindJSON(&p)
		p.ID = c.Param("id")
		pu.UpdateProduct(p)
		c.Status(204)
	}
}

func getCustomers(cl customerlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, cl.GetCustomers())
	}
}

func addCustomer(ca customeradding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p customeradding.Customer
		c.BindJSON(&p)
		ca.AddCustomer(p)
		c.Status(201)
	}
}

func deleteCustomer(cd customerdeleting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		cd.DeleteCustomer(id)
		c.Status(204)
	}
}

func updateCustomer(cu customerupdating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p customerupdating.Customer
		c.BindJSON(&p)
		p.ID = c.Param("id")
		cu.UpdateCustomer(p)
		c.Status(204)
	}
}

func getSuppliers(sl supplierlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, sl.GetSuppliers())
	}
}

func addSupplier(sa supplieradding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s supplieradding.Supplier
		c.BindJSON(&s)
		sa.AddSupplier(s)
		c.Status(201)
	}
}

func deleteSupplier(cd supplierdeleting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		cd.DeleteSupplier(id)
		c.Status(204)
	}
}

func updateSupplier(su supplierupdating.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s supplierupdating.Supplier
		c.BindJSON(&s)
		s.ID = c.Param("id")
		su.UpdateSupplier(s)
		c.Status(204)
	}
}

func addSellOrder(soa sellorderadding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var so sellorderadding.SellOrder
		c.BindJSON(&so)
		soa.AddSellOrder(so)
		c.Status(201)
	}
}

func getSellOrders(sol sellorderlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, sol.GetSellOrders())
	}
}

func addBuyOrder(boa buyorderadding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bo buyorderadding.BuyOrder
		c.BindJSON(&bo)
		boa.AddBuyOrder(bo)
		c.Status(201)
	}
}

func getBuyOrders(bol buyorderlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, bol.GetBuyOrders())
	}
}

func addPayment(pa paymentadding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p paymentadding.Payment
		c.BindJSON(&p)
		pa.AddPayment(p)
		c.Status(201)
	}
}

func getPayments(pl paymentlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, pl.GetPayments())
	}
}

func addDebtCollection(dca debtcollectionadding.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p debtcollectionadding.DebtCollection
		c.BindJSON(&p)
		dca.AddDebtCollection(p)
		c.Status(201)
	}
}

func getDebtCollections(dcl debtcollectionlisting.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, dcl.GetDebtCollections())
	}
}
