package main

import (
	"fmt"
	"net/http"
	"webpage"
)

func zhibo8Page(w http.ResponseWriter, r* http.Request) {
	fmt.Fprint(w, webpage.NewHtmlContents)
}

func main()  {
	go zhibo8Spider()

	http.HandleFunc("/zhibo8", zhibo8Page)
	http.ListenAndServe(":9595",  nil)
}
