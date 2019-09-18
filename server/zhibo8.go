package main

import (
	"bug"
	"data"
	"db"
	"encoding/json"
	"time"
)

func zhibo8Spider() {
	var key = "neymar"
	zhibo8Searcher := &bug.Seacher{NetSet: "https://news.zhibo8.cc/zuqiu/more.htm?label=%E8%A5%BF%E7%94%B2"}
	zhibo8Searcher.KeyWords = append(zhibo8Searcher.KeyWords, data.GetKeyWords(key)...)
	as := zhibo8Searcher.Run()
	for i := len(as)-1; i >= 0; i-- {
		db.LPush(key+"-news", as[i].Url)
		raw, _ := json.Marshal(as[i])
		db.Set(as[i].Url, string(raw))
	}

	time.Sleep(10 * time.Minute)
	go zhibo8Spider()
}

func zhibo8SpiderFull() {
	keys := []string {"barca", "man-city", "man-unit", "atm", "spurs", "real-madrid"}

	for _, key := range keys {
		zhibo8Searcher := &bug.SearcherFull{NetSet: "https://www.zhibo8.cc/zuqiu/luxiang.htm"}
		zhibo8Searcher.KeyWords = append(zhibo8Searcher.KeyWords, data.GetKeyWords(key)...)
		article := zhibo8Searcher.Run()
		if article != nil {
			db.LPush(key+"-full", article.Url)
			raw, _ := json.Marshal(*article)
			db.Set(article.Url, string(raw))
		}
	}

	time.Sleep(10 * time.Minute)
	go zhibo8SpiderFull()
}

func zhibo8SpiderNext() {
	keys := []string {"barca", "man-city", "man-unit", "atm", "spurs", "real-madrid"}

	for _, key := range keys {
		zhibo8Searcher := &bug.SearcherNext{NetSet: "https://www.zhibo8.cc"}
		zhibo8Searcher.KeyWords = append(zhibo8Searcher.KeyWords, data.GetKeyWords(key)...)
		article := zhibo8Searcher.Run()
		if article != nil {
			db.LPush(key+"-next", article.Url)
			raw, _ := json.Marshal(*article)
			db.Set(article.Url, string(raw))
		}
	}

	time.Sleep(10 * time.Minute)
	go zhibo8SpiderNext()
}