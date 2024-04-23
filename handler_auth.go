package main

import (
	"net/http"
	"time"

	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

// handleCreateUser use for [V1Router]
func (apiCfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request, user database.User) {
	userJSON := databaseLoginRespToLoginResp(user)
	respondWithJSON(w, 200, userJSON)
}

type LoginResponse struct {
	User          User              `json:"user"`
	APIKey        APIKeyResponse    `json:"api_key"`
	RefreshAPIKey RefAPIKeyResponse `json:"refresh_api_key"`
}

func databaseLoginRespToLoginResp(dbUser database.User) LoginResponse {
	return LoginResponse{
		User: User{
			ID:          dbUser.ID,
			CreatedAt:   dbUser.CreateAt,
			UpdateAt:    dbUser.UpdateAt,
			FirstName:   dbUser.FirstName,
			LastName:    dbUser.LastName,
			Email:       dbUser.Email,
			NickName:    dbUser.NickName,
			NumberPhone: dbUser.NumberPhone,
			DayOfBirth:  dbUser.DayOfBirth,
			Address:     dbUser.Address,
			Role:        dbUser.Role,
		},
		APIKey: APIKeyResponse{
			APIKey: dbUser.ApiKey,
			Iat:    dbUser.ApiIat,
			Exp:    dbUser.ApiExp,
		},
		RefreshAPIKey: RefAPIKeyResponse{
			RefAPIKey: dbUser.RefreshApiKey,
			Iat:       dbUser.RefApiIat,
			Exp:       dbUser.RefApiExp,
		},
	}
}

type APIKeyResponse struct {
	APIKey string    `json:"key"`
	Iat    time.Time `json:"iat"`
	Exp    time.Time `json:"exp"`
}

type RefAPIKeyResponse struct {
	RefAPIKey string    `json:"key"`
	Iat       time.Time `json:"iat"`
	Exp       time.Time `json:"exp"`
}
