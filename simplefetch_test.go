package simplefetch_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	fetch "github.com/11me/simplefetch"
)

func TestFetch(t *testing.T) {

	t.Run("GET method", func(t *testing.T) {
		// create a local listening server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprintf(w, "Test data")
		}))

		res, err := fetch.Get(fetch.Options{
			URL: server.URL,
		})
		if err != nil {
			t.Error(err.Error())
		}

		if res.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %v", res.StatusCode)
		}
		defer server.Close()
	})

	t.Run("POST method", func(t *testing.T) {
		dataMap := make(map[string]interface{})
		// create a local listening server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)

			data, _ := io.ReadAll(r.Body)
			json.Unmarshal(data, &dataMap)

			if dataMap["id"] != float64(1) {
				t.Errorf("Expected id 1, but got %d", dataMap["id"])
			}

			if dataMap["name"] != "lime" {
				t.Errorf("Expected name lime, but got %s", dataMap["name"])
			}

			fmt.Fprintf(w, "Test data")
		}))

		res, err := fetch.Post(fetch.Options{
			URL: server.URL,
			Data: fetch.Data{
				"id":   1,
				"name": "lime",
			},
		})
		if res.StatusCode != 200 {
			t.Error(err)
		}
		defer server.Close()
	})
}
