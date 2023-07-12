package main

import (
	"log"
	"net/http"

	"github.com/jakeshoemaker/me.dev/server/handlers"
	"github.com/jakeshoemaker/me.dev/server/views"
)

const addr = ":8080"

func main() { 
    templates , err := views.GenerateTemplates()
    if err != nil {
        log.Fatal(err)
    }

    handler, err := handlers.CreateHandler(views.NewIndexView(templates))
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
