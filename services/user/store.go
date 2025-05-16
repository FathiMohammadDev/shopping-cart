package user

import (
	"database/sql"

	"github.com/FathiMohammadDev/shopping-cart/types"
)

type Store struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *Store {
	return &Store{db}
}


func(s *Store)GetUserByEmail(email string) (*types.User, error){}
func(s *Store)GetUserByID(id int) (*types.User, error){}
func(s *Store)CreateUser(user types.User) error{}