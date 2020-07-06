package todo

import (
	"net/http"
	"net/http/httptest"
    "strings"
    "testing"

    "../handler"
)

func TestGetSamples(t *testing.T) {
    testServer := setupServer()

    req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/samples", nil)
    if err != nil {
        t.Fatal(err)
    }

    rec := httptest.NewRecorder()
    testServer.ServeHTTP(rec, req)

    got := strings.TrimSpace(rec.Body.String())

    want := `[{"id":1,"title":"Practice GoLang","note":"Calculator","due_date":"2000-01-01T00:00:00Z"},{"id":2,"title":"Do homework","note":"ToDoApp","due_date":"2000-01-01T00:00:00Z"},{"id":3,"title":"Start Final Project","note":"Finalize Requirements","due_date":"2000-01-01T00:00:00Z"}]`
	
    if got != want {
        t.Fatalf("Want: %v, Got: %v", want, got)
    }
}

func setupServer() *http.ServeMux {
    return handler.SetUpRouting()
}