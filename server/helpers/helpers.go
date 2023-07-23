package helpers

func Of[Value any](v Value) *Value {
    return &v
}

type State struct {
    DarkMode bool
    OppositeTheme string
    ViewInFocus string
}
