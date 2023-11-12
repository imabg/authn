package store

import (
	"errors"
	"github.com/imabg/authn/types"
	"github.com/jmoiron/sqlx"
)

type ILoginStore interface {
	CheckAccess(value string) (*types.User, error)
	Log(login *types.LoginLog) error
}

type LoginStore struct {
	db *sqlx.DB
}

func NewLoginStore(db *sqlx.DB) *LoginStore {
	return &LoginStore{db: db}
}

func (l *LoginStore) CheckAccess(value string) (*types.User, error) {
	user := types.User{}
	source := types.Source{}
	//TODO: check maximum session that are allowed against a source

	//check user
	err := l.db.Get(&user, `SELECT * FROM "users" WHERE email=$1`, value)
	if err != nil {
		return nil, err
	}
	//check source
	err = l.db.Get(&source, `SELECT * FROM sources WHERE id=$1`, user.SourceID)
	if !source.IsActive {
		return nil, errors.New("operation not allowed on this source")
	}
	if err != nil {
		return nil, err
	}
	if user.IsBlacklisted {
		return nil, errors.New("user is blacklisted")
	}
	return &user, err
}

func (l *LoginStore) Log(data *types.LoginLog) error {
	// mark is-active False to the earlier log record
	_, err := l.db.Query(`UPDATE "logins" SET is_active=false WHERE user_id=$1 AND is_active=true`, data.UserId)
	_, err = l.db.Query(`INSERT INTO "logins" (ip, platform, user_agent, access_token, user_id, is_active) VALUES ($1, $2, $3,$4, $5, $6)`, data.Ip, data.Platform, data.UserAgent, data.AccessToken, data.UserId, data.IsActive)
	if err != nil {
		return err
	}
	return nil
}
