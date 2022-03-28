package main

import (
	"crypto/sha1"
	"fmt"
)

//SHA1 散列（hash)经常用于生成二进制文件或者文本块的短标识
//Go 在多个 crypto/* 包中实现了一系列散列函数
func main() {
	s := "sha1 this string"
	//产生一个散列值的方式是 sha1.New()
	//sha1.Write(bytes)，然后sha1.Sum([]byte{})
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，
	// 需要使用 []byte(s) 将其强制转换成字节数组
	h.Write([]byte(s))
	//Sum 得到最终的散列值的字符切片
	//Sum 接收一个参数， 可以用来给现有的字符切片追加
	//额外的字节切片：但是一般都不需要这样做。
	bs := h.Sum(nil)

	//SHA1值经常以16进制输出，例如在gitcommit中
	//我们这里也使用%x来将散列结果格式化为16进制字符串。

	fmt.Println(bs)
	fmt.Printf("%x\n", bs)

	//你可以使用和上面相似的方式来计算其他形式的散列值。 例如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。
}
