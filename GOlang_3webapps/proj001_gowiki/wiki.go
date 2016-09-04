package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

//data structure
type Page struct {
    Title string
    Body  []byte
}

//create a simple file page
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

//create a simple web server
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
		//create a simple file page
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))

		//create a simple web server
		fmt.Println("running the server at localhost:8080...")
		http.HandleFunc("/", handler)
	 	http.ListenAndServe(":8080", nil)


}
