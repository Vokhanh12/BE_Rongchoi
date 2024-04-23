package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

// handleCreateUser use for [V1Router]
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		NickName    string `json:"nick_name"`
		NumberPhone string `json:"number_phone"`
		DayOfBirth  string `json:"day_of_birth"`
		Address     string `json:"address"`
		Role        string `json:"role"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// check DayOfBirth
	var pDayOfBirth = sql.NullTime{}
	if params.DayOfBirth != "" {
		date, err := time.Parse("2006-01-02", params.DayOfBirth)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error parsing date: %v", err))
			return
		}
		pDayOfBirth = sql.NullTime{Time: date, Valid: true}
	} else {
		// declare Time is zero
		pDayOfBirth = sql.NullTime{Time: time.Time{}, Valid: false}
	}

	// check NickName
	var pNickName = sql.NullString{}
	if params.NickName != "" {
		pNickName = sql.NullString{String: params.NickName, Valid: true}
	} else {
		// declare String is Empty
		pNickName = sql.NullString{String: "", Valid: false}
	}

	// check Email
	var pAddress = sql.NullString{}
	if params.Address != "" {
		pAddress = sql.NullString{String: params.Address, Valid: true}
	} else {
		// declare String is Empty
		pAddress = sql.NullString{String: "", Valid: false}
	}

	// check Role [Role default is Employee]
	var pRole database.UserRole
	switch params.Role {
	case string(database.UserRoleBuyer):
		pRole = database.UserRoleBuyer
	case string(database.UserRoleEmployee):
		pRole = database.UserRoleEmployee
	default:
		respondWithError(w, 400, fmt.Sprintf("Can't not find role %s:", params.Role))
		return
	}

	// Create User used to [Main]
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		CreateAt:    time.Now().UTC(),
		UpdateAt:    time.Now().UTC(),
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		NumberPhone: params.NumberPhone,
		Email:       params.Email,
		NickName:    pNickName,   // able null
		DayOfBirth:  pDayOfBirth, // able null
		Address:     pAddress,    // able null
		Role:        pRole,
	})

	// Check error for Create User
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	// Create user Success
	respondWithJSON(w, 201, databaseUserToUser(user))
}

// handleCreateUser use for [V1Router]
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	userJSON := databaseUserToUser(user)
	respondWithJSON(w, 200, map[string]User{"user": userJSON})
}
