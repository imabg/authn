package types

import "encoding/json"

type Config struct {
	ID                 string          `db:"id" json:"id"`
	PasswordConfig     json.RawMessage `db:"password_config" json:"password_config"`
	MaxConcurrentLogin int             `db:"max_concurrent_login" json:"max_concurrent_login"`
	IsActive           bool            `db:"is_active" json:"-"`
	SourceId           string          `db:"source_id" json:"source_id"`
	CreatedAt          string          `db:"created_at" json:"created_at"`
}

type PasswordConfig struct {
	Length      int
	SmallCase   bool
	UpperCase   bool
	SpecialCase bool
	Numeric     bool
}
