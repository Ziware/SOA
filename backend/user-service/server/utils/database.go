package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDuplicateKey = errors.New("user login or email mistmatch")
	ErrInvalidData  = errors.New("incorrect user data")
)

func NewDatabase(config TDBConfig) (*TDatabase, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to posgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging to posgres: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("successfull init of posgres")
	return &TDatabase{db}, nil
}

func (r *TDatabase) Close() error {
	return r.db.Close()
}

func (r *TDatabase) GetUserByID(id uuid.UUID) (*TUser, error) {
	query := `
		SELECT 
			user_id, login, email, password_hash, 
			name, surname, birth_date, phone_number,
			created_at, updated_at
		FROM users.users 
		WHERE user_id = $1
	`

	var user TUser

	err := r.db.QueryRow(query, id).Scan(
		&user.UserId, &user.Login, &user.Email, &user.PassHash,
		&user.Name, &user.Surname, &user.BirthDate, &user.PhoneNumber,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by id error: %w", err)
	}

	return &user, nil
}

func (r *TDatabase) GetUserByLogin(login string) (*TUser, error) {
	query := `
		SELECT 
			user_id, login, email, password_hash, 
			name, surname, birth_date, phone_number,
			created_at, updated_at
		FROM users.users 
		WHERE login = $1
	`

	var user TUser

	err := r.db.QueryRow(query, login).Scan(
		&user.UserId, &user.Login, &user.Email, &user.PassHash,
		&user.Name, &user.Surname, &user.BirthDate, &user.PhoneNumber,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by login error: %w", err)
	}

	return &user, nil
}

func (r *TDatabase) GetUserByEmail(email string) (*TUser, error) {
	query := `
		SELECT 
			user_id, login, email, password_hash, 
			name, surname, birth_date, phone_number,
			created_at, updated_at
		FROM users.users 
		WHERE email = $1
	`

	var user TUser

	err := r.db.QueryRow(query, email).Scan(
		&user.UserId, &user.Login, &user.Email, &user.PassHash,
		&user.Name, &user.Surname, &user.BirthDate, &user.PhoneNumber,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by email error: %w", err)
	}

	return &user, nil
}

func (r *TDatabase) CreateUser(user *TUser) error {
	if user.Login == "" || user.Email == "" || user.PassHash == "" {
		return ErrInvalidData
	}

	query := `
		INSERT INTO users.users (
			login, email, password_hash, 
			name, surname, birth_date, phone_number,
			user_id, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(
		query,
		user.Login, user.Email, user.PassHash,
		user.Name, user.Surname, user.BirthDate, user.PhoneNumber,
		user.UserId, user.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("insertion error: %w", err)
	}

	return nil
}

func (r *TDatabase) UpdateUser(user *TUser) error {
	exists, err := r.CheckUserExistsById(user.UserId)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}

	query := `
		UPDATE users.users SET 
			email = $1, 
			name = $2, 
			surname = $3, 
			birth_date = $4, 
			phone_number = $5
		WHERE user_id = $6
		RETURNING updated_at
	`

	err = r.db.QueryRow(
		query,
		user.Email, user.Name, user.Surname,
		user.BirthDate, user.PhoneNumber, user.UserId,
	).Scan(&user.UpdatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return ErrDuplicateKey
		}
		return fmt.Errorf("update user error: %w", err)
	}

	return nil
}

func (r *TDatabase) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users.users WHERE user_id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error on user deletion: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error on user deletion: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *TDatabase) ListUsers(limit, offset int) ([]TUser, error) {
	query := `
		SELECT 
			user_id, login, email, password_hash, 
			name, surname, birth_date, phone_number,
			created_at, updated_at
		FROM users.users 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error on user lists: %w", err)
	}
	defer rows.Close()

	var users []TUser

	for rows.Next() {
		var user TUser

		if err := rows.Scan(
			&user.UserId, &user.Login, &user.Email, &user.PassHash,
			&user.Name, &user.Surname, &user.BirthDate, &user.PhoneNumber,
			&user.CreatedAt, &user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error on user lists: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on user lists: %w", err)
	}

	return users, nil
}

func (r *TDatabase) CountUsers() (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users.users`

	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error on users count: %w", err)
	}

	return count, nil
}

func (r *TDatabase) CheckUserExistsById(id uuid.UUID) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users.users WHERE user_id = $1)`

	err := r.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error on user existance check by id: %w", err)
	}

	return exists, nil
}

func (r *TDatabase) CheckUserExistsByLogin(login string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users.users WHERE login = $1)`

	err := r.db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error on user existance check by login: %w", err)
	}

	return exists, nil
}
