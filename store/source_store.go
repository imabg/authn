package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/imabg/authn/types"
	"github.com/jmoiron/sqlx"
)

type SourceStoreInterface interface {
	Create(*types.Source) (int, error)
	GetByID(id int) (*types.Source, error)
}

type SourceStore struct {
	db *sqlx.DB
}

func NewSourceStore(db *sqlx.DB) *SourceStore {
	return &SourceStore{
		db: db,
	}
}

func (s *SourceStore) Create(data *types.Source) (int, error) {
	var id int
	query := `INSERT INTO sources (name, description) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(query, data.Name, data.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SourceStore) GetByID(id int) (*types.Source, error) {
	var source types.Source
	query := `SELECT * FROM sources WHERE id=$1 AND is_active=true`
	fmt.Print(id)
	err := s.db.Get(&source, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("source not found")
	} else if err != nil {
		return nil, err
	} else {
		return &source, nil
	}
}
