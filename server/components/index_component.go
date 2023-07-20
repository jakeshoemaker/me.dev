package components

import (
	"html/template"
)

type IndexComponent struct {
    Templ *template.Template
}

func NewIndexView(templ *template.Template) *IndexComponent {
    return &IndexComponent{Templ: templ}
}
