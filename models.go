package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

type User struct {
	ID          uuid.UUID         `json:"id"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdateAt    time.Time         `json:"updated_at"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	NickName    sql.NullString    `json:"nick_name"`
	NumberPhone string            `json:"number_phone"`
	DayOfBirth  sql.NullTime      `json:"day_of_birth"`
	Address     sql.NullString    `json:"address"`
	Role        database.UserRole `json:"role"`
	APIKey      string            `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
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
		APIKey:      dbUser.ApiKey,
	}
}
