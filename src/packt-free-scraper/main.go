package main

import (
	"bytes"
	"net/http"
	"strings"
    "fmt"
    "log"
    "io/ioutil"
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

    text := doc.Find("#deal-of-the-day > div > div > div.dotd-main-book-summary.float-left > div.dotd-title > h2").Text()
    text = strings.TrimSpace(text)
    return text
}


func postJson() {
    url := "http://httpbin.org/post"
    fmt.Println("URL:>", url)

    myStr := "The Book Name Is Good"
    _ = myStr

    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
    req, err := http.NewRequest("POSt", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // log vs panic?
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}