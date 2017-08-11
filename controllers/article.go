package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sagarmj76/basicgoapi/model"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := model.Articles{
		model.Article{Title: "Hello", Desc: "Article Desc", Content: "Article Content"},
		model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
}
