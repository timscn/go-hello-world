package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h2Hello from Web :D")
// }

var tpl = template.Must(template.ParseFiles("static/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func mainWeb() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server starting...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}
