package utils

import (
	encoreUUID "encore.dev/types/uuid"
	googleUUID "github.com/google/uuid"
)

func GenerateUUID() (encoreUUID.UUID, error) {
	uuidValueGoogle, err := googleUUID.NewV7()
	if err != nil {
		return encoreUUID.Nil, err
	}
	// Convert to encore.dev/types/uuid.UUID

	uuidValueEncore, err := encoreUUID.FromString(uuidValueGoogle.String())
	if err != nil {
		return encoreUUID.Nil, err
	}

	return uuidValueEncore, nil
}
