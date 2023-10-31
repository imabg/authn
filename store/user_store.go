package store

import (
	"github.com/imabg/authn/types"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) Create(data *types.User) (string, error) {
	var id string
	query := `INSERT INTO users (email, phone, country_code) VALUES ($1, $2, $3) RETURNING display_id`
	err := u.db.QueryRow(query, data.Email, data.Phone, data.CountryCode).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
