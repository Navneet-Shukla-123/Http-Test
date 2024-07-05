package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")

	if len(query) == 0 {
		w.Write([]byte("Query is empty"))
		return
	}

	cnt := 0

	for _, val := range query {
		if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
			cnt++
		}
	}

	w.Write([]byte(fmt.Sprintf("Total vowel is %d ", cnt)))
}

type RequestBody struct {
	Name string `json:"name"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	name := reqBody.Name
	if len(name) == 0 {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	cnt := 0
	for _, val := range name {
		if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
			cnt++
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Total vowel is %d", cnt)
}

func main() {
	mux := chi.NewRouter()
	mux.Get("/get", Get)
	http.ListenAndServe(":8080", nil)
}
