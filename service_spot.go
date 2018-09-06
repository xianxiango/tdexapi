package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) SpotBuy(wq SpotBuyRequest) (*SpotBuy, error) {
	params := make(map[string]interface{})
	params["amount"] = wq.Amount
	params["price"] = wq.Price
	params["symbol"] = wq.Symbol

	res, err := as.request("POST", "/spot/buy", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotBuy.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotBuy := SpotBuy{}

	if err := json.Unmarshal(textRes, &rawSpotBuy); err != nil {
		return nil, errors.Wrap(err, "rawSpotBuy unmarshal failed")
	}

	return &rawSpotBuy, nil
}

func (as *apiService) SpotSell(wq SpotSellRequest) (*SpotSell, error) {
	params := make(map[string]interface{})
	params["amount"] = wq.Amount
	params["price"] = wq.Price
	params["symbol"] = wq.Symbol

	res, err := as.request("POST", "/spot/sell", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotSell.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotSell := SpotSell{}

	if err := json.Unmarshal(textRes, &rawSpotSell); err != nil {
		return nil, errors.Wrap(err, "rawSpotSell unmarshal failed")
	}

	return &rawSpotSell, nil
}

func (as *apiService) SpotCancel(wq SpotCancelRequest) (*SpotCancel, error) {
	params := make(map[string]interface{})
	params["id"] = wq.ID
	params["symbol"] = wq.Symbol

	res, err := as.request("POST", "/spot/cancel", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotCancel.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotCancel := SpotCancel{}

	if err := json.Unmarshal(textRes, &rawSpotCancel); err != nil {
		return nil, errors.Wrap(err, "rawSpotCancel unmarshal failed")
	}

	return &rawSpotCancel, nil
}
func (as *apiService) SpotOrders() (*SpotOrders, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/spot/orders", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotOrders.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotOrders := SpotOrders{}

	if err := json.Unmarshal(textRes, &rawSpotOrders); err != nil {
		return nil, errors.Wrap(err, "rawSpotOrders unmarshal failed")
	}

	return &rawSpotOrders, nil
}

func (as *apiService) SpotHistory(wq SpotHistoryRequest) (*SpotHistory, error) {
	params := make(map[string]interface{})
	params["beginTime"] = wq.BeginTime
	params["endTime"] = wq.EndTime
	params["pageSize"] = wq.PageSize
	params["page"] = wq.Page

	res, err := as.request("POST", "/spot/history", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotHistory.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotHistory := SpotHistory{}

	if err := json.Unmarshal(textRes, &rawSpotHistory); err != nil {
		return nil, errors.Wrap(err, "rawSpotHistory unmarshal failed")
	}

	return &rawSpotHistory, nil
}
func (as *apiService) SpotStat(wq SpotStatRequest) (*SpotStat, error) {
	params := make(map[string]interface{})
	params["symbol"] = wq.Symbol
	params["beginTime"] = wq.BeginTime
	params["endTime"] = wq.EndTime

	res, err := as.request("POST", "/spot/stat", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SpotStat.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSpotStat := SpotStat{}

	if err := json.Unmarshal(textRes, &rawSpotStat); err != nil {
		return nil, errors.Wrap(err, "rawSpotStat unmarshal failed")
	}

	return &rawSpotStat, nil
}
