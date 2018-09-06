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

	// rawProductCoins := struct {
	// 	Status uint32 `json:"status"`
	// 	Data   struct {
	// 		List []struct {
	// 			ID            int64   `json:"id"`
	// 			Code          string  `json:"code"`
	// 			Symbol        string  `json:"symbol"`
	// 			Family        string  `json:"family"`
	// 			MinTrade      float64 `json:"minTrade"`
	// 			MinCash       float64 `json:"minCash"`
	// 			Fee           float64 `json:"fee"`
	// 			FeeType       string  `json:"feeType"`
	// 			Digits        uint32  `json:"digits"`
	// 			ShowDigits    uint32  `json:"showDigits"`
	// 			Confirmations uint32  `json:"confirmations"`
	// 			Op            uint32  `json:"op"`
	// 			Types         uint32  `json:"types"`
	// 			Exchange      struct {
	// 				Max uint32   `json:"max"`
	// 				Min float64  `json:"min"`
	// 				To  []uint32 `json:"to"`
	// 			}
	// 		}
	// 	}
	// }{}
	rawProductCoins := ProductCoins{}
	if err := json.Unmarshal(textRes, &rawProductCoins); err != nil {
		return nil, errors.Wrap(err, "rawProductCoins unmarshal failed")
	}

	// pc := &ProductCoins{
	// 	Status: rawProductCoins.Status,
	// 	Data: ProductCoinsData{
	// 		List: []*ProductCoinsList{},
	// 	},
	// }
	// for _, b := range rawProductCoins.Data.List {
	// 	// free, err := floatFromString(b.AvailableBalance)
	// 	// if err != nil {
	// 	// 	level.Error(as.Logger).Log("wsUnmarshal", err, "body", b.AvailableBalance)
	// 	// 	return
	// 	// }
	// 	// locked, err := floatFromString(b.Locked)
	// 	// if err != nil {
	// 	// 	level.Error(as.Logger).Log("wsUnmarshal", err, "body", b.Locked)
	// 	// 	return
	// 	// }
	// 	pc.Data.List = append(pc.Data.List, &ProductCoinsList{
	// 		ID:            b.ID,
	// 		Code:          b.Code,
	// 		Symbol:        b.Symbol,
	// 		Family:        b.Family,
	// 		MinTrade:      b.MinTrade,
	// 		MinCash:       b.MinCash,
	// 		Fee:           b.Fee,
	// 		FeeType:       b.FeeType,
	// 		Digits:        b.Digits,
	// 		ShowDigits:    b.ShowDigits,
	// 		Confirmations: b.Confirmations,
	// 		Op:            b.Op,
	// 		Types:         b.Types,
	// 		Exchange: ProductCoinsExchange{
	// 			Max: b.Exchange.Max,
	// 			Min: b.Exchange.Min,
	// 			To:  b.Exchange.To,
	// 		},
	// 	})
	// }

	return &rawProductCoins, nil
}
