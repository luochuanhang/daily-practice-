package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5加密
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	//EncodeToString返回src的十六进制编码。
	return hex.EncodeToString(m.Sum(nil))
}
