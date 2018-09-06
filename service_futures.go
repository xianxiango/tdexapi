package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) FuturesOpen(wq FuturesOpenRequest) (*FuturesOpen, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["side"] = wq.Side
	params["scale"] = wq.Scale
	params["volume"] = wq.Volume
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["strategy"] = wq.Strategy
	params["better"] = wq.Better
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant

	res, err := as.request("POST", "/futures/open", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesOpen.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesOpen := FuturesOpen{}

	if err := json.Unmarshal(textRes, &rawFuturesOpen); err != nil {
		return nil, errors.Wrap(err, "rawFuturesOpen unmarshal failed")
	}

	return &rawFuturesOpen, nil
}
func (as *apiService) FuturesClose(wq []FuturesCloseRequest) (*FuturesClose, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/close", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesClose.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesClose := FuturesClose{}

	if err := json.Unmarshal(textRes, &rawFuturesClose); err != nil {
		return nil, errors.Wrap(err, "rawFuturesClose unmarshal failed")
	}

	return &rawFuturesClose, nil
}
func (as *apiService) FuturesCloseAll(wq []FuturesCloseAllRequest) (*FuturesCloseAll, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/closeAll", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesCloseAll.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesCloseAll := FuturesCloseAll{}

	if err := json.Unmarshal(textRes, &rawFuturesCloseAll); err != nil {
		return nil, errors.Wrap(err, "rawFuturesCloseAll unmarshal failed")
	}

	return &rawFuturesCloseAll, nil
}
func (as *apiService) FururesCancel(wq []FururesCancelRequest) (*FururesCancel, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/cancel", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FururesCancel.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFururesCancel := FururesCancel{}

	if err := json.Unmarshal(textRes, &rawFururesCancel); err != nil {
		return nil, errors.Wrap(err, "rawFururesCancel unmarshal failed")
	}

	return &rawFururesCancel, nil
}
func (as *apiService) FururesReplace(wq FururesReplaceRequest) (*FururesReplace, error) {
	params := make(map[string]interface{})
	params["id"] = wq.ID
	params["relative"] = wq.Relative
	params["cid"] = wq.Cid
	params["side"] = wq.Side
	params["scale"] = wq.Scale
	params["volume"] = wq.Volume
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["strategy"] = wq.Strategy
	params["better"] = wq.Better
	params["constant"] = wq.Constant

	res, err := as.request("POST", "/futures/replace", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FururesReplace.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFururesReplace := FururesReplace{}

	if err := json.Unmarshal(textRes, &rawFururesReplace); err != nil {
		return nil, errors.Wrap(err, "rawFururesReplace unmarshal failed")
	}

	return &rawFururesReplace, nil
}

func (as *apiService) Setsl(wq SetslRequest) (*Setsl, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["strategy"] = wq.Strategy
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["better"] = wq.Better

	res, err := as.request("POST", "/futures/setsl", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Setsl.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSetsl := Setsl{}

	if err := json.Unmarshal(textRes, &rawSetsl); err != nil {
		return nil, errors.Wrap(err, "rawSetsl unmarshal failed")
	}

	return &rawSetsl, nil
}
func (as *apiService) Settp(wq SettpRequest) (*Settp, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["strategy"] = wq.Strategy
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["better"] = wq.Better

	res, err := as.request("POST", "/futures/settp", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Settp.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSettp := Settp{}

	if err := json.Unmarshal(textRes, &rawSettp); err != nil {
		return nil, errors.Wrap(err, "rawSettp unmarshal failed")
	}

	return &rawSettp, nil
}

func (as *apiService) Merge(wq MergeRequest) (*Merge, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["list"] = wq.List

	res, err := as.request("POST", "/futures/merge", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Merge.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawMerge := Merge{}

	if err := json.Unmarshal(textRes, &rawMerge); err != nil {
		return nil, errors.Wrap(err, "rawMerge unmarshal failed")
	}

	return &rawMerge, nil
}
func (as *apiService) Split(wq SplitRequest) (*Split, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["volume"] = wq.Volume

	res, err := as.request("POST", "/futures/split", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Split.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSplit := Split{}

	if err := json.Unmarshal(textRes, &rawSplit); err != nil {
		return nil, errors.Wrap(err, "rawSplit unmarshal failed")
	}

	return &rawSplit, nil
}
func (as *apiService) Setup(wq SetupRequest) (*Setup, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["target"] = wq.Target

	res, err := as.request("POST", "/futures/setup", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Setup.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSetup := Setup{}

	if err := json.Unmarshal(textRes, &rawSetup); err != nil {
		return nil, errors.Wrap(err, "rawSetup unmarshal failed")
	}

	return &rawSetup, nil
}
func (as *apiService) SchemeSet(wq SchemeSetRequest) (*SchemeSet, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["options"] = wq.Options

	res, err := as.request("PUT", "/futures/scheme", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SchemeSet.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSchemeSet := SchemeSet{}

	if err := json.Unmarshal(textRes, &rawSchemeSet); err != nil {
		return nil, errors.Wrap(err, "rawSchemeSet unmarshal failed")
	}

	return &rawSchemeSet, nil
}
func (as *apiService) SchemeGet(wq SchemeGetRequest) (*SchemeGet, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid

	res, err := as.request("POST", "/futures/scheme", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from SchemeGet.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawSchemeGet := SchemeGet{}

	if err := json.Unmarshal(textRes, &rawSchemeGet); err != nil {
		return nil, errors.Wrap(err, "rawSchemeGet unmarshal failed")
	}

	return &rawSchemeGet, nil
}
func (as *apiService) GetOrders() (*GetOrders, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/futures/orders", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from GetOrders.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawGetOrders := GetOrders{}

	if err := json.Unmarshal(textRes, &rawGetOrders); err != nil {
		return nil, errors.Wrap(err, "rawGetOrders unmarshal failed")
	}

	return &rawGetOrders, nil
}
func (as *apiService) GetPosition() (*GetPosition, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/futures/position", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from GetPosition.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawGetPosition := GetPosition{}

	if err := json.Unmarshal(textRes, &rawGetPosition); err != nil {
		return nil, errors.Wrap(err, "rawGetPosition unmarshal failed")
	}

	return &rawGetPosition, nil
}
func (as *apiService) GetHistory(wq GetHistoryRequest) (*GetHistory, error) {
	params := make(map[string]interface{})
	params["page"] = wq.Page
	params["pageSize"] = wq.PageSize

	res, err := as.request("POST", "/futures/history", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from GetHistory.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawGetHistory := GetHistory{}

	if err := json.Unmarshal(textRes, &rawGetHistory); err != nil {
		return nil, errors.Wrap(err, "rawGetHistory unmarshal failed")
	}

	return &rawGetHistory, nil
}
func (as *apiService) GetContract(wq GetContractRequest) (*GetContract, error) {
	params := make(map[string]interface{})
	params["symbol"] = wq.Symbol

	res, err := as.request("POST", "/futures/contract", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from GetContract.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawGetContract := GetContract{}

	if err := json.Unmarshal(textRes, &rawGetContract); err != nil {
		return nil, errors.Wrap(err, "rawGetContract unmarshal failed")
	}

	return &rawGetContract, nil
}
