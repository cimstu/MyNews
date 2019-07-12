package bug

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Seacher struct {
	NetSet string
	KeyWords []string
}

type Article struct {
	Title string
	Url string
}

func (s* Seacher)Run()([]Article) {
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
	doc.Find(".articleTitle").Each(func(i int, sel* goquery.Selection) {
		t := sel.Find("a").Text()
		u, _ := sel.Find("[href]").Attr("href")
		articles = append(articles, Article{t, u})
	})

	var result []Article
	for _, article := range articles {
		for _, key := range s.KeyWords {
			if strings.Index(article.Title, key) >= 0 {
				result = append(result, article)
			}
		}
	}

	return result
}