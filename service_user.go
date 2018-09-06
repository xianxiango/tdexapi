package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) UserInfo() (*UserInfo, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/user/info", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from UserInfo.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawUserInfo := UserInfo{}

	if err := json.Unmarshal(textRes, &rawUserInfo); err != nil {
		return nil, errors.Wrap(err, "rawUserInfo unmarshal failed")
	}

	return &rawUserInfo, nil
}
