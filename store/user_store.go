package store

import (
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"github.com/jmoiron/sqlx"
)

type UserStoreInterface interface {
	CreateViaEmail(user *types.User, password string) (string, error)
	CreateViaPhone(user *types.User) (string, error)
}

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) CreateViaEmail(data *types.User, password string) (string, error) {
	var id string
	tx, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	err = tx.QueryRow(`INSERT INTO users (id, email, source_id) VALUES ($1, $2, $3) RETURNING id`, utils.GenerateUUID(), data.Email, data.SourceID).Scan(&id)
	if err != nil {
		err = tx.Rollback()
		return "", err
	}
	_, err = tx.Query(`INSERT INTO credentials (id, password, user_id) VALUES ($1, $2, $3)`, utils.GenerateUUID(), password, id)
	if err != nil {
		err = tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserStore) CreateViaPhone(data *types.User) (string, error) {
	var id string
	query := `INSERT INTO users (id, phone, country_code, source_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := u.db.QueryRow(query, utils.GenerateUUID(), data.Phone, data.CountryCode, data.SourceID).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
