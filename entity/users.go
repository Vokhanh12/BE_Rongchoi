package entity

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
)

type User struct {
	ID          uuid.UUID         "json:id"
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	NickName    sql.NullString    `json:"nick_name"`
	NumberPhone string            `json:"number_phone"`
	DayOfBirth  sql.NullTime      `json:"day_of_birth"`
	Address     sql.NullString    `json:"address"`
	Role        database.UserRole `json:"role"`
}
