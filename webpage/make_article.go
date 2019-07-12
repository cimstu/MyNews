package webpage

import (
	"bug"
	"fmt"
	"strings"
)

var NewHtmlContents string

func MakePage(articles []bug.Article, top int) {
	var htmlElements string
	for i, article := range articles {
		if i >= top {
			break
		}

		if strings.Index(article.Url, "https") != 0 && strings.Index(article.Url, "http") != 0 {
			htmlElements += fmt.Sprintf("<p><a href=\"%s\">%s</a></p>", "http:"+ article.Url, article.Title)
		} else {
			htmlElements += fmt.Sprintf("<p><a href=\"%s\">%s</a></p>", article.Url, article.Title)
		}
	}

	NewHtmlContents = fmt.Sprintf("<html><body>%s</body></html>", htmlElements)
}
