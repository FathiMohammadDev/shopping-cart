package user

import (
	"database/sql"
	"fmt"

	"github.com/FathiMohammadDev/shopping-cart/types"
)

type Store struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.DB.Exec("INSER INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.DB.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)

	if rows.Next() {
		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found.")
	}
	return u, nil

}
func (s *Store) GetUserByID(id int) (*types.User, error) {}
