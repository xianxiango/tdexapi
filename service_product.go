package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) ProductCoins() (*ProductCoins, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/product/coins", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from ProductCoins.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	
	rawProductCoins := ProductCoins{}
	if err := json.Unmarshal(textRes, &rawProductCoins); err != nil {
		return nil, errors.Wrap(err, "rawProductCoins unmarshal failed")
	}

	
	return &rawProductCoins, nil
}
