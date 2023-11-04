package types

type Source struct {
	ID                  string `db:"id" json:"id"`
	Name                string `db:"name" json:"name"`
	Description         string `db:"description" json:"description"`
	DisableUserCreation bool   `db:"disable_user_creation" json:"disable_user_creation"`
	PasswordConfig      string `db:"password_config" json:"password_config"`
	IsActive            bool   `db:"is_active" json:"-"`
	CreatedAt           string `db:"created_at" json:"created_at"`
}

type SourceRTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
