package main

import (
	"bytes"
	"net/http"
	"strings"
    "fmt"
    "log"
    goquery "github.com/PuerkitoBio/goquery"
)

var packtURL = "https://www.packtpub.com/packt/offers/free-learning"
var slackURL = "https://httpbin.org/post"

func main() {
    doc, err := goquery.NewDocument(packtURL)
    if err != nil {
      log.Fatal(err)
    }

    title := getBookTitle(doc)
    fmt.Println(title)
    postJSON(title)
}

func getBookTitle(doc *goquery.Document) string {
    text := doc.Find("#deal-of-the-day > div > div > div.dotd-main-book-summary.float-left > div.dotd-title > h2").Text()
    text = strings.TrimSpace(text)
    return text
}


func postJSON(bookTitle string) {
    payload := fmt.Sprintf(`{
        "channel": "#packt-free",
        "username": "packtpubfree",
        "text": "<%s|%s>",
        "icon_emoji": ":books:"
    }`, packtURL, bookTitle)

    var jsonStr = []byte(payload)
    req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonStr))
    if err != nil {
        fmt.Println("Could not complete POST request to slack.")
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // log vs panic?
        panic(err)
    }
    defer resp.Body.Close()
}