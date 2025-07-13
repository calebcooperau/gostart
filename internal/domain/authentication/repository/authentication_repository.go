package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"gostart/internal/domain/authentication/data"
	"gostart/internal/domain/authentication/entities"
)

type AuthenticationRepository interface {
	FindUserIDByProvider(ctx context.Context, providerUserID string) (uuid.UUID, error)
	FindAuthUserByID(ctx context.Context, id uuid.UUID) (*entities.AuthUser, error)
	RegisterAuthUser(ctx context.Context, gothUser goth.User) (uuid.UUID, error)
}

type AuthenticationSqlRepository struct {
	queries *data.Queries
	db      *pgxpool.Pool
}

func NewAuthenticationSqlRepository(pool *pgxpool.Pool) *AuthenticationSqlRepository {
	queries := data.New(pool)
	return &AuthenticationSqlRepository{
		db:      pool,
		queries: queries,
	}
}

func (repository *AuthenticationSqlRepository) FindUserIDByProvider(ctx context.Context, providerUserID string) (uuid.UUID, error) {
	id, err := repository.queries.FindUserIDByProvider(ctx, providerUserID)

	if err == sql.ErrNoRows {
		return uuid.Nil, nil
	}
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (repository *AuthenticationSqlRepository) FindAuthUserByID(ctx context.Context, id uuid.UUID) (*entities.AuthUser, error) {
	authUserRow, err := repository.queries.FindAuthUserByID(ctx, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	authUser := &entities.AuthUser{
		ID:           authUserRow.ID,
		Email:        authUserRow.Email,
		FirstName:    authUserRow.FirstName,
		LastName:     authUserRow.LastName,
		MobileNumber: authUserRow.MobileNumber,
	}

	return authUser, nil
}

func (repository *AuthenticationSqlRepository) RegisterAuthUser(ctx context.Context, gothUser goth.User) (uuid.UUID, error) {
	authUserParams := data.CreateAuthUserParams{
		Email:          gothUser.Email,
		FirstName:      gothUser.FirstName,
		LastName:       gothUser.LastName,
		Provider:       gothUser.Provider,
		ProviderUserID: gothUser.UserID,
	}

	registeredID, err := repository.queries.CreateAuthUser(ctx, authUserParams)
	if err != nil {
		return uuid.Nil, err
	}

	return registeredID, err
}
