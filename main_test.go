package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	t.Run("Query is empty", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/get?name=", nil)
		w := httptest.NewRecorder()

		Get(w, r)

		res := w.Result()

		defer res.Body.Close()

		expected := "query is empty"

		data, err := io.ReadAll(res.Body)
		ans := string(data)
		if err != nil {
			t.Errorf("Expected is %s and result is %s ", expected, ans)
		}
	})

	t.Run("query is not empty and answer is correct", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/get?name=navneet", nil)
		w := httptest.NewRecorder()

		Get(w, r)

		res := w.Result()

		defer res.Body.Close()

		expected := "Total vowel is 3"

		data, err := io.ReadAll(res.Body)
		ans := strings.TrimSpace(string(data))
		if err != nil {
			t.Errorf("Expected is %s and result is %s ", expected, ans)
		}
		if ans != expected {
			t.Errorf("Expected is %s and result is %s ", expected, ans)

		}

	})

	t.Run("query is not empty and answer is incorrect", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/get?name=navneetShukla", nil)
		w := httptest.NewRecorder()

		Get(w, r)

		res := w.Result()

		defer res.Body.Close()

		expected := "Total vowel is 3"

		data, err := io.ReadAll(res.Body)
		ans := strings.TrimSpace(string(data))
		if err != nil {
			t.Errorf("Expected is %s and result is %s ", expected, ans)
		}
		if ans == expected {
			t.Errorf("Expected is %s and result is %s ", expected, ans)

		}

	})

}
