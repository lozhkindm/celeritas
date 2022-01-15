package render

import (
	"github.com/CloudyKit/jet/v6"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./testdata/views"),
	jet.InDevelopmentMode(),
)

var render = Render{
	Renderer: "",
	RootPath: "./testdata",
	JetViews: views,
}

var r *http.Request
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) {
	var err error

	if r, err = http.NewRequest("GET", "", nil); err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
}
