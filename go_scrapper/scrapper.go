// by re9ulus 30.05.2016

package main

import (
	"flag"
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
	filename = strings.Replace(filename, "https://", "", -1)
	filename = strings.Replace(filename, "/", "_", -1)
	return filename
}

func getUrlsFromPage(url string, newLinkChan chan string, saveFolder string) {
	fmt.Println("Url to process: ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("error occured", "error: %s", err.Error())
		return
	}

	content, err := doc.Html()
	if err != nil {
		fmt.Println("error occured", "error: %s", err.Error())
		return
	}

	go savePage(saveFolder+urlToFilename(url), content)

	checkError(err)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && checkLink(link) {
			newLinkChan <- link
		}
	})
}

func isFolderExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func mkDir(path string) error {
	isExists, err := isFolderExists(path)
	checkError(err)
	if !isExists {
		return os.MkdirAll(path, 0777)
	}
	return nil
}

func main() {
	var urlRoot, saveFolder string
	flag.StringVar(&urlRoot, "root", "", "page root")
	flag.StringVar(&saveFolder, "save_folder", "F:/test/", "folder to save pages")
	flag.Parse()
	visitedLinks := make(map[string]bool)
	visitedLinks[""] = true
	defaultUrl := ""
	newLinkChan := make(chan string)

	go getUrlsFromPage(urlRoot+defaultUrl, newLinkChan, saveFolder)

	stopCounter := 0
	for stopCounter < 3 {
		select {
		case newLink := <-newLinkChan:
			stopCounter = 0
			_, ok := visitedLinks[newLink]
			if !ok {
				visitedLinks[newLink] = true
				go getUrlsFromPage(urlRoot+newLink, newLinkChan, saveFolder)
			}
		case <-time.After(time.Second):
			fmt.Println("exit?")
			stopCounter += 1
		}
	}

	for key, _ := range visitedLinks {
		fmt.Println(key)
	}

	fmt.Println("Pages found:", len(visitedLinks))
	fmt.Println("Exit from main")
}
