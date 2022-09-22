package apis

import "lyrics-extraction/modules/trasfer"

type TransferApi interface {
	// 上传文件
	Upload(filePath string) (*trasfer.UploadResponse, error)
	// 获取转写结果
	GetResult(orderID string) (*trasfer.GetResultResponse, error)
}
