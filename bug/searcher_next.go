package bug

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type SearcherNext struct {
	NetSet string
	KeyWords []string
}

func (s *SearcherNext)Run() (*Article) {
	var article *Article
	resp, _ := http.Get(s.NetSet)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("ERROR:Seacher full got page failed url:%s, code:%d", s.NetSet, resp.StatusCode)
		return article
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("ERROR:Seacher full got page html contents url:%s, failed:%s", s.NetSet, err)
		return article
	}

	found := false
	doc.Find(".box").Each(func(i int, sel *goquery.Selection) {
		if found {
			return
		}

		date := ""
		sel.Find(".titlebar").Each(func(i int, sel *goquery.Selection) {
			date = sel.Find("h2").Text()
			//fmt.Println(date)
		})

		find_content := false
		sel.Find(".content").Each(func(i int, sel *goquery.Selection) {
			sel.Find("li").Each(func(j int, sel *goquery.Selection) {
				find_content = true

				if found {
					return
				}

				tags_label, got := sel.Attr("label")
				if !got {
					return
				}

				tags := strings.Split(tags_label, ",")
				if tags[len(tags)-1] != "足球" {
					return
				}

				time_lable, got := sel.Attr("data-time")
				if !got {
					return
				}
				title :=strings.TrimSpace(sel.Find("b").Text())

				url, _ := sel.Find("a").Attr("href")

				for _, word := range s.KeyWords  {
					if strings.Contains(title, word) {
						found = true
						article = &Article{fmt.Sprintf("%s-%s", date, title), url,time_lable}
						break
					}
				}
			})
		})

	})


	return article
}
