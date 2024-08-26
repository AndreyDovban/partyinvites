package main

import (
	"html/template"
)

type Rsvp struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()
}

func loadTemplates() {
	// TODO - load templates here
}
