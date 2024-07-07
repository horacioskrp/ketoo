package user

import (
	"database/sql"
	"fmt"

	"github.com/horacioskrp/kitoo/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not Found")
	}
	return u, nil
}

func (s *Store) GetUserByUSername(username string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("User not Found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
