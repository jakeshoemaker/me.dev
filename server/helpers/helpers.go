package helpers

func Of[Value any](v Value) *Value {
    return &v
}


type SiteData struct {
    DarkMode bool
    OppositeTheme string
}
