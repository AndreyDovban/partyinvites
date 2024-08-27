package main

import (
	"html/template"
	"log"
	"net/http"
)

type Rsvp struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	log.Println("http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for i, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err != nil {
			log.Panicln(err.Error())
		} else {
			templates[name] = t
			log.Println("Loaded template", i, name)
		}
	}
}

func welcomHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			Rsvp:   &Rsvp{},
			Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		responseData := Rsvp{
			Name:       request.Form["name"][0],
			Email:      request.Form["email"][0],
			Phone:      request.Form["phone"][0],
			WillAttend: request.Form["willattend"][0] == "true",
		}

		responses = append(responses, &responseData)

		if responseData.WillAttend {
			templates["thanks"].Execute(writer, responseData.Name)
		} else {
			templates["sorry"].Execute(writer, responseData.Name)
		}
	}
}
