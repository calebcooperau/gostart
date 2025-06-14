package dtos

import (
	"time"

	"github.com/google/uuid"
)

type UserDetailApiDto struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
