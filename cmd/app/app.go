package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Get("/search", getSearch)

	log.Println("Starting up on http://localhost: 3333")

	http.ListenAndServe(":3333", r)
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("value")

	base := `{
        "search_type": "match",
        "query":
        {
            "term": "%s"
        },
        "from": 0,
        "max_results": 10,
        "_source": ["From", "To"]
    }`
	query := fmt.Sprintf(base, value)
	req, err := http.NewRequest("POST", "http://localhost:4080/api/mails/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
