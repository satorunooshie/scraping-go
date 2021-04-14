package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
)

func main() {
	wUrl := "http://machicon.jp/areas/tokyo"
	doc, err := goquery.NewDocument(wUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	title := doc.Find("title").Text()
	fmt.Println(title)

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		/* 相対パス
		a := s.Find("a")

		h, _ := a.Last().Attr("href")

		u.Path = h
		fmt.Println(u.String())
		*/
		a := s.Find("a")
		h, _ := a.Attr("href")
		bUrl, _ := url.Parse(wUrl)
		fullUrl, err := toAbsUrl(bUrl, h)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(fullUrl)
	})
	fmt.Println(doc.Find(".eventItem-bookmark-title").Text())
}
func toAbsUrl(baseUrl *url.URL, webUrl string) (string, error) {
	relUrl, err := url.Parse(webUrl)
	if err != nil {
		return "", err
	}
	absUrl := baseUrl.ResolveReference(relUrl)
	return absUrl.String(), nil
}
