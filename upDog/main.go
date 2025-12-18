package main

import (
	"log"
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

func main() {

	//logger aufsetzten
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	log.Println(string(p2.Body))
}
