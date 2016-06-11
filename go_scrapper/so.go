// by re9ulus 11.06.2016

/*
TODO:
1. Separate common logic in getSearch and getAnsById functions
*/

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
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

func APIRequest(request string) io.ReadCloser {
	urlRequest := APIUrl + request
	fmt.Println("url request: ", urlRequest)
	resp, err := http.Get(urlRequest)
	checkError(err)
	return resp.Body
}

func getAnsById(id int) (string, bool) {
	body := APIRequest(fmt.Sprintf(AnsUrl, id))
	defer body.Close()

	decoder := json.NewDecoder(body)
	var data AnsList
	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println("error occured")
	}

	if len(data.Answers) > 0 {
		return html.UnescapeString(data.Answers[0].Body), true
	} else {
		return "", false
	}
}

func getSearch(question string) SearchInfoList {
	body := APIRequest(fmt.Sprintf(SearchUrl, url.QueryEscape(question)))
	defer body.Close()

	decoder := json.NewDecoder(body)
	var data SearchInfoList
	err := decoder.Decode(&data)

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
		fmt.Println(getAnsById(item.AcceptedAnswerId))
		fmt.Println("\n===\n")
	}
}
