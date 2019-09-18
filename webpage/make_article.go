package webpage

import (
	"bug"
	"fmt"
	"strings"
)


func MakePage(articles []bug.Article) string{
	var htmlElements string
	for _, article := range articles {

		if strings.Index(article.Url, "https") != 0 && strings.Index(article.Url, "http") != 0 {
			htmlElements += fmt.Sprintf("<p><a href=\"%s\" target=\"_blank\">%s</a></p>", "http:"+article.Url, article.Title)
		} else {
			htmlElements += fmt.Sprintf("<p><a href=\"%s\" target=\"_blank\">%s</a></p>", article.Url, article.Title)
		}
	}

	return fmt.Sprintf("<html><body>%s</body></html>", htmlElements)
}
