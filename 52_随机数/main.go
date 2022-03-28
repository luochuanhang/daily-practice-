package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//给定的种子是确定的，每次都会产生相同的随机数数字序列。
	//要产生不同的数字序列，需要给定一个不同的种子。

	//对于想要加密的随机数,使用此方法并不安全,应该使用crypto/rand。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
//如果使用相同种子生成的随机数生成器，会生成相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
