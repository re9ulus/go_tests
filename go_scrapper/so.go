// by re9ulus 11.06.2016

/*
TODO:
1. Move SO requst logic to APIRequest function
*/

package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "io"
	"net/http"
	"net/url"
	"os"
)

const (
	APIUrl    = "https://api.stackexchange.com//2.2/"
	SearchUrl = "search?order=desc&sort=relevance&site=stackoverflow&intitle=%s"
	AnsUrl    = "answers/%d?order=desc&sort=activity&site=stackoverflow&filter=!9YdnSM68f"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func getAnsById(id int) (Ans, bool) {
	request := APIUrl + fmt.Sprintf(AnsUrl, id)
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

	if len(data.Answers) > 0 {
		return data.Answers[0], true
	} else {
		return nil, false
	}
}

func getSearch(question string) SearchInfoList {
	request := APIUrl + fmt.Sprintf(SearchUrl, url.QueryEscape(question))
	fmt.Println("request: ", request)

	resp, err := http.Get(request)
	checkError(err)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var data SearchInfoList
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println("error occured")
	}

	return data
}

type SearchInfo struct {
	IsAnswered       bool   `json:"is_answered"`
	QuestionId       int    `json:"question_id"`
	AcceptedAnswerId int    `json:"accepted_answer_id"`
	Title            string `json:"title"`
}

type SearchInfoList struct {
	Infos []SearchInfo `json:"items"`
}

type Ans struct {
	Body       string `json:"body_markdown"`
	AnswerId   int    `json:"answer_id"`
	IsAccepted bool   `json"is_accepted"`
}

type AnsList struct {
	Answers []Ans `json:"items"`
}

type Question struct {
}

func main() {
	question := "python compare lists"
	info := getSearch(question)

	fmt.Println("data length: ", len(info.Infos))
	for i, item := range info.Infos {
		fmt.Println(i, item.Title, item.AcceptedAnswerId)
		getAnsById(item.AcceptedAnswerId)
		// getAnsById(33448575)
	}
}
