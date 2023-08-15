package components

import (
	"html/template"
)

type Component struct {
	Templ *template.Template
}

func NewComponent(templ *template.Template) *Component {
	return &Component{Templ: templ}
}
