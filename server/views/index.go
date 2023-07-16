package views

import (
    "html/template"
    "net/http"

    // some internal model?
)

type IndexView struct {
    templ *template.Template

}

type Site struct {
    DarkMode bool
}

func NewIndexView(templ *template.Template) *IndexView {
    return &IndexView{templ: templ}
}

func (t *IndexView) Index(writer http.ResponseWriter) {
    site := Site {
        DarkMode: true,
    }
    if err := t.templ.ExecuteTemplate(writer, "index", site); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}
