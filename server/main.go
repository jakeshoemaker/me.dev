package main

import (
	"log"
	"net/http"

	"github.com/jakeshoemaker/me.dev/server/controllers"
	"github.com/jakeshoemaker/me.dev/server/components"
)

const addr = ":8080"

func main() { 
    templates , err := components.GenerateTemplates()
    if err != nil {
        log.Fatal(err)
    }

    controller, err := controllers.CreateController(components.NewComponent(templates))
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("listening on %s", addr)
	  
  if err := http.ListenAndServe(addr, controller.Router); err != nil {
	  	log.Fatal(err)
	}
}
