package trasfer

type UploadRequest struct {
	FileName string `json:""`
	FileSize int64  `json:"fileSize"`
	Duration int64  `json:"duration"`
	Language string `json:"language"`
}

type UploadResponse struct {
	Code     string           `json:"code"`
	DescInfo string           `json:"descInfo"`
	Content  uploadResContent `json:"content"`
}

type uploadResContent struct {
	OrderId          string `json:"orderId"`
	TaskEstimateTime int    `json:"taskEstimateTime"`
}
