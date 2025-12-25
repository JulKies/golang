package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	filename := p.Title + ".txt"
	if len(filename) <= 4 {
		return
	}

	err := os.WriteFile(filename, p.Body, 0600)
	if err != nil {
		log.Println("Error writing file $v", filename)
	}
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)

	if err != nil {
		log.Println("Error reading file $v", fileName)
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate("view", w, page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	page.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate("edit", w, page)
}

func renderTemplate(title string, w http.ResponseWriter, p *Page) {
	err := templates.ExecuteTemplate(w, title+".html", p)
	if err != nil {
		log.Println("Cannot Parse template '$s'", title)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	//logger aufsetzten
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
