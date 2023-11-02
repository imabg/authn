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
	query := `INSERT INTO users (id,email, phone, country_code, source_id, display_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING display_id`
	err := u.db.QueryRow(query, data.ID, data.Email, data.Phone, data.CountryCode, data.SourceID, data.DisplayID).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
