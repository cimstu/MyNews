package data

import (
	"bug"
	"db"
	"encoding/json"
)


func GetKeyWords(k string)[]string {
	var words []string

	switch k {
	case "barca":
		words = append(words, "巴萨", "巴塞罗那")
		break
	case "real-madrid":
		words = append(words, "皇马", "皇家马德里")
		break
	case "man-city":
		words = append(words, "曼城", "曼彻斯特城")
		break
	case "man-unit":
		words = append(words, "曼联", "曼彻斯特联")
		break
	case "atm":
		words = append(words, "马竞", "马德里竞技")
		break
	case "spurs":
		words = append(words, "热刺")
		break
	case "neymar":
		words = append(words, "内马尔")
		break
	}

	return words
}

func GetArticle(k string, len int) []bug.Article {
	var articles []bug.Article

	urls := db.GetList(k, len)
	raw_articles := db.MGet(urls)
	for _, raw_article := range raw_articles {
		var article bug.Article
		err := json.Unmarshal([]byte(raw_article.(string)), &article)

		if err == nil {
			articles = append(articles, article)
		}
	}

	return articles
}