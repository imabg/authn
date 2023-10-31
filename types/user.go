package types

import "github.com/google/uuid"

type User struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	DisplayID       string    `json:"display_id"`
	Phone           int       `json:"phone"`
	IsEmailVerified bool      `json:"is_email_verified"`
	IsPhoneVerified bool      `json:"is_phone_verified"`
	CountryCode     string    `json:"country_code"`
	IsBlacklisted   bool      `json:"is_blacklisted"`
	SourceID        int64     `json:"source_id"`
	CreatedAt       string    `json:"created_at"`
}

type UserDTO struct {
	Email       string `json:"email" binding:"required, email"`
	Phone       int    `json:"phone" binding:"max=10"`
	CountryCode string `json:"country_code" binding:"iso3166_1_alpha2"`
}
