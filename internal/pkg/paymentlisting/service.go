package paymentlisting

type Repository interface {
	GetPayments() []Payment
}

type Service interface {
	GetPayments() []Payment
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetPayments() []Payment {
	return s.r.GetPayments()
}
