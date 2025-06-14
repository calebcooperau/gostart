package delete

import "github.com/google/uuid"

type UserDeleteCommand struct {
	Id uuid.UUID
}
