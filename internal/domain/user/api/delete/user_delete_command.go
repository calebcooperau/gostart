package delete

import "github.com/google/uuid"

type UserDeleteCommand struct {
	ID uuid.UUID
}
