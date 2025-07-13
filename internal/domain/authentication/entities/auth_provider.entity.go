package entities

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider struct {
	id             uuid.UUID
	userID         uuid.UUID
	provider       string
	providerUserID string
	createdAt      time.Time
}

func (authProvider AuthProvider) Create(userID uuid.UUID, provider string, providerUserID string) (*AuthProvider, error) {
	newProvider := &AuthProvider{
		userID:         userID,
		provider:       provider,
		providerUserID: providerUserID,
	}
	return newProvider, nil
}
