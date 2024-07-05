package main

import (
	"bytes"
	"encoding/json"
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

func TestPost(t *testing.T) {
	t.Run("error in decoding the body", func(t *testing.T) {

		// invalid json
		req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		Post(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		expected := "invalid request\n"

		body := new(bytes.Buffer)
		body.ReadFrom(res.Body)
		if body.String() != expected {
			t.Errorf("expected body %q, got %q", expected, body.String())
		}

	})

	t.Run("name is empty", func(t *testing.T) {

		// valid json

		body := RequestBody{
			Name: "",
		}

		bodyBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "applicattion/json")
		w := httptest.NewRecorder()

		Post(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		expected := "name is empty\n"
		bodyBuf := new(bytes.Buffer)

		bodyBuf.ReadFrom(res.Body)

		if bodyBuf.String() != expected {
			t.Errorf("expected body %q, got %q", expected, bodyBuf.String())
		}
	})

	t.Run("valid name", func(t *testing.T) {

		// valid json with name

		body := RequestBody{
			Name: "Navneet",
		}

		bodyBytes, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		Post(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, res.StatusCode)
		}

		expected := "Total vowel is 3"

		bodyBuf := new(bytes.Buffer)
		bodyBuf.ReadFrom(res.Body)

		if strings.TrimSpace(bodyBuf.String()) != expected {
			t.Errorf("expected body %q, got %q", expected, bodyBuf.String())

		}
	})

}
