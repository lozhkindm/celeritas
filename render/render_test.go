package render

import "testing"

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"page_go", "go", "home", false, "Error rendering tmpl using go renderer"},
	{"page_go_no_tmpl", "go", "no-tmpl", true, "No error parsing non-existent tmpl using go renderer"},
	{"page_jet", "jet", "home", false, "Error rendering tmpl using jet renderer"},
	{"page_jet_no_tmpl", "jet", "no-tmpl", true, "No error parsing non-existent tmpl using jet renderer"},
	{"invalid_renderer", "invalid", "home", true, "No error rendering tmpl using invalid renderer"},
}

func TestRender_Page(t *testing.T) {
	setup(t)
	for _, c := range pageData {
		render.Renderer = c.renderer
		err := render.Page(w, r, c.template, nil, nil)

		if c.errorExpected && err == nil {
			t.Errorf("%s: %s", c.name, c.errorMessage)
		} else if !c.errorExpected && err != nil {
			t.Errorf("%s: %s: %s", c.name, c.errorMessage, err.Error())
		}
	}
}
