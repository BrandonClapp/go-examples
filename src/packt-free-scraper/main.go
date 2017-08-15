package main

import (
	"strings"
    "fmt"
    "log"
    goquery "github.com/PuerkitoBio/goquery"
)

var url = "https://www.packtpub.com/packt/offers/free-learning"

func main() {
    doc, err := goquery.NewDocument(url)
    if err != nil {
      log.Fatal(err)
    }

    title := getBookTitle(doc)
    fmt.Println(title)
}

func getBookTitle(doc *goquery.Document) string {

    // todo: error checking?

    text := doc.Find("#deal-of-the-day > div > div > div.dotd-main-book-summary.float-left > div.dotd-title > h2").Text()
    text = strings.TrimSpace(text)
    return text
}