package main

import ( //import different go packages

	"io/ioutil"
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

func main() {

}
