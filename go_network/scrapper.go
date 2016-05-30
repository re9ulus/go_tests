// by re9ulus 30.05.2016

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func main() {
	url := "http://www.tornadoweb.org/"
	fmt.Println(getPage(url))
}
