package authentication

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider struct {
	id             uuid.UUID
	userId         uuid.UUID
	provider       string
	providerUserId string
	createdAt      time.Time
}

func (authProvider AuthProvider) Create(userId uuid.UUID, provider string, providerUserId string) (*AuthProvider, error) {
	newProvider := &AuthProvider{
		userId:         userId,
		provider:       provider,
		providerUserId: providerUserId,
	}
	return newProvider, nil
}
