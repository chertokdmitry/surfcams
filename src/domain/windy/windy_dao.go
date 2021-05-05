package windy

import (
	"encoding/json"
	"gitlab.com/chertokdmitry/surfcams/src/domain/images"
	"gitlab.com/chertokdmitry/surfcams/src/env"
	"gitlab.com/chertokdmitry/surfcams/src/message"
	"gitlab.com/chertokdmitry/surfcams/src/utils/logger"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Result struct {
	Result Response `json:"result"`
}

type Response struct {
	Offset  int       `json:"offset"`
	Limit   int       `json:"limit"`
	Total   int       `json:"total"`
	Webcams []Webcams `json:"webcams"`
}

type Webcams struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Image  Image  `json:"image"`
}

type Image struct {
	Current  Current     `json:"current"`
	Sizes    interface{} `json:"sizes"`
	Daylight interface{} `json:"daylight"`
	Update   int         `json:"update"`
}

type Current struct {
	Icon      string `json:"icon"`
	Thumbnail string `json:"thumbnail"`
	Preview   string `json:"preview"`
	Toenail   string `json:"toenail"`
}

// get data from windy api
func Save(id int64) {
	idStr := strconv.FormatInt(id, 10)
	url := env.URL_PATH + idStr + env.URL_PARAMS + env.KEY
	now := strconv.FormatInt(time.Now().Unix(), 10)

	response, err := http.Get(url)

	if err != nil {
		logger.Error(message.ErrGetOWM, err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logger.Error(message.ErrReadAll, err)
	}

	var r Result
	json.Unmarshal(body, &r)

	if len(r.Result.Webcams) > 0 {
		image := r.Result.Webcams[0].Image.Current.Preview
		imagePath := "/var/www/surfweather.ru/webcams/" + idStr + ".jpg"

		err = SaveImage(imagePath, image)
		if err != nil {
			logger.Error(message.ErrSaveImage, err)
		}

		imagePath = "/var/www/surfweather.ru/webcams/" + idStr + "/" + now + ".jpg"

		err = SaveImage(imagePath, image)
		if err != nil {
			logger.Error(message.ErrSaveImage, err)
		}

		images.Insert(id, now)
	}
}

func SaveKurort() {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	imageKurort := "http://api.ipeye.ru/v1/stream/poster/1/a577c5effa7b4ca7a0cabba6741fbef6/pre.jpg"
	imagePathKurort := "/var/www/surfweather.ru/webcams/1000000000.jpg"

	err := SaveImage(imagePathKurort, imageKurort)
	if err != nil {
		logger.Error(message.ErrSaveImage, err)
	}

	imagePathKurort = "/var/www/surfweather.ru/webcams/1000000000/" + now + ".jpg"

	err = SaveImage(imagePathKurort, imageKurort)
	if err != nil {
		logger.Error(message.ErrSaveImage, err)
	}

	images.Insert(1000000000, now)
}

func SaveImage(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}
