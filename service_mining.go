package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) RobotClose(rr RobotCloseRequest) (*RobotClose, error) {
	params := make(map[string]interface{})
	params["type"] = rr.Type
	res, err := as.request("POST", "/mining/robotClose", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from RobotClose.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawRobotClose := RobotClose{}

	if err := json.Unmarshal(textRes, &rawRobotClose); err != nil {
		return nil, errors.Wrap(err, "rawRobotClose unmarshal failed")
	}

	return &rawRobotClose, nil
}
