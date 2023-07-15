package views

import (
    "html/template"
    "net/http"

    // some internal model?
)

type IndexView struct {
    templ *template.Template
}

type Self struct {
    Name string
    Age int16
}

func NewIndexView(templ *template.Template) *IndexView {
    return &IndexView{templ: templ}
}

func (t *IndexView) Index(writer http.ResponseWriter) {
    me := Self {
        Name: "urmom",
        Age: 6,
    }
    if err := t.templ.ExecuteTemplate(writer, "index", me); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}
