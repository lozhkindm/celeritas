package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CsrfToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (r *Render) Page(w http.ResponseWriter, req *http.Request, view string, vars, data interface{}) error {
	switch strings.ToLower(r.Renderer) {
	case "go":
		return r.goPage(w, req, view, data)
	case "jet":
	}
	return nil
}

func (r *Render) goPage(w http.ResponseWriter, req *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", r.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	if err := tmpl.Execute(w, td); err != nil {
		return err
	}

	return nil
}
