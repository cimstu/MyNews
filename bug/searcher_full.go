package bug

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type SearcherFull struct {
	NetSet string
	KeyWords []string
}

func (s *SearcherFull)Run() (*Article) {
	host := "https://www.zhibo8.cc"
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
			sel.Find("b").Each(func(j int, sel *goquery.Selection) {
				find_content = true

				if found {
					return
				}

				title := sel.Text()
				u, _ := sel.Find("[href]").Attr("href")

				//fmt.Println(title)

				for _, word := range s.KeyWords  {
					if strings.Contains(title, word) {
						found = true
						article = &Article{fmt.Sprintf("%s-%s", date, title), fmt.Sprintf("%s%s", host, u), date}
						break
					}
				}
			})
		})

		if !find_content {
			sel.Find("span").Each(func(i int, selection *goquery.Selection) {
				if found {
					return
				}

				c := selection.Nodes[0].FirstChild
				n := 0

				title := ""
				u := ""

				for ; c != nil; c = c.NextSibling {
					if n % 2 == 0 {
						title = c.Data
					} else {
						u = c.Attr[0].Val
						//fmt.Println(fmt.Sprintf("%s-%s", title, u))

						for _, word := range s.KeyWords  {
							if strings.Contains(title, word) {
								found = true

								title = strings.TrimPrefix(title, " | ")
								article = &Article{fmt.Sprintf("%s-%s", date, title), fmt.Sprintf("%s%s", host, u), date}

								break
							}
						}

						if found {
							break
						}
					}

					n += 1
				}
			})
		}

	})


	return article
}