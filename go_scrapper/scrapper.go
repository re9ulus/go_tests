// by re9ulus 30.05.2016

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func getPage(url string) string {
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	return string(body)
}

func getUrlsFromPage(url string) map[string]bool {
	doc, err := goquery.NewDocument(url)
	checkError(err)
	links := make(map[string]bool)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && !(strings.HasPrefix(link, "http")) {
			links[link] = true
		}
	})
	return links
}

func main() {
	urlRoot := "http://www.tornadoweb.org/"
	links := make(map[string]bool)
	visitedLinks := make(map[string]bool)

	links[""] = true

	for len(links) > 0 {
		for link, _ := range links {
			fmt.Println(len(visitedLinks))
			delete(links, link)
			visitedLinks[link] = true
			newLinks := getUrlsFromPage(urlRoot + link)
			for newLink, _ := range newLinks {
				_, ok := visitedLinks[newLink]
				if !ok {
					links[newLink] = true
				}
			}
		}
	}

	// for key, _ := range links {
	// 	fmt.Println(key)
	// }
}
