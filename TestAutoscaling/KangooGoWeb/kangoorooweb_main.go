package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
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
	fmt.Printf("Request received : %s\n", r.URL)
	fmt.Fprintf(w, "<html><head><title>Go title</title></head></head><body>")
	fmt.Fprintf(w, "<h1>Hi there, I love %s!</h1>", r.URL.Path[1:])
	fmt.Fprintf(w, "<hr/>")
	fmt.Fprintf(w, "<b>OS</b>: [%s]<br/>", runtime.GOOS)
	fmt.Fprintf(w, "<b>ARCH</b>: [%s]<br/>", runtime.GOARCH)
	fmt.Fprintf(w, "<b>TIME</b>: [%s]<br/>", time.Now())
	fmt.Fprintf(w, "</body></html>")
}

func main() {
	var port = ":8080"
	fmt.Println("KangooGoWeb server starting ...")
	fmt.Printf("running on OS: [%s]   Architecture: [%s]\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Listening on [%s]\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}