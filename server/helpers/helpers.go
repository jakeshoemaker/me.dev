package helpers

func Of[Value any](v Value) *Value {
    return &v
}

type State struct {
    Theme string
    ViewInFocus string
}
