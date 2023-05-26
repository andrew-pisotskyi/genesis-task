package domain

type SubscribeRequest struct {
	Email string `validate:"required,email"`
}

type BtcUahRate struct {
	Price string `json:"price"`
}

type BtcRateUsecaseInterface interface {
	GetRate() (int64, error)
	Subscribe(email string) error
	SendEmails() error
}

type BtcRateRepositoryInterface interface {
	ExistsEmail(email string) bool
	SaveEmail(email string) error
	GetAllEmails() []string
}
