package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getTemplateString(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error opening template file: %v\n", err)
	}
	markup := string(input)
	return markup
}

func indexPage(writer http.ResponseWriter, request *http.Request) {
	markup := getTemplateString("index.html")
	var template = template.Must(template.New("front_page").Parse(markup))
	status := ""
	if request.Method == "POST" {
		err := request.ParseForm()

		fmt.Printf("\n\n%v\n\n", request.Form)

		if err != nil {
			log.Fatalf("Error parsing form data: %v\n", err)
			status = "FAILURE"
		} else {
			status = processPage(request.Form)
		}
	}
	template.Execute(writer, status)
}

func processPage(formData url.Values) string {
	// the map contains one single string at index 0, everything is there at 0 index in the slice.
	wholeText := formData["highlight-text"][0]

	// This slicing removes the first and last character which are '[' and ']'
	// IF text is empty it will PANIC here, [:-1]
	wholeText = wholeText[0 : len(wholeText)-1]
	email := strings.Join(formData["email"], "")
	book_map := Parser(wholeText)

	user_file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer user_file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Add User
	if _, err := user_file.WriteString(email + "\n"); err != nil {
		log.Fatal(err)
	}

	// write json to a file
	datafile, err := os.Create(email + ".json")
	if err != nil {
		log.Fatal(err)
		return "FAILURE"
	}
	json_data, err := json.MarshalIndent(book_map, "", "\t")
	if err != nil {
		log.Fatal(err)
		return "FAILURE"
	}
	_, err = datafile.Write(json_data)
	if err != nil {
		log.Fatal(err)
		return "FAILURE"
	}

	return "SUCCESS"
}

func getHighlights(writer http.ResponseWriter, request *http.Request) {
	markup := getTemplateString("get_highlights.html")
	var template = template.Must(template.New("front_page").Parse(markup))

	// Render template, diff for POST and GET
	if request.Method == "GET" {
		template.Execute(writer, nil)
	}
	// else POST
	err := request.ParseForm()
	if err != nil {
		log.Fatalf("Error parsing form data: %v\n", err)
	}
	email := strings.Join(request.Form["email"], "")
	datafile, err := os.ReadFile(email + ".json")
	if err != nil {
		log.Fatal(err)
	}
	user_data := &map[string]Bookdata{}
	err = json.Unmarshal(datafile, user_data)

	template.Execute(writer, user_data)
}
