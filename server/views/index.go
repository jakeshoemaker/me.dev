package views

import (
	"html/template"
)

type IndexView struct {
    Templ *template.Template

}

func NewIndexView(templ *template.Template) *IndexView {
    return &IndexView{Templ: templ}
}
