package images

import (
	"bytes"
	"encoding/json"
	"gitlab.com/chertokdmitry/surfcams/src/env"
	"gitlab.com/chertokdmitry/surfcams/src/message"
	"gitlab.com/chertokdmitry/surfcams/src/utils/logger"
	"net/http"
)

func Insert(CameraId int64, name string) {
	sub := &Image{0, CameraId, name}
	jsonReq, err := json.Marshal(sub)
	resp, err := http.Post(env.API_HOST+"images", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Error(message.ErrHttpPost, err)
	}

	defer resp.Body.Close()
}
