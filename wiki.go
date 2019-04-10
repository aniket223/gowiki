package main

import ( //import different go packages

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct { //struct of pages
	Title string
	Body  []byte
}

func (p *Page) save() error { //function for saving pages
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //simple library writing to files
}

func loadPage(title string) (*Page, error) { //function for loading pages
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) // 2 arguments to check if error occurs
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) { //for viewing
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) { //for editing
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
