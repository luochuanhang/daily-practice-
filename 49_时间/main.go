package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	//获取当前时间
	now := time.Now()
	p(now)
	//通过提供年月日等信息，你可以构建一个 time。 时间总是与 Location 有关，也就是时区
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	p(then)
	//可以提取出时间的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	//比较两个时间，第一个时间是否在第二个时间前面
	p(then.Before(now))
	//比较两个时间，第一个时间是否在第二个时间之后
	p(then.After(now))
	////比较两个时间，第一个时间是否和第二个时间同一时刻
	p(then.Equal(now))

	//返回连个时间的时间差
	diff := now.Sub(then)

	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//可以将第一个时间移动第二个时间的时间长度
	p(then.Add(diff))
	p(then.Add(-diff))
}
