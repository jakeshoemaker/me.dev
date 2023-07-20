package controllers

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	//"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jakeshoemaker/me.dev/server/components"
  "github.com/jakeshoemaker/me.dev/server/helpers"
)

var (
    //go:embed static
    static embed.FS
)


type Controller struct {
    Router http.Handler
    root *components.Component
    darkMode *bool
    oppositeTheme *string
    curr_route *string
}


func CreateController(root *components.Component) (*Controller, error) {
    router := mux.NewRouter()
    controller := &Controller {
        Router: router,
        root: root,
        darkMode: helpers.Of(true),
        oppositeTheme: helpers.Of("light"),
        curr_route: helpers.Of("/"),
    }

    static_handler, err := create_static_handler()
    if err != nil {
        return nil, fmt.Errorf("couldnt create static handler: %w", err)
    }

    router.HandleFunc("/", controller.get_index).Methods(http.MethodGet)
    router.HandleFunc("/theme/icon/light", controller.put_light_theme_icon).Methods(http.MethodGet)
    router.HandleFunc("/theme/icon/dark", controller.put_dark_theme_icon).Methods(http.MethodGet)
	  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_handler))

    return controller, nil
}

func (t *Controller) get_index(writer http.ResponseWriter, _ *http.Request) {
    data := helpers.State {
        DarkMode: *t.darkMode,
        OppositeTheme: *t.oppositeTheme,
    }
    
    if err := t.root.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func (t *Controller) put_dark_theme_icon(writer http.ResponseWriter, _ *http.Request) {
    *t.darkMode = true
    *t.oppositeTheme = "light"
    data := helpers.State {
        DarkMode: true,
        OppositeTheme: "light",
    }
    if err := t.root.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }

}

func (t *Controller) put_light_theme_icon(writer http.ResponseWriter, _ *http.Request) {
    *t.darkMode = false
    *t.oppositeTheme = "dark"
    data := helpers.State {
        DarkMode: false,
        OppositeTheme: "dark",
    }
    if err := t.root.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func create_static_handler() (http.Handler, error) {
    l, err := fs.Sub(static, "static")
    if err != nil {
        return nil, err
    }
    return http.FileServer(http.FS(l)), nil
}
