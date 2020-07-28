package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var layoutDir string = "views/layouts"

type View struct {
	Template *template.Template
	Layout   string
}

type ViewData struct {
	Data interface{}
}

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
