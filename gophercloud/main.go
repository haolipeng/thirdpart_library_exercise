package main

import (
	"fmt"
	"time"
)

func main() {
	var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
	t := time.Now()                           //当前时间
	t.Unix()                                  //时间戳

	ts := t.Format(timeLayoutStr) //time转string
	fmt.Println(ts)
	st, _ := time.Parse(timeLayoutStr, ts) //string转time
	fmt.Println(st)
}
