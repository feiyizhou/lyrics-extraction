package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lyrics-extraction/common"
	"lyrics-extraction/modules/trasfer"
	"lyrics-extraction/utils"
	"net/http"
	"os"
	"time"
)

type transferService struct{}

func NewTransferService() *transferService {
	return &transferService{}
}

//https://raasr.xfyun.cn/v2/api/upload?
//duration=200
//&signa=Je5YsBvPcsbB4qy8Qvzd367fiv0%3D
//&fileName=%E9%98%B3%E5%85%89%E6%80%BB%E5%9C%A8%E9%A3%8E%E9%9B%A8%E5%90%8E.speex-wb
//&fileSize=11895
//&sysDicts=uncivilizedLanguage
//&appId=3e79d91c
//&ts=1662101767
func (ts *transferService) Upload(filePath string, duration int) (*trasfer.UploadResponse, error) {
	nowTime := time.Now().Unix()
	signa := utils.GetXFSignAStr(nowTime)
	fileName, fileSize, err := utils.GetFileInfo(filePath)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s?appId=%s&signa=%s&ts=%d&fileSize=%d&fileName=%s&sysDicts=%s"+
		"&duration=%d", common.BaseURL, common.Upload, common.AppID, signa, nowTime, fileSize,
		fileName, "uncivilizedLanguage", duration)
	fmt.Println(url)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := bufio.NewReader(file)
	httpReq, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()
	resBody, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	var response trasfer.UploadResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return nil, nil
}

//https://raasr.xfyun.cn/v2/api/getResult?
//signa=Wv23VLOg%2F6sQ1BDx4DKnnxtgiwQ%3D
//&orderId=DKHJQ2022090217220902175209AAEBD000015
//&appId=3e79d91c
//&resultType=predict
//&ts=1662112340
func (ts *transferService) GetResult(orderID string) (*trasfer.GetResultResponse, error) {
	nowTime := time.Now().Unix()
	signa := utils.GetXFSignAStr(nowTime)
	url := fmt.Sprintf("%s%s?signa=%s&orderId=%s&appId=%s&resultType=transfer&ts=%d",
		common.BaseURL, common.GetResult, signa, orderID, common.AppID, nowTime)
	httpReq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()
	resBody, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	var response trasfer.GetResultResponse
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return nil, nil
}
