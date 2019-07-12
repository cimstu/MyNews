package main

import (
	"bug"
	"time"
	"webpage"
)

func zhibo8Spider()  {
	zhibo8Searcher := &bug.Seacher{NetSet:"https://news.zhibo8.cc/zuqiu/more.htm?label=%E8%A5%BF%E7%94%B2"}
	zhibo8Searcher.KeyWords = append(zhibo8Searcher.KeyWords, "内马尔")
	webpage.MakePage(zhibo8Searcher.Run(), 10)

	time.Sleep(10*time.Minute)
	go zhibo8Spider()
}
