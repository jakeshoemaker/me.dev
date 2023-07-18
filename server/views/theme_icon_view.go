package views

import (
	"html/template"
	"net/http"

    "github.com/jakeshoemaker/me.dev/server/helpers"
)

type ThemeView struct {
    templ *template.Template
}

func NewThemeView(templ *template.Template) *ThemeView {
    return &ThemeView{templ: templ}
}

func (t *ThemeView) UpdateThemeIcon(w http.ResponseWriter, isDarkMode bool, data helpers.SiteData) {
    if isDarkMode {
        if err := t.templ.ExecuteTemplate(w, "theme_icon_light", data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else {
        if err := t.templ.ExecuteTemplate(w, "theme_icon_dark", data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}
