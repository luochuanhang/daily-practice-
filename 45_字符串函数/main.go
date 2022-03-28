package main

import (
	"fmt"
	s "strings"
)

//我们给 fmt.Println 一个较短的别名， 因为我们随后会大量的使用它。
var p = fmt.Println

func main() {
	//是否包含子串
	p("Contains:  ", s.Contains("test", "es"))
	//包含几个子串
	p("Count:     ", s.Count("testt", "t"))
	//是否以子串为开头
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	//是否以子串为结尾
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	//子串所在的下标
	p("Index:     ", s.Index("test", "e"))
	//用第二个分隔符连接字符串
	p("Join:      ", s.Join([]string{"aa", "cc", "b"}, "-"))
	//将字符串重复5次
	p("Repeat:    ", s.Repeat("a", 5))
	//用后面的字符替换前面的，如果为-1 替换所有
	p("Replace:   ", s.Replace("fooooo", "o", "0", -1))
	//用后面的字符替换前面的，如果为正数替换数量个数
	p("Replace:   ", s.Replace("fooooo", "o", "1", 3))
	//将字符串按照后面的分隔符分开
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	//转换为小写
	p("ToLower:   ", s.ToLower("TEST"))
	//转换为大写
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	//长度
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
