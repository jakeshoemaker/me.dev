package views 

import (
    "embed"
    "html/template"
)

var (
    //go:embed "templates/*"
    templates embed.FS
)

func GenerateTemplates() (*template.Template, error) {
    return template.ParseFS(templates, "templates/*.html")
}
