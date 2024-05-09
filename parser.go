package main

import (
	"log"
	"strings"
	"time"
)

func TitleAuthor(line string) (title, author string) {
	// Try parsing on parenthesis basis

	author_start := strings.LastIndex(line, "(")
	pdf_author_start := strings.LastIndex(line, "-")

	if author_start == -1 || pdf_author_start > author_start {
		// This highlight is from a PDF
		// second condition is for a PDF file case like: Title = Golang (3rd Ed.) - Jack Sparrow
		author = strings.Trim(line[pdf_author_start+1:len(line)-1], "\r\t\n ")
		title = strings.Trim(line[:pdf_author_start], "\r\t\n ")

	} else {

		author = strings.Trim(line[author_start+1:len(line)-1], "\r\t\n ")
		title = strings.Trim(line[:author_start], "\r\t\n ")
	}

	return title, author
}

func Parser(wholeText string) map[string]Bookdata {
	ParsedData := make(map[string]Bookdata)
	var title, author string
	var quote []string
	var structuredQuote Quote

	state := EMPTY

	for _, line := range strings.Split(wholeText, "\n") {
		line = strings.Trim(line, "\r\t\n ")
		if line == "" || line == "\n" || line == " " || line == "\r" {
			continue
		}

		// Mark state for new highlight, and move to next line
		equal_literal := "=========="
		if line == equal_literal {
			if state == LOC_TIME {
				full_quote := strings.Join(quote, "\n")
				full_quote = strings.Trim(full_quote, " \t\n\r")
				if full_quote == "" {
					// skip if empty quote
					state = EMPTY
					title = ""
					author = ""
					quote = nil

					continue
				}
				structuredQuote.Text = full_quote

				if val, ok := ParsedData[title]; ok {
					val.QuoteList = append(val.QuoteList, structuredQuote)
					ParsedData[title] = val
				} else {
					var data Bookdata
					data.Title = title
					data.Author = author
					data.QuoteList = append(data.QuoteList, structuredQuote)
					ParsedData[title] = data
				}
			}
			state = EMPTY
			title = ""
			author = ""
			quote = nil
			continue
		}

		// Here the string encountered should be non-equal-to and non-newline strings.
		switch state {
		case EMPTY:
			title, author = TitleAuthor(line)
			state = TITLE_AUTHOR
		case TITLE_AUTHOR:
			structuredQuote.Location = "0"
			_, time_string, _ := strings.Cut(line, " Added on ")
			// Maybe, if we need robustness here, we need to try multiple time formats for a succesful parse
			// Buse i have seen the below 2 patterns in highlights, currently using whats presnt in my own highlights file
			// timeAdded, err := time.Parse("Monday, 2 January 2006 15:04:05 PM", time_string)
			timeAdded, err := time.Parse("Monday, January 2, 2006 15:04:05 PM", time_string)
			structuredQuote.TimeAdded = timeAdded
			if err != nil {
				log.Printf("Error parsing time: %v", err)
			}
			state = LOC_TIME
		case LOC_TIME:
			quote = append(quote, line)
		}

	}
	return ParsedData
}
