package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Go 提供内建的 JSON 编码解码（序列化反序列化）支持， 包括内建及自定义类型与 JSON 数据之间的转化。
type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	//JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	//可以给结构字段声明标签来自定义编码的 JSON 数据的键名
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	//将序列化的数据解析到dat键是string，值是interface
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//打印解析的值
	fmt.Println(dat)
	//为了使用解码 map 中的值，我们需要将他们进行适当的类型转换
	num := dat["num"].(float64)
	fmt.Println(num)

	//访问嵌套的值需要一系列的转化
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	//我们可以将json数据解码为自定义的数据类型
	//解码的时候不需要类型断言
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	//可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
