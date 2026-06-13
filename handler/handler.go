package handler

import (
	"encoding/json"
	types "go-back/types"
	store "go-back/vars"
	"net/http"
	"strconv"
	"time"
)

var counter int16 = 0

func HandlePosts(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var data types.CreatePost

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		post := types.Post{
			Id:        counter,
			Title:     data.Title,
			Desc:      data.Desc,
			Timestamp: time.Now(),
		}

		store.Posts = append(store.Posts, post)

		counter++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)

		return
	}

	if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Posts)

		return

	}

	if r.Method == http.MethodPut {

		var data types.CreatePost

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, "Invalid Data to Update", http.StatusBadRequest)
		}

		idParam := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}

		for idx := 0; idx < len(store.Posts); idx++ {

			if store.Posts[idx].Id == int16(id) {
				store.Posts[idx].Title = data.Title
				store.Posts[idx].Desc = data.Desc
				store.Posts[idx].Timestamp = time.Now()

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(store.Posts[idx])
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

		for i, post := range store.Posts {
			if post.Id == int16(id) {
				store.Posts = append(store.Posts[:i], store.Posts[i+1:]...)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Post deleted successfully"))
				return
			}
		}

		http.Error(w, "Post not found", http.StatusNotFound)
	}

}
