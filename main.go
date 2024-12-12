package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static/*
var staticFiles embed.FS

//go:embed templates/*
var templateFiles embed.FS

var DisplayApiKey = ""
var TrainStationId = ""
var TagLine = ""

type config struct {
	DisplayApiKey  string
	TrainStationId string
	TagLine        string
}

func main() {
	staticFs, err := fs.Sub(fs.FS(staticFiles), "static")
	if err != nil {
		log.Fatalf("could not load static files: %v", err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFs))))

	if DisplayApiKey == "" {
		DisplayApiKey = os.Getenv("DISPLAY_API_KEY")
	}

	if TrainStationId == "" {
		TrainStationId = os.Getenv("TRAIN_STATION_ID")
	}

	if TagLine == "" {
		TagLine = os.Getenv("TAG_LINE")
	}

	tmpl, err := template.ParseFS(templateFiles, "templates/*.html")
	if err != nil {
		log.Fatalf("could not load templates: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		err = tmpl.ExecuteTemplate(w, "index.html", config{
			DisplayApiKey:  DisplayApiKey,
			TrainStationId: TrainStationId,
			TagLine:        TagLine,
		})
		if err != nil {
			log.Fatalf("could not parse template: %v", err)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("listening on http://localhost:%s \n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
