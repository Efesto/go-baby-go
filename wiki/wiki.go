package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Page represents a web page
type Page struct {
	Title string
	Body  []byte
}

// Save saves a web page on disk
func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// non existent page
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := Page{Title: title, Body: []byte(body)}
	err := p.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

var templates = loadTemplates()

func loadTemplates() *template.Template {
	templatefiles := []string{"tmpl/edit.html", "tmpl/view.html"}
	templates := template.New("master")
	for _, file := range templatefiles {
		splits := strings.Split(file, "/")
		templatename := splits[len(splits)-1]

		tmpldata, _ := ioutil.ReadFile(file)
		tmpl, _ := template.New(templatename).Parse(string(replaceLinks(tmpldata)))

		templates.AddParseTree(templatename, tmpl.Tree)
	}

	return templates
}

var interpolationregex = regexp.MustCompile("\\[([a-zA-Z0-9]+)\\]")

func replaceLinks(tmpldata []byte) []byte {
	return interpolationregex.ReplaceAllFunc(tmpldata, func(data []byte) []byte {
		group := interpolationregex.ReplaceAllString(string(data), `$1`)
		newgroup := "<a href='/view/" + group + "'>" + group + "</a>"
		return []byte(newgroup)
	})
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
