package images

type Image struct {
	Id       int64  `json:"id"`
	CameraId int64  `json:"camera_id"`
	Name     string `json:"name"`
}
