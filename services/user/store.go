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

}
func (s *Store) GetUserByID(id int) (*types.User, error) {}
