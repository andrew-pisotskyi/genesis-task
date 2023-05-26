package route

import (
	btcApi "github.com/andrew-pisotskyi/genesis-task/internal/btc_rate/api"
	btcRateHandler "github.com/andrew-pisotskyi/genesis-task/internal/btc_rate/delivery/http"
	btrRateRepository "github.com/andrew-pisotskyi/genesis-task/internal/btc_rate/repository/file"
	btcRateUsecase "github.com/andrew-pisotskyi/genesis-task/internal/btc_rate/usecase"
	"github.com/andrew-pisotskyi/genesis-task/internal/mail"
	"github.com/labstack/echo/v4"
)

func (r *Router) NewBtcRateRoute(groupRoute *echo.Group) {
	bApi := btcApi.NewApi(r.config.BtcUahApiURl)
	bRepo := btrRateRepository.NewBtcRateRepository(r.config.EmailsDataFile)
	mailService := mail.NewMail(r.config.SmtpHost, r.config.SmtpPort, r.config.EmailFrom, r.config.EmailFromPassword)

	btcRateUsecase := btcRateUsecase.NewBtcRateUsecase(bApi, bRepo, mailService)
	handler := btcRateHandler.NewBtcRateHandler(btcRateUsecase)

	groupRoute.GET("/rate", handler.GetRate)
	groupRoute.POST("/subscribe", handler.Subscribe)
	groupRoute.POST("/sendEmails", handler.SendEmails)
}
