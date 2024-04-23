// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleBuyer    UserRole = "buyer"
	UserRoleEmployee UserRole = "employee"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Post struct {
	ID          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Content     string
	NumberPhone string
	Address     string
	NickName    string
	UserID      uuid.UUID
}

type User struct {
	ID            uuid.UUID
	CreateAt      time.Time
	UpdateAt      time.Time
	FirstName     string
	LastName      string
	Email         string
	NickName      sql.NullString
	NumberPhone   string
	DayOfBirth    sql.NullTime
	Address       sql.NullString
	Role          UserRole
	ApiKey        string
	ApiIat        time.Time
	ApiExp        time.Time
	RefreshApiKey string
	RefApiIat     time.Time
	RefApiExp     time.Time
}
