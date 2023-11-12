package types

type Login struct {
	ID            int    `db:"id" json:"id"`
	Ip            string `db:"ip" json:"ip"`
	Platform      string `db:"platform" json:"platform"`
	UserAgent     string `db:"user_agent" json:"user_agent"`
	AccessToken   string `db:"access_token" json:"access_token"`
	IsActive      bool   `db:"is_active" json:"is_active"`
	IsBlacklisted bool   `db:"is_blacklisted" json:"is_blacklisted"`
	UserId        string `db:"user_id" json:"user_id"`
	LogoutAt      string `db:"logout_at" json:"logout_at"`
	CreatedAt     string `db:"created_at"`
}

type LoginViaEmailDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginViaPhoneDTO struct {
	Phone       string `json:"phone" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
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
