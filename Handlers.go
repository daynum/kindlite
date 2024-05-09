package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func indexPage(writer http.ResponseWriter, request *http.Request) {
	index_template_file := "index.html"
	input, err := os.ReadFile(index_template_file)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	front_page := string(input)
	var template = template.Must(template.New("front_page").Parse(front_page))
	template.Execute(writer, nil)
}

func processPage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing form data: %v\n", err)
	}

	// the map contains one single string at index 0, everything is there at 0 index in the slice.
	wholeText := request.Form["highlight-text"][0]

	// This slicing removes the first and last character which are '[' and ']'
	// IF text is empty it will PANIC here, [:-1]
	wholeText = wholeText[0 : len(wholeText)-1]
	email := strings.Join(request.Form["email"], "")
	book_map := Parser(wholeText)

	client := &FirebaseClient{credsFile: "/home/sharpfox/.keys/kindle-highlights-4116d-2d1ab71a0842.json"}
	client.Client()

	client.AddToEmail(email, book_map)
	//TODO: display a succesful response with a template, maybe include how many books/highlights were imported.
}

func getHighlights(writer http.ResponseWriter, request *http.Request) {
	template_file := "get_highlights.html"
	input, err := os.ReadFile(template_file)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	front_page := string(input)
	var template = template.Must(template.New("front_page").Parse(front_page))

	// Render template, diff for POST and GET
	if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			log.Fatalf("Error parsing form data: %v\n", err)
		}
		email := strings.Join(request.Form["email"], "")
		client := &FirebaseClient{credsFile: "/home/sharpfox/.keys/kindle-highlights-4116d-2d1ab71a0842.json"}
		client.Client()

		user_data := client.Read(email)

		template.Execute(writer, user_data)
	} else {
		template.Execute(writer, nil)
	}
}
