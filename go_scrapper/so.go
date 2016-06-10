// by re9ulus 11.06.2016

package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	SOUrl = "https://api.stackexchange.com//2.2/search?order=desc&sort=relevance&site=stackoverflow&intitle="
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func getSOAnswers(question string) AnsList {
	request := SOUrl + url.QueryEscape(question)
	fmt.Println("request: ", request)

	resp, err := http.Get(request)
	checkError(err)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var data AnsList
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println("error occured")
	}

	return data
}

type AnsInfo struct {
	IsAnswered       bool   `json:"is_answered"`
	QuestionId       int    `json:"question_id"`
	AcceptedAnswerId int    `json:"accepted_answer_id"`
	Title            string `json:"title"`
}

type AnsList struct {
	Infos []AnsInfo `json:"items"`
}

type Ans struct {
}

type Question struct {
}

func main() {
	question := "python compare lists"
	info := getSOAnswers(question)

	fmt.Println("data length: ", len(info.Infos))
	for i, item := range info.Infos {
		fmt.Println(i, item.Title)
	}

}
