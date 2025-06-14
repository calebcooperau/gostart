package authentication

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

type AuthenticationRepository interface {
	FindUserIdByProvider(provider string, providerUserId string) (uuid.UUID, error)
	GetAuthUserById(id uuid.UUID) (*AuthUser, error)
	RegisterAuthUser(gothUser goth.User) (uuid.UUID, error)
}

type AuthenticationSqlRepository struct {
	db *sql.DB
}

func NewAuthenticationSqlRepository(db *sql.DB) *AuthenticationSqlRepository {
	return &AuthenticationSqlRepository{
		db: db,
	}
}

func (repository *AuthenticationSqlRepository) FindUserIdByProvider(provider string, providerUserId string) (uuid.UUID, error) {
	var userId uuid.UUID
	query := `
	SELECT user_id
	FROM auth_user_providers
	WHERE provider_user_id = $1
	`

	err := repository.db.QueryRow(query, providerUserId).Scan(&userId)
	if err == sql.ErrNoRows {
		return uuid.Nil, nil
	}
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

func (repository *AuthenticationSqlRepository) GetAuthUserById(id uuid.UUID) (*AuthUser, error) {
	usr := &AuthUser{}
	query := `
	SELECT id, email, first_name, last_name, mobile_number, created_at, updated_at
	FROM users
	WHERE id = $1
	`

	err := repository.db.QueryRow(query, id).Scan(
		&usr.Id,
		&usr.Email,
		&usr.CreatedAt,
		&usr.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (repository *AuthenticationSqlRepository) RegisterAuthUser(gothUser goth.User) (uuid.UUID, error) {
	transaction, err := repository.db.Begin()
	if err != nil {
		return uuid.Nil, err
	}
	defer transaction.Rollback()

	// Create Auth User record
	query := `
	INSERT INTO auth_users (email, first_name, last_name)	
	VALUES ($1, $2, $3)
	RETURNING id 
	`

	var authUserId uuid.UUID
	err = repository.db.QueryRow(query, gothUser.Email, gothUser.FirstName, gothUser.LastName).Scan(&authUserId)
	if err != nil {
		return uuid.Nil, err
	}

	// Create Auth Provider record
	query = `
	INSERT INTO auth_user_providers(user_id, provider, provider_user_id)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	var providerId uuid.UUID
	err = repository.db.QueryRow(query, authUserId, gothUser.Provider, gothUser.UserID).Scan(&providerId)
	if err != nil {
		return uuid.Nil, err
	}

	query = `
	INSERT INTO users (id, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`

	err = repository.db.QueryRow(query, authUserId, gothUser.Email, gothUser.FirstName, gothUser.LastName).Scan()
	if err != nil {
		return uuid.Nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return uuid.Nil, err
	}
	return authUserId, err
}
