package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "golang.org/x/crypto/bcrypt"
	"gostart/internal/domain/user/data"
	"gostart/internal/domain/user/entities"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	RegisterUser(ctx context.Context, cmp *entities.User) (uuid.UUID, error)
	UpdateUser(ctx context.Context, cmp *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, cmp *entities.User) error
}

type UserSqlRepository struct {
	queries *data.Queries
	db      *pgxpool.Pool
}

func NewUserSqlRepository(db *pgxpool.Pool) *UserSqlRepository {
	queries := data.New(db)
	return &UserSqlRepository{
		queries: queries,
		db:      db,
	}
}

func (repository *UserSqlRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	userData, err := repository.queries.FindUserByID(ctx, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:           userData.ID,
		Email:        userData.Email,
		FirstName:    userData.FirstName,
		LastName:     userData.LastName,
		MobileNumber: *userData.MobileNumber,
		CreatedAt:    userData.CreatedAt.Time,
		UpdatedAt:    userData.UpdatedAt.Time,
	}

	return user, nil
}

func (repository *UserSqlRepository) RegisterUser(ctx context.Context, user *entities.User) (uuid.UUID, error) {
	addUserParams := data.AddUserParams{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: &user.MobileNumber,
	}
	userResult, err := repository.queries.AddUser(ctx, addUserParams)
	if err != nil {
		return uuid.Nil, err
	}
	return userResult.ID, err
}

func (repository *UserSqlRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	updateUserParams := data.UpdateUserParams{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: &user.MobileNumber,
		ID:           user.ID,
	}

	result, err := repository.queries.UpdateUser(ctx, updateUserParams)
	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return user, nil
}

func (repository *UserSqlRepository) DeleteUser(ctx context.Context, user *entities.User) error {
	result, err := repository.queries.DeleteUser(ctx, user.ID)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
