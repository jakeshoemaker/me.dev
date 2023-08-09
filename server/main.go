package main

import (
    "os"
    "log"
    "net/http"
    "fmt"
    "context"
    "database/sql"

    "github.com/jakeshoemaker/me.dev/server/components"
    "github.com/jakeshoemaker/me.dev/server/controllers"
    "github.com/spf13/viper"
    _ "github.com/libsql/libsql-client-go/libsql"
)

const addr = ":8080"

func main() { 
    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil { 
        log.Fatalf("error while loading config file %s", err)
    }

    turso_auth := viper.GetString("TURSO_AUTH_TOKEN")
    log.Printf("turso auth token: %s", turso_auth)
    url := "libsql://me-dot-dev-db-jakeshoemaker.turso.io?authToken=" + turso_auth
    log.Printf("url: %s", url)
    db, err := sql.Open("libsql", url)
    if err != nil {
        log.Fatalf("failed to open db %s %s", url, err)  
    }
    rows ,err := db.Query("SELECT * FROM jobs;")
    if err != nil {
        log.Fatal(err)
    }
    
    for rows.Next() {
        var job struct {
            id int
            title string
            company string
            tenure string
            skill_id int
        }
        if err := rows.Scan(&job.id, &job.title, &job.company, &job.tenure, &job.skill_id); err != nil {
            log.Fatal(err)
        }
        fmt.Println(job)
    }

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


  func query(ctx context.Context, db *sql.DB, stmt string, args ...any) *sql.Rows {
      res, err := db.QueryContext(ctx, stmt, args...)
      if err != nil {
          fmt.Fprintf(os.Stderr, "failed to execute query %s: %s", stmt, err)
          os.Exit(1)
      }
      return res
  }
