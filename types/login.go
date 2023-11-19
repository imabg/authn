package types

import "time"

type Login struct {
	ID          int    `db:"id" json:"id"`
	Ip          string `db:"ip" json:"ip"`
	Platform    string `db:"platform" json:"platform"`
	UserAgent   string `db:"user_agent" json:"user_agent"`
	AccessToken string `db:"access_token" json:"access_token"`
	IsActive    bool   `db:"is_active" json:"is_active"`
	UserId      string `db:"user_id" json:"user_id"`
	LogoutAt    string `db:"logout_at" json:"logout_at"`
	CreatedAt   string `db:"created_at"`
}

type LoginViaEmailDTO struct {
	Contract string `json:"contract" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordLessLoginDTO struct {
	Contract string `json:"contract" binding:"required"`
}

type Platform int

const (
	WebPlatform Platform = iota
	MobilePlatform
	RestClientPlatform
)

type LoginLog struct {
	Ip          string   `json:"ip"`
	Platform    Platform `json:"platform"`
	UserAgent   string   `json:"user_agent"`
	AccessToken string   `json:"access_token"`
	UserId      string   `json:"user_id"`
	IsActive    bool     `json:"is_active"`
}

type LoginResponse struct {
	UserId      string    `json:"user_id"`
	AccessToken string    `json:"access_token"`
	SourceId    string    `json:"source_id"`
	UserAgent   string    `json:"user_agent"`
	IP          string    `json:"ip"`
	ExpireAt    time.Time `json:"expire_at"`
	CreatedAt   string    `json:"created_at"`
}
