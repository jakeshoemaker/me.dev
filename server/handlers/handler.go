package handlers

import (
    "embed"
    "fmt"
    "io/fs"
    "net/http"

    //"github.com/google/uuid"
    "github.com/gorilla/mux"
    "github.com/jakeshoemaker/me.dev/server/views"
)

var (
    //go:embed static
    static embed.FS
)

type Handler struct {
    http.Handler
    //view *views.ModelView[models.View]
    index *views.IndexView
}

func CreateHandler(view *views.IndexView) (*Handler, error) {
    router := mux.NewRouter()
    handler := &Handler {
        Handler: router,
        index: view,
    }

    static_handler, err := createStaticHandler()
    if err != nil {
        return nil, fmt.Errorf("couldnt create static handler: %w", err)
    }

    router.HandleFunc("/", handler.get_index).Methods(http.MethodGet)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_handler))

    return handler, nil
}

func (t *Handler) get_index(writer http.ResponseWriter, _ *http.Request) {
    t.index.Index(writer)
}

func createStaticHandler() (http.Handler, error) {
    l, err := fs.Sub(static, "static")
    if err != nil {
        return nil, err
    }
    return http.FileServer(http.FS(l)), nil
}
