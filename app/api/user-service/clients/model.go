package clients

import (
	"time"

	"github.com/google/uuid"
)

type ClientInfo struct {
	ID           uuid.UUID `db:"id" json:"id"`
	FirstName    string    `db:"first_name" json:"first-name"`
	LastName     string    `db:"last_name" json:"last-name"`
	PasswordHash []byte    `db:"password_hash" json:"-"`

	BirthDate time.Time `db:"birth_date" json:"birth-date"`
	Gender    string    `db:"gender" json:"gender"`
	Phone     string    `db:"phone" json:"phone"`
	Address   string    `db:"address" json:"address"`
	Email     string    `db:"email" json:"email"`
}

type ClientWithPassword struct {
	ClientInfo
	Password string `json:"password"`
}
