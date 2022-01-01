package simplefetch_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	fetch "github.com/11me/simplefetch"
)

func TestFetch(t *testing.T) {

	// create a local listening server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Test data")
	}))

	t.Run("Get method", func(t *testing.T) {
		res, err := fetch.Get(fetch.Options{
			URL: server.URL,
		})
		if err != nil {
			t.Error(err.Error())
		}

		if res.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %v", res.StatusCode)
		}
	})

	defer server.Close()
}
