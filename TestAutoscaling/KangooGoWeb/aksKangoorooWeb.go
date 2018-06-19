package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received : %s", r.URL)
	fmt.Fprintf(w, "<html><head><title>Go title</title></head></head><body><h1>Hi there, I love %s!</h1></body></html>", r.URL.Path[1:])
}

func main() {
	fmt.Println("KangooGo web server started.")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
