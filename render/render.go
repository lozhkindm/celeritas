package render

import (
	"errors"
	"fmt"
	"github.com/CloudyKit/jet/v6"
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
	JetViews   *jet.Set
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
		return r.jetPage(w, req, view, vars, data)
	default:

	}
	return errors.New("no renderer specified")
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

func (r *Render) jetPage(w http.ResponseWriter, req *http.Request, view string, vars, data interface{}) error {
	var variables jet.VarMap

	if vars == nil {
		variables = make(jet.VarMap)
	} else {
		variables = vars.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	tmpl, err := r.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, variables, td); err != nil {
		return err
	}

	return nil
}
