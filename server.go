//go:build ignore


package main


import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body string
}

const html string = "<!DOCTYPE html> <html><head><meta charset='utf-8' /><meta name='viewport' content='width=device-width'/><link src='../style.css' rel='stylesheet' type='text/css'>"

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("pages/"+filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("pages/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body[:])}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/wiki/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "An error occured when reading this file.")
	} else {
		page := html + "<h1>"+p.Title+"</h1><div>"+p.Body+"</div>"
		fmt.Fprintf(w, page)
	}
}

func viewRoot(w http.ResponseWriter, r *http.Request) {
	file_name := r.URL.Path[len("/"):]
	file, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Fprintf(w, "An error occured when reading this file.")
	} else {
		fmt.Fprintf(w, string(file))
	}
}

func main() {
	http.HandleFunc("/", viewRoot)
	http.HandleFunc("/wiki/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
