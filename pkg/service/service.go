package service

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func JQGet(url string, docHandler func(*goquery.Document)) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	docHandler(doc)
}
