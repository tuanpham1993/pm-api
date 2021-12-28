package debtcollectionlisting

type Repository interface {
	GetDebtCollections() []DebtCollection
}

type Service interface {
	GetDebtCollections() []DebtCollection
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetDebtCollections() []DebtCollection {
	return s.r.GetDebtCollections()
}
