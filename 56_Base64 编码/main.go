package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	//这是要编解码的字符串
	data := "abc123!?$*&()'-=@~"
	/*
		Go同时支持标准base64以及URL兼容base64。
		这是使用标准编码器进行编码的方法。
		编码器需要一个[]byte因此我们将string转换为
		该类型。
	*/
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	//解码可能会返回错误，如果不确定输入信息格式是否正确
	//需要进行错误检查
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	//使用 URL base64 格式进行编解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
