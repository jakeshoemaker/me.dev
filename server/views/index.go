package views

import (
    "html/template"
    "net/http"

    // some internal model?
)

type IndexView struct {
    templ *template.Template
}

func NewIndexView(templ *template.Template) *IndexView {
    return &IndexView{templ: templ}
}

func (t *IndexView) Index(writer http.ResponseWriter) {
    if err := t.templ.ExecuteTemplate(writer, "index", nil); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}
