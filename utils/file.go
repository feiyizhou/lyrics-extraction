package utils

import (
	"fmt"
	"os"
)

func GetFileInfo(filePath string) (string, int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", 0, err
	}
	return fileInfo.Name(), fileInfo.Size(), err
}

func CheckFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("%s is not exist", filePath)
		}
		return false, err
	}
	return true, err
}
