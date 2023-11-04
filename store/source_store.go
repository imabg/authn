package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"github.com/jmoiron/sqlx"
)

type SourceStoreInterface interface {
	Create(*types.Source) (string, error)
	GetByID(id string) (*types.Source, error)
	GetConfig(id string) (*types.Config, error)
}

type SourceStore struct {
	db *sqlx.DB
}

func NewSourceStore(db *sqlx.DB) *SourceStore {
	return &SourceStore{
		db: db,
	}
}

func (s *SourceStore) Create(data *types.Source) (string, error) {
	var id string
	query := `INSERT INTO sources (id, name, description) VALUES ($1, $2, $3) RETURNING id`
	configQuery := `INSERT INTO "sourceConfigs" (id, source_id) VALUES ($1, $2)`
	err := s.db.QueryRow(query, data.ID, data.Name, data.Description).Scan(&id)
	err = s.db.QueryRow(configQuery, utils.GenerateUUID(), id).Err()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *SourceStore) GetByID(id string) (*types.Source, error) {
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

func (s *SourceStore) GetConfig(sourceId string) (*types.Config, error) {
	var config types.Config
	query := `SELECT * FROM "sourceConfigs" WHERE source_id=$1`
	err := s.db.Get(&config, query, sourceId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("source not found")
	} else if err != nil {
		return nil, err
	} else {
		return &config, nil
	}
}
