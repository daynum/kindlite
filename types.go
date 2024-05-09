package main

import (
	"time"
)

// Need fields to be capital for firebase to access them and store in firestore
type Quote struct {
	Location  string    `firestore:"location,omitempty"`
	Text      string    `firestore:"text,omitempty"`
	TimeAdded time.Time `firestore:"time_added,omitempty"`
}

type Bookdata struct {
	Title     string  `firestore:"title,omitempty"`
	Author    string  `firestore:"author,omitempty"`
	QuoteList []Quote `firestore:"quote_list,omitempty"`
}

type ParseState int8

const (
	EMPTY ParseState = iota
	TITLE_AUTHOR
	LOC_TIME
	QUOTE
)
