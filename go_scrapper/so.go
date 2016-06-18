// by re9ulus 11.06.2016

/*
TODO:
1. Separate common logic in getSearch and getAnsById functions
2. Add rating fields to questions and answers
3. Add while loop for user requests
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	APIUrl = "https://api.stackexchange.com//2.2/"
	// sort=[activity, votes, relevance]
	SearchUrl   = "search?order=desc&sort=relevance&site=stackoverflow&intitle=%s"
	AnsUrl      = "answers/%d?order=desc&sort=activity&site=stackoverflow&filter=!9YdnSM68f"
	QuestionUrl = "questions/%d?order=desc&sort=activity&site=stackoverflow"
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
	// ToDo: Should return Answer obj, not string
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

func getQuestionById(id int) (string, bool) {
	// ToDo: Should return Question obj, not string
	body := APIRequest(fmt.Sprintf(QuestionUrl, id))
	defer body.Close()

	decoder := json.NewDecoder(body)
	var data QuestionList
	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println("error occured")
	}

	if len(data.Questions) > 0 {
		return html.UnescapeString(data.Questions[0].Title), true
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
	QuestionId int    `json:"question_id"`
	IsAnswered bool   `json:"is_answered"`
	Title      string `json:"title"`
}

type QuestionList struct {
	Questions []Question `json:"items"`
}

func main() {
	var question string
	flag.StringVar(&question, "q", "", "question to search for")
	flag.Parse()

	info := getSearch(question)

	fmt.Println("\nquestion found: ", len(info.Infos))
	fmt.Println("\n")

	if len(info.Infos) > 0 {
		for i, item := range info.Infos {
			if item.IsAnswered {
				fmt.Println(i, item.Title, "(", item.QuestionId, ")")
			}
		}

		fmt.Print("\n>> Select question: ")

		var i int
		_, err := fmt.Scanf("%d", &i)

		fmt.Println("\n===\n")
		checkError(err)
		if 0 <= i && i < len(info.Infos) {
			answer, _ := getAnsById(info.Infos[i].AcceptedAnswerId)
			fmt.Println(answer)
		} else {
			fmt.Println("Wrong input.")
		}
		fmt.Println("\n===\n")
	}
}
