package cameras

import (
	"encoding/json"
	"gitlab.com/chertokdmitry/surfcams/src/env"
	"gitlab.com/chertokdmitry/surfcams/src/message"
	"gitlab.com/chertokdmitry/surfcams/src/utils/logger"
	"io/ioutil"
	"net/http"
)

// get all cameras
func GetAll() []int64 {
	response, err := http.Get(env.API_HOST + "cameras/all/")
	if err != nil {
		logger.Error(message.ErrHttpGet, err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(message.ErrReadAll, err)
	}

	var res []int64
	json.Unmarshal(responseData, &res)

	return res
}
