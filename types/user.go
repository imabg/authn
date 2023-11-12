package types

import "database/sql"

type User struct {
	ID              string         `db:"id" json:"id"`
	Email           string         `db:"email" json:"email"`
	Phone           sql.NullString `db:"phone" json:"phone"`
	IsEmailVerified bool           `db:"is_email_verified" json:"is_email_verified"`
	IsPhoneVerified bool           `db:"is_phone_verified" json:"is_phone_verified"`
	CountryCode     sql.NullString `db:"country_code" json:"country_code"`
	IsBlacklisted   bool           `db:"is_blacklisted" json:"is_blacklisted"`
	SourceID        string         `db:"source_id" json:"source_id"`
	CreatedAt       string         `db:"created_at" json:"created_at"`
}

type UserEmailDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	SourceID string `json:"source_id" binding:"required"`
}

type UserPhoneDTO struct {
	Phone       string `json:"phone" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
	SourceID    string `json:"source_id" binding:"required"`
}
