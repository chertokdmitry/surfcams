package app

import (
	"gitlab.com/chertokdmitry/surfcams/src/domain/cameras"
	"gitlab.com/chertokdmitry/surfcams/src/domain/windy"
)

// run the app
func Run() {
	cameras := cameras.GetAll()

	for _, camera := range cameras {
		windy.Save(camera)
	}

	windy.SaveKurort()
}
