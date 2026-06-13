package tests

import (
	"bytes"
	"encoding/json"
	"go-back/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *http.Server

// func setUpServer() {

// 	server = &http.Server{
// 		Addr:    ":8080",
// 		Handler: router.SetUpRouter(),
// 	}

// 	err := server.ListenAndServe()

// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestPostAPI(t *testing.T) {

	server := httptest.NewServer(router.SetUpRouter())
	defer server.Close()

	t.Run("POST API", func(t *testing.T) {
		post := map[string]string{
			"title": "1 post",
			"desc":  "post 1",
		}

		postStr, err := json.Marshal(post)

		if err != nil {
			t.Fatalf("Failed to marshal post: %v", err)
		}

		res, err := http.Post(server.URL+"/posts", "application/json", bytes.NewBuffer(postStr))

		if err != nil {
			t.Fatalf("Api call Failed : %v", err)
		}

		if res.StatusCode != 201 {
			t.Fatalf("Failed to Create Post")
		}
	})

	t.Run("GET API", func(t *testing.T) {
		res, err := http.Get(server.URL + "/posts")

		if err != nil {
			t.Fatalf("Failed Fetch posts API : %v", err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("Failed to Get Posts")
		}
	})
}

// func TestPostGet(t *testing.T) {

// }

// func TestPostDelete()

// func TestPostUpdate()
