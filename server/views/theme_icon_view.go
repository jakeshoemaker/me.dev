package views

import (
    "html/template"
    "net/http"
)

type ThemeView struct {
    templ *template.Template
}

func NewThemeView(templ *template.Template) *ThemeView {
    return &ThemeView{templ: templ}
}

func (t *ThemeView) UpdateThemeIcon(w http.ResponseWriter, isDarkMode bool) bool {
    if isDarkMode {
        if err := t.templ.ExecuteTemplate(w, "theme_icon_light", nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return false
        }
        return true
    }

    if err := t.templ.ExecuteTemplate(w, "theme_icon_dark", nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return false
    }
    return true
}
