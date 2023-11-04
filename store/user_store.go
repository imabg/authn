package store

import (
	"github.com/imabg/authn/types"
	"github.com/jmoiron/sqlx"
)

type UserStoreInterface interface {
	Create(user *types.User) (string, error)
}

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
	query := `INSERT INTO users (id,email, country_code, source_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := u.db.QueryRow(query, data.ID, data.Email, data.CountryCode, data.SourceID).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
