package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

// handleCreatePost use for [V1Router]
func (apiCfg *apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title       string `json:"title"`
		Content     string `json:"content"`
		NumberPhone string `json:"number_phone"`
		Address     string `json:"address"`
		NickName    string `json:"nick_name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// check NumberPhone
	var pUserAddress string
	if user.Address.Valid {
		pUserAddress = user.Address.String
	} else {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create Post has address is null: %v", err))
		return
	}

	// check NickName
	var pUserNickName string
	if user.NickName.Valid {
		pUserNickName = user.NickName.String
	} else {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create Post has NickName is null: %v", err))
		return
	}

	// Create post used to [Main]
	post, err := apiCfg.DB.CreatePost(r.Context(), database.CreatePostParams{
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       params.Title,
		Content:     params.Content,
		NumberPhone: user.NumberPhone,
		Address:     pUserAddress,
		NickName:    pUserNickName,
		UserID:      user.ID,
	})

	// Check error for Create post
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	// Create user Success
	respondWithJSON(w, 201, databasePostToPost(post))
}

// handleCreateUser use for [V1Router]
func (apiCfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {

	// Get Post used to [Main]
	posts, err := apiCfg.DB.GetPosts(r.Context())

	// Check error for get posts
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get post: %v", err))
		return
	}

	// Create user Success
	respondWithJSON(w, 201, databasePostsToPosts(posts))
}
