// by re9ulus 30.05.2016

/*
TODO:
1. Pass url as command arg
2. Save '.html' pages to disc
3. Pass path to 'save' folder as command arg
*/

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func checkLink(url string) bool {
	return !(strings.HasPrefix(url, "http") || strings.Contains(url, "#"))
}

func savePage(filename, html string) {
	fmt.Println("Write ", filename)
	ioutil.WriteFile(filename, []byte(html), 0222)
}

func urlToFilename(url string) string {
	filename := strings.Replace(url, "http://", "", -1)
	filename = strings.Replace(filename, "/", "_", -1)
	// filename = strings.Replace(filename, ":", "", -1)
	return filename
}

func getUrlsFromPage(url string, newLinkChan chan string) {
	fmt.Println("Url to process: ", url)
	doc, err := goquery.NewDocument(url)
	checkError(err)

	// fmt.Println(doc.Html())
	if strings.HasSuffix(url, ".html") {
		content, err := doc.Html()
		checkError(err)

		go savePage("F:/test/"+urlToFilename(url), content)
	}

	checkError(err)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && checkLink(link) {
			newLinkChan <- link
		}
	})
}

func main() {
	urlRoot := "http://www.tornadoweb.org/en/stable/"
	visitedLinks := make(map[string]bool)
	visitedLinks[""] = true
	defaultUrl := ""
	newLinkChan := make(chan string)

	go getUrlsFromPage(urlRoot+defaultUrl, newLinkChan)

	stopCounter := 0
	for stopCounter < 3 {
		select {
		case newLink := <-newLinkChan:
			stopCounter = 0
			_, ok := visitedLinks[newLink]
			if !ok {
				visitedLinks[newLink] = true
				go getUrlsFromPage(urlRoot+newLink, newLinkChan)
			}
		case <-time.After(time.Second):
			fmt.Println("exit?")
			stopCounter += 1
		}
	}

	for key, _ := range visitedLinks {
		fmt.Println(key)
	}
	fmt.Println(len(visitedLinks))

	fmt.Println("Exit from main")
}
