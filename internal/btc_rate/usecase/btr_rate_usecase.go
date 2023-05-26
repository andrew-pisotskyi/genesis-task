package usecase

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/andrew-pisotskyi/genesis-task/domain"
	"github.com/andrew-pisotskyi/genesis-task/internal/btc_rate/api"
)

type btcRateUsecase struct {
	btcApi      *api.Api
	btcRateRepo domain.BtcRateRepositoryInterface
	mailService domain.MailInterface
}

func NewBtcRateUsecase(
	btcApi *api.Api,
	btcRateRepo domain.BtcRateRepositoryInterface,
	mailService domain.MailInterface,
) domain.BtcRateUsecaseInterface {
	return &btcRateUsecase{btcApi, btcRateRepo, mailService}
}

func (btu *btcRateUsecase) GetRate() (int64, error) {
	rate, err := btu.btcApi.GetBtcUahRate()
	if err != nil {
		log.Errorf("Something went wrong in btc api request. err: %s", err.Error())
		return 0, err
	}

	return int64(rate), nil
}

func (btu *btcRateUsecase) Subscribe(email string) error {
	existsEmail := btu.btcRateRepo.ExistsEmail(email)
	if existsEmail {
		return domain.ErrEmailExists
	}
	btu.btcRateRepo.SaveEmail(email)
	return nil
}

func (btu *btcRateUsecase) SendEmails() error {
	rate, err := btu.btcApi.GetBtcUahRate()
	if err != nil {
		log.Errorf("Something went wrong in btc api request. err: %s", err.Error())
		return err
	}
	emails := btu.btcRateRepo.GetAllEmails()
	message := fmt.Sprintf("At the moment, the exchange rate of BTC to UAH is %d", int64(rate))

	return btu.mailService.Send(emails, message)
}
