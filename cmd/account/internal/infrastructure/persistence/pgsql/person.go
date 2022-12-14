package pgsql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
)

// Person pgsql model.
type Person struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	AccountID   uuid.UUID      `json:"accountId" db:"account_id"`
	FirstName   string         `json:"firstName" db:"first_name"`
	LastName    string         `json:"lastName" db:"last_name"`
	Email       string         `json:"email" db:"email"`
	Phone       string         `json:"phone" db:"phone"`
	DateOfBirth time.Time      `json:"dateOfBirth" db:"date_of_birth"`
	Avatar      sql.NullString `json:"avatar" db:"avatar"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time      `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the person model to account entity.
func (p Person) ToEntity() person.Entity {
	return person.Entity{
		ID:        p.ID,
		AccountID: p.AccountID,
		Details: person.Details{
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			Email:       p.Email,
			Phone:       p.Phone,
			DateOfBirth: p.DateOfBirth,
		},
		Avatar: person.Avatar(p.Avatar.String),
	}
}
