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
    theme *string
    viewInFocus *string
}


func CreateController(root *components.Component) (*Controller, error) {
    router := mux.NewRouter()
    controller := &Controller {
        Router: router,
        root: root,
        theme: helpers.Of("dark"),
        viewInFocus: helpers.Of("main_greeting"),
    }

    static_handler, err := create_static_handler()
    if err != nil {
        return nil, fmt.Errorf("couldnt create static handler: %w", err)
    }

    router.HandleFunc("/", controller.get_index).Methods(http.MethodGet)
    router.HandleFunc("/themes/{theme}", controller.set_theme).Methods(http.MethodGet)
    router.HandleFunc("/resume", controller.resume).Methods(http.MethodGet)
	  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_handler))

    return controller, nil
}

func (t *Controller) projects(writer http.ResponseWriter, _ *http.Request) {
    *t.viewInFocus = "projects"
    data := helpers.State {
        Theme: *t.theme,
        ViewInFocus: *t.viewInFocus,    
    }
    
    if err := t.root.Templ.ExecuteTemplate(writer, "resume", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func (t *Controller) resume(writer http.ResponseWriter, _ *http.Request) {
    *t.viewInFocus = "resume"
    data := helpers.State {
        Theme: *t.theme,
        ViewInFocus: "resume",    
    }


    if err := t.root.Templ.ExecuteTemplate(writer, "resume", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func (t *Controller) get_index(writer http.ResponseWriter, _ *http.Request) {
    *t.viewInFocus = "main_greeting"
    data := helpers.State {
        Theme: *t.theme,
        ViewInFocus: *t.viewInFocus,
    }
    
    if err := t.root.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func (t *Controller) set_theme(writer http.ResponseWriter, req *http.Request) {
    vars :=mux.Vars(req)
    *t.theme = vars["themes"]
    data := helpers.State {
        Theme: *t.theme,
        ViewInFocus: *t.viewInFocus,
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
