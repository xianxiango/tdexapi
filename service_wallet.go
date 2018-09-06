package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) Balances(bq BalancesRequest) (*Balances, error) {
	params := make(map[string]interface{})
	params["type"] = bq.Type
	res, err := as.request("POST", "/wallet/balances", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Balances.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	// rawBalances := struct {
	// 	Status uint32 `json:"status"`
	// 	Data   struct {
	// 		List []struct {
	// 			Type      uint32  `json:"type"`
	// 			Currency  uint32  `json:"currency"`
	// 			Quantity  float64 `json:"quantity"`
	// 			Withdraw  float64 `json:"withdraw"`
	// 			Transfer  float64 `json:"transfer"`
	// 			Deposit   float64 `json:"deposit"`
	// 			Lock      uint32  `json:"lock"`
	// 			Available uint32  `json:"available"`
	// 			Status    uint32  `json:"status"`
	// 		}
	// 	}
	// }{}
	rawBalances := Balances{}

	if err := json.Unmarshal(textRes, &rawBalances); err != nil {
		return nil, errors.Wrap(err, "rawBalances unmarshal failed")
	}

	// bls := &Balances{
	// 	// Status: rawProductCoins.Status,
	// }

	return &rawBalances, nil
}
func (as *apiService) Balance(bq BalanceRequest) (*Balance, error) {
	params := make(map[string]interface{})
	params["currency"] = bq.Currency
	params["type"] = bq.Type

	res, err := as.request("POST", "/wallet/balance", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Balance.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	// rawBalance := struct {
	// 	Status uint32 `json:"status"`
	// 	Data   struct {
	// 		Type      uint32  `json:"type"`
	// 		Currency  uint32  `json:"currency"`
	// 		Quantity  float64 `json:"quantity"`
	// 		Withdraw  float64 `json:"withdraw"`
	// 		Transfer  float64 `json:"transfer"`
	// 		Deposit   float64 `json:"deposit"`
	// 		Lock      uint32  `json:"lock"`
	// 		Available uint32  `json:"available"`
	// 		Status    uint32  `json:"status"`
	// 	}
	// }{}
	rawBalance := Balance{}

	if err := json.Unmarshal(textRes, &rawBalance); err != nil {
		return nil, errors.Wrap(err, "rawBalance unmarshal failed")
	}

	// bl := &Balance{
	// 	// Status: rawProductCoins.Status,
	// }

	return &rawBalance, nil
}

func (as *apiService) Withdraw(wq WithdrawRequest) (*Withdraw, error) {
	params := make(map[string]interface{})
	params["currency"] = wq.Currency
	params["address"] = wq.Address
	params["amount"] = wq.Amount

	res, err := as.request("POST", "/wallet/withdraw", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from withdraw.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	// rawWithdraw := struct {
	// 	Status uint32 `json:"status"`
	// 	Data   struct {
	// 		Currency uint32  `json:"currency"`
	// 		Quantity float64 `json:"quantity"`
	// 		Withdraw float64 `json:"withdraw"`
	// 		Transfer float64 `json:"transfer"`
	// 		Deposit  float64 `json:"deposit"`
	// 		OrderID  string  `json:"orderId"`
	// 	}
	// }{}
	rawWithdraw := Withdraw{}

	if err := json.Unmarshal(textRes, &rawWithdraw); err != nil {
		return nil, errors.Wrap(err, "rawWithdraw unmarshal failed")
	}

	// wd := &Withdraw{
	// 	// Status: rawProductCoins.Status,
	// }

	return &rawWithdraw, nil
}
func (as *apiService) Switch(wq SwitchRequest) (*Switch, error) {
	params := make(map[string]interface{})
	params["currency"] = wq.Currency
	params["direction"] = wq.Direction
	params["amount"] = wq.Amount

	res, err := as.request("POST", "/wallet/switch", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Switch.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSwitch := Switch{}

	if err := json.Unmarshal(textRes, &rawSwitch); err != nil {
		return nil, errors.Wrap(err, "rawSwitch unmarshal failed")
	}

	return &rawSwitch, nil
}
