package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dustin-ward/minecraft-time-logging/parser"
	"github.com/dustin-ward/minecraft-time-logging/routes"
	"github.com/dustin-ward/minecraft-time-logging/util"

	"github.com/gorilla/mux"
)

func main() {
	// Parse each log file in /logs/
	dir, _ := os.Getwd()
	err := filepath.Walk(dir+"/logs/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".log" {
			parser.Parse(path)
		}
		return nil
	})
	util.ErrorCheck(err)

	// Setup HTTP router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomePage)
	r.HandleFunc("/user/{username}", routes.UserPage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	fmt.Println("Page running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
