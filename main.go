package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var counter int16 = 0

type Post struct {
	Id        int16     `json:id`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Timestamp time.Time `json:"timestamp"`
}

var posts = []Post{}

type CreatePost struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func handlePosts(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var data CreatePost

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		post := Post{
			Id:        counter,
			Title:     data.Title,
			Desc:      data.Desc,
			Timestamp: time.Now(),
		}

		posts = append(posts, post)

		counter++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)

		return
	}

	if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)

		return

	}

	if r.Method == http.MethodPut {

		var data CreatePost

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, "Invalid Data to Update", http.StatusBadRequest)
		}

		idParam := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}

		for idx := 0; idx < len(posts); idx++ {

			if posts[idx].Id == int16(id) {
				posts[idx].Title = data.Title
				posts[idx].Desc = data.Desc
				posts[idx].Timestamp = time.Now()

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(posts[idx])
				return
			}
		}

		http.Error(w, "Post Don't Exist To Update", http.StatusNotFound)
	}

	if r.Method == http.MethodDelete {

		idParam := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}

		for i, post := range posts {
			if post.Id == int16(id) {
				posts = append(posts[:i], posts[i+1:]...)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post deleted successfully"))
				return
			}
		}

		http.Error(w, "Post not found", http.StatusNotFound)
	}

}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/posts", handlePosts)

	router.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"status":"ok"}`))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Go Server Starting on 8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		log.Fatal(err)
	}
}
