/*Package usermodel*/
package usermodel

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/common"
	"github.com/jmoiron/sqlx"
)

// Users structure
type Users struct {
	DB *sqlx.DB
}

// User structure
type User struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

// NewUser constructor for Users
func NewUser(db *sqlx.DB) Users {
	return Users{DB: db}
}

// Register method for user registration logic
func (u Users) Register(user User) (User, error) {
	hashedPassword, err := common.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	sqlStatement := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id, password, created_at"
	err = u.DB.Get(&user, sqlStatement, user.Email, hashedPassword)
	return user, err
}

// Login method for user login logic
func (u Users) Login(user User) (User, error) {
	clearPassword := user.Password
	sqlStatement := "SELECT id, password, created_at FROM users WHERE email=$1"
	err := u.DB.Get(&user, sqlStatement, user.Email)
	if err != nil {
		return User{}, err
	}
	err = common.CheckPasswordHash(clearPassword, user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
