package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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

	return &Page{Title: fileName, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	//logger aufsetzten
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
