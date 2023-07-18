package handlers

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	//"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jakeshoemaker/me.dev/server/views"
    "github.com/jakeshoemaker/me.dev/server/helpers"
)

var (
    //go:embed static
    static embed.FS
)


type Handler struct {
    http.Handler
    //view *views.ModelView[models.View]
    index *views.IndexView
    theme_view *views.ThemeView
    darkMode *bool
    oppositeTheme *string
}


func CreateHandler(view *views.IndexView, theme_view *views.ThemeView) (*Handler, error) {
    router := mux.NewRouter()
    handler := &Handler {
        Handler: router,
        index: view,
        theme_view: theme_view,
        darkMode: helpers.Of(true),
        oppositeTheme: helpers.Of("light"),
    }

    static_handler, err := createStaticHandler()
    if err != nil {
        return nil, fmt.Errorf("couldnt create static handler: %w", err)
    }

    router.HandleFunc("/", handler.get_index).Methods(http.MethodGet)
    router.HandleFunc("/theme/icon/light", handler.put_light_theme_icon).Methods(http.MethodGet)
    router.HandleFunc("/theme/icon/dark", handler.put_dark_theme_icon).Methods(http.MethodGet)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static_handler))

    return handler, nil
}

func (t *Handler) get_index(writer http.ResponseWriter, _ *http.Request) {
    data := helpers.SiteData {
        DarkMode: *t.darkMode,
        OppositeTheme: *t.oppositeTheme,
    }
    if err := t.index.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func (t *Handler) put_dark_theme_icon(writer http.ResponseWriter, _ *http.Request) {
    *t.darkMode = true
    *t.oppositeTheme = "light"
    data := helpers.SiteData {
        DarkMode: true,
        OppositeTheme: "light",
    }
    if err := t.index.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }

}

func (t *Handler) put_light_theme_icon(writer http.ResponseWriter, _ *http.Request) {
    *t.darkMode = false
    *t.oppositeTheme = "dark"
    data := helpers.SiteData {
        DarkMode: false,
        OppositeTheme: "dark",
    }
    if err := t.index.Templ.ExecuteTemplate(writer, "index", data); err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}

func createStaticHandler() (http.Handler, error) {
    l, err := fs.Sub(static, "static")
    if err != nil {
        return nil, err
    }
    return http.FileServer(http.FS(l)), nil
}
