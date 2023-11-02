package utils

import (
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func GenerateDisplayID() (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		return "", err
	}
	return id, nil
}
