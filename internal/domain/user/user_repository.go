package user

import (
	"database/sql"

	"github.com/google/uuid"
	_ "golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*User, error)
	RegisterUser(cmp *User) (uuid.UUID, error)
	UpdateUser(cmp *User) (*User, error)
	DeleteUser(cmp *User) error
}

type UserSqlRepository struct {
	db *sql.DB
}

func NewUserSqlRepository(db *sql.DB) *UserSqlRepository {
	return &UserSqlRepository{db: db}
}

func (repository *UserSqlRepository) RegisterUser(user *User) (uuid.UUID, error) {
	query := `
	INSERT INTO users (email, first_name, last_name, mobile_number)	
	VALUES ($1, $2, $3, $4,)
	RETURNING id, created_at, updated_at
	`

	err := repository.db.QueryRow(query, user.Email, user.FirstName, user.LastName, user.MobileNumber).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		// will update
		return uuid.Nil, err
	}
	return user.Id, err
}

func (repository *UserSqlRepository) GetUserById(id uuid.UUID) (*User, error) {
	user := &User{}
	query := `
	SELECT id, email, first_name, last_name, mobile_number, created_at, updated_at
	FROM users
	WHERE id = $1
	`

	err := repository.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.MobileNumber,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserSqlRepository) UpdateUser(user *User) (*User, error) {
	query := `
	UPDATE users 
	SET email = $1, first_name = $2, last_name = $3, mobile_number = $4, updated_at = CURRENT_TIMESTAMP
	WHERE id = $5
	RETURNING updated_at
	`

	result, err := repository.db.Exec(query, user.Email, user.FirstName, user.LastName, user.MobileNumber, user.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return user, nil
}

func (repository *UserSqlRepository) DeleteUser(user *User) error {
	query := `DELETE from users WHERE id = $1`

	result, err := repository.db.Exec(query, user.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
