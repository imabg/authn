package types

type Source struct {
	ID                 string `db:"id" json:"id"`
	Name               string `db:"name" json:"name"`
	Description        string `db:"description" json:"description"`
	DisableNewUser     bool   `db:"disable_new_user" json:"disable_new_user"`
	OptForPasswordless bool   `db:"opt_for_passwordless" json:"opt_for_passwordless"`
	IsActive           bool   `db:"is_active" json:"-"`
	CreatedAt          string `db:"created_at" json:"created_at"`
}

type SourceRTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GenerateSourceKeysRO struct {
	Key string `json:"key"`
}
