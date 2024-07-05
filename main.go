package main

import (
	"fmt"
	"net/http"
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

func main() {
	http.HandleFunc("/get", Get)
	http.ListenAndServe(":8080", nil)
}
