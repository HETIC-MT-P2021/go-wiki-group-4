package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var layoutDir string = "views/layouts"

// View is the struct used to generate rendered web pages
type View struct {
	Template *template.Template
	Layout   string
}

// ViewData is an interface to contain Data
type ViewData struct {
	Data interface{}
}

// NewView generates a View using a layout and a file name
func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render insert Data to a web page
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{
		Data: data,
	}
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "/*.html")
	if err != nil {
		panic(err)
	}
	return files
}
