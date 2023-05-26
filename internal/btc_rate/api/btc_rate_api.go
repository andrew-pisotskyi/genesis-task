package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrew-pisotskyi/genesis-task/domain"
	log "github.com/sirupsen/logrus"
)

type Api struct {
	baseUrl string
}

func NewApi(btcUahApiUrl string) *Api {
	return &Api{
		baseUrl: btcUahApiUrl,
	}
}

func (api *Api) GetBtcUahRate() (float64, error) {
	response, err := http.Get(api.baseUrl)
	if err != nil {
		log.Errorf("Error when getting btc-uah rate: %s", err.Error())
		return 0.0, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("something went wrong; response status code: %d", response.StatusCode)
		log.Error(err)
		return 0.0, err
	}
	defer response.Body.Close()

	var btcUahRate domain.BtcUahRate
	if err := json.NewDecoder(response.Body).Decode(&btcUahRate); err != nil {
		log.Errorf("Could not decode incoming update. err: %s", err.Error())
		return 0.0, err
	}
	price, err := strconv.ParseFloat(btcUahRate.Price, 64)
	if err != nil {
		log.Errorf("Could not parse float by string. err: %s", err.Error())
		return 0.0, err
	}

	return price, err
}
