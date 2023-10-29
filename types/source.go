package types

type Source struct {
	ID                  int64  `db:"id" json:"id"`
	Name                string `db:"name" json:"name"`
	Description         string `db:"description" json:"description"`
	DisableUserCreation bool   `db:"disable_user_creation" json:"disable_user_creation"`
	IsActive            bool   `db:"is_active" json:"-"`
	CreatedAt           string `db:"created_at" json:"created_at"`
}

type SourceRTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
