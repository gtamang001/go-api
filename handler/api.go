package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-go/model"
)

const (
	uatAPIURL  = "https://jsonplaceholder.typicode.com/users"
	prodAPIURL = "https://jsonplaceholder.typicode.com/posts"
)

// fetchData fetches and converts UAT or PROD API data to simplified APIResponse
func fetchData(env string) ([]model.APIResponse, error) {
	var url string

	switch env {
	case "uat":
		url = uatAPIURL
	case "prod":
		url = prodAPIURL
	default:
		return nil, fmt.Errorf("invalid environment: %s", env)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if env == "uat" {
		// Decode into []User model
		var users []model.User
		err = json.NewDecoder(resp.Body).Decode(&users)
		if err != nil {
			return nil, err
		}

		var data []model.APIResponse
		for _, user := range users {
			data = append(data, model.APIResponse{
				ID:    user.ID,
				Name:  user.Name,
				Value: user.Email,
			})
		}
		return data, nil
	}

	// Decode into []Post model
	var posts []model.Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	var data []model.APIResponse
	for _, post := range posts {
		data = append(data, model.APIResponse{
			ID:    post.ID,
			Name:  post.Title,
			Value: post.Body,
		})
	}
	return data, nil
}

// APIHandler handles requests and returns JSON for table display
func APIHandler(w http.ResponseWriter, r *http.Request) {
	env := r.URL.Query().Get("env")

	data, err := fetchData(env)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
