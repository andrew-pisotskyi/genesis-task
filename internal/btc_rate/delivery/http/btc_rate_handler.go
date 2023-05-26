package http

import (
	"net/http"

	"github.com/andrew-pisotskyi/genesis-task/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type BtcRateHandler struct {
	BtcRateUsecase domain.BtcRateUsecaseInterface
}

func NewBtcRateHandler(btcRateUsecase domain.BtcRateUsecaseInterface) *BtcRateHandler {
	return &BtcRateHandler{btcRateUsecase}
}

func (brh *BtcRateHandler) GetRate(c echo.Context) error {
	rate, err := brh.BtcRateUsecase.GetRate()
	if err != nil {
		return c.JSON(getStatusCode(err), domain.ResponseError{Message: "Invalid status value"})
	}

	return c.JSON(http.StatusOK, rate)
}

func (brh *BtcRateHandler) Subscribe(c echo.Context) error {
	email := c.FormValue("email")
	subscribeRequest := domain.SubscribeRequest{Email: c.FormValue("email")}
	if ok, err := isRequestValid(&subscribeRequest); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := brh.BtcRateUsecase.Subscribe(email)
	if err != nil {
		return c.JSON(getStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	c.Response().Header().Add("content-type", "application/json")
	return c.NoContent(http.StatusOK)
}

func (brh *BtcRateHandler) SendEmails(c echo.Context) error {
	err := brh.BtcRateUsecase.SendEmails()
	if err != nil {
		return c.JSON(getStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	c.Response().Header().Add("content-type", "application/json")
	return c.NoContent(http.StatusOK)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)
	switch err {
	case domain.ErrEmailExists:
		return http.StatusConflict
	default:
		return http.StatusBadRequest
	}
}

func isRequestValid(m *domain.SubscribeRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
