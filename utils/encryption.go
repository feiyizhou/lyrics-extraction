package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"lyrics-extraction/common"
)

func GetXFSignAStr(time int64) string {
	return HmacSHA1(common.SecretKey, MD5(GetBaseStr(time)))
}

func HmacSHA1(keyStr, value string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GetBaseStr(time int64) string {
	return fmt.Sprintf("%s%d", common.AppID, time)
}
