package types

type Config struct {
	ID                     string `db:"id" json:"id"`
	PasswordLength         int    `db:"password_length" json:"password_length"`
	PasswordLowerAllowed   bool   `db:"password_lower_allowed" json:"password_lower_allowed"`
	PasswordUpperAllowed   bool   `db:"password_upper_allowed" json:"password_upper_allowed"`
	PasswordNumericAllowed bool   `db:"password_numeric_allowed" json:"password_numeric_allowed"`
	PasswordSpecialAllowed bool   `db:"password_special_allowed" json:"password_special_allowed"`
	IsSystemGenerated      bool   `db:"is_system_generated" json:"is_system_generated"`
	IsActive               bool   `db:"is_active" json:"-"`
	SourceId               string `db:"source_id" json:"source_id"`
	CreatedAt              string `db:"created_at" json:"created_at"`
}
