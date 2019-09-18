package main

import (
	"data"
	"fmt"
	"mylog"
	"net/http"
	"net/url"
	"webpage"
)

func zhibo8Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webpage.MakePage(data.GetArticle("neymar"+"-news", 15)))
}

func zhibo8PageFull(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	team := m["team"][0]
	fmt.Fprint(w, webpage.MakePage(data.GetArticle(team+"-full", 1)))
}

func zhibo8PageNext(w http.ResponseWriter, r *http.Request) {
	m, _ := url.ParseQuery(r.URL.RawQuery)
	team := m["team"][0]
	fmt.Fprint(w, webpage.MakePage(data.GetArticle(team+"-next", 1)))
}

func main() {
	mylog.InitLog("./myapp.log")

	go zhibo8Spider()
	go zhibo8SpiderFull()
	go zhibo8SpiderNext()

	http.HandleFunc("/zhibo8", zhibo8Page)
	http.HandleFunc("/zhibo8_full", zhibo8PageFull)
	http.HandleFunc("/zhibo8_next", zhibo8PageNext)

	http.ListenAndServe(":9595", nil)
}

