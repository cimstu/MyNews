package bug

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Seacher struct {
	NetSet   string
	KeyWords []string
}

type Article struct {
	Title string
	Url   string
	T	  string
}

func (s *Seacher) Run() []Article {
	resp, _ := http.Get(s.NetSet)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("ERROR:Seacher got page failed url:%s, code:%d", s.NetSet, resp.StatusCode)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("ERROR:Seacher got page html contents url:%s, failed:%s", s.NetSet, err)
		return nil
	}

	var articles []Article

	doc.Find(".articleTitle").Each(func(i int, sel *goquery.Selection) {
		title := sel.Find("a").Text()
		for _, key := range s.KeyWords {
			if strings.Index(title, key) >= 0 {
				u, _ := sel.Find("[href]").Attr("href")
				date := sel.Parent().Find(".postTime").Text()

				article := Article{title, u, date}

				articles = append(articles, article)
			}
		}
	})

	return articles
}
