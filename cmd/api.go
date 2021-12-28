package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"

	"productmanagement/internal/pkg/buyorderadding"
	"productmanagement/internal/pkg/buyorderlisting"
	"productmanagement/internal/pkg/customeradding"
	"productmanagement/internal/pkg/customerdeleting"
	"productmanagement/internal/pkg/customerlisting"
	"productmanagement/internal/pkg/customerupdating"
	"productmanagement/internal/pkg/debtcollectionadding"
	"productmanagement/internal/pkg/debtcollectionlisting"
	"productmanagement/internal/pkg/http/rest"
	"productmanagement/internal/pkg/paymentadding"
	"productmanagement/internal/pkg/paymentlisting"
	"productmanagement/internal/pkg/productadding"
	"productmanagement/internal/pkg/productdeleting"
	"productmanagement/internal/pkg/productlisting"
	"productmanagement/internal/pkg/productupdating"
	"productmanagement/internal/pkg/sellorderadding"
	"productmanagement/internal/pkg/sellorderlisting"
	"productmanagement/internal/pkg/storage/pg"
	"productmanagement/internal/pkg/supplieradding"
	"productmanagement/internal/pkg/supplierdeleting"
	"productmanagement/internal/pkg/supplierlisting"
	"productmanagement/internal/pkg/supplierupdating"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if os.Getenv("ENV") != "production" {
		godotenv.Load()
	}

	storage, err := pg.NewStorage()

	if err != nil {
		panic(err)
	}

	productLister := productlisting.NewService(storage)
	productAdder := productadding.NewService(storage)
	productDeleter := productdeleting.NewService(storage)
	productUpdater := productupdating.NewService(storage)

	customerLister := customerlisting.NewService(storage)
	customerAdder := customeradding.NewService(storage)
	customerDeleter := customerdeleting.NewService(storage)
	customerUpdater := customerupdating.NewService(storage)

	supplierLister := supplierlisting.NewService(storage)
	supplierAdder := supplieradding.NewService(storage)
	supplierDeleter := supplierdeleting.NewService(storage)
	supplierUpdater := supplierupdating.NewService(storage)

	sellOrderAdder := sellorderadding.NewService(storage)
	sellOrderLister := sellorderlisting.NewService(storage)

	buyOrderAdder := buyorderadding.NewService(storage)
	buyOrderLister := buyorderlisting.NewService(storage)

	paymentAdder := paymentadding.NewService(storage)
	paymentLister := paymentlisting.NewService(storage)

	debtCollectionAdder := debtcollectionadding.NewService(storage)
	debtCollectionLister := debtcollectionlisting.NewService(storage)

	r := rest.Handler(
		productLister,
		productAdder,
		productDeleter,
		productUpdater,
		customerLister,
		customerAdder,
		customerDeleter,
		customerUpdater,
		supplierLister,
		supplierAdder,
		supplierDeleter,
		supplierUpdater,
		sellOrderAdder,
		sellOrderLister,
		buyOrderAdder,
		buyOrderLister,
		paymentAdder,
		paymentLister,
		debtCollectionAdder,
		debtCollectionLister,
	)

	r.Run()
}
