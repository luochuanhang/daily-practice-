package util

import (
	"crypto/md5"
	"encoding/hex"
)

//一般不会直接将上传的图片名暴露出来，
//因此我们对图片名进行MD5来达到这个效果
//md5编码
func EncodeMD5(value string) string {
	/*
		New返回一个新的哈希值。计算MD5校验和的哈希。
		Hash也实现编码。BinaryMarshaler和编码。
		BinaryUnmarshaler来编组和解组散列的内部状态。
	*/
	m := md5.New()
	//
	m.Write([]byte(value))
	//Sum将当前哈希添加到b并返回结果片。它不会改变底层的哈希状态。
	//EncodeToString返回src的十六进制编码。
	return hex.EncodeToString(m.Sum(nil))
}
