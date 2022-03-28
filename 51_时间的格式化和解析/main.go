package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	//这是一个遵循 RFC3339，并使用对应的布局
	//（layout）常量进行格式化的基本例子。
	t := time.Now()
	p(t.Format(time.RFC3339))
	
	//时间解析使用与 Format 相同的布局值。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	//format和Parse使用基于例子的布局来决定日期格式，
	//一般你只要使用time包中提供的布局常量就行了，
	//但是你也可以实现自定义布局.
	//布局时间必须使用 Mon Jan 2 15:04:05 MST 2006 
	//的格式，来指定格式化/解析给定时间/字符串的布局。 
	//时间一定要遵循：2006为年，15为小时，Monday代表星期几等规则。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

}
