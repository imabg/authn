package store

import "github.com/jmoiron/sqlx"

type ICredentialStore interface {
	GetUserCred(userId string) (string, error)
}

type CredentialStore struct {
	db *sqlx.DB
}

func NewCredentialStore(db *sqlx.DB) *CredentialStore {
	return &CredentialStore{
		db: db,
	}
}

func (s *CredentialStore) GetUserCred(userId string) (string, error) {
	var hashPwd string
	err := s.db.Get(&hashPwd, `SELECT password FROM credentials WHERE user_id=$1`, userId)
	if err != nil {
		return "", err
	}
	return hashPwd, nil
}
