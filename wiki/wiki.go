package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, err
}

func main() {
	var title string = "TestPage"
	p1 := &Page{Title: title, Body: []byte("This is a simple Page.")}
	p1.save()
	p2, _ := loadPage(title)
	fmt.Println(string(p2.Body))
}
