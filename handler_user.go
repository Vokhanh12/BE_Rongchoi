package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName   string            `json:"first_name"`
		LastName    string            `json:"last_name"`
		Email       string            `json:"email"`
		NickName    sql.NullString    `json:"nick_name"`
		NumberPhone string            `json:"number_phone"`
		DayOfBirth  sql.NullTime      `json:"day_of_birth"`
		Address     sql.NullString    `json:"address"`
		Role        database.UserRole `json:"role"`
	}

	result, err := json.Marshal(r.Body)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Marshal JSON: %v", err))
		return
	}

	reader := bytes.NewReader(result)

	decoder := json.NewDecoder(reader)

	params := parameters{}
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:         uuid.New(),
		CreateAt:   time.Now().UTC(),
		UpdateAt:   time.Now().UTC(),
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Email:      params.Email,
		NickName:   params.NickName,
		DayOfBirth: params.DayOfBirth,
		Address:    params.Address,
		Role:       params.Role,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v %s", err, string(result)))
		return
	}

	respondWithJSON(w, 200, user)
}
