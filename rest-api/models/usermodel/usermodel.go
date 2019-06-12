/*Package usermodel*/
package usermodel

import (
	"database/sql"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
)

//
type Users struct {
	DB *sql.DB
}

//
type User struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

//
func New(db *sql.DB) Users {
	return Users{DB: db}
}

//
func (u Users) Register(user User) (User, error) {
	hashedPassword, err := core.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	sqlStatement := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id, created_at"
	err = u.DB.QueryRow(sqlStatement, user.Email, hashedPassword).Scan(&user.ID, &user.CreatedAt)
	return user, err
}

//
func (u Users) Login(user User) (User, bool) {
	clearPassword := user.Password
	sqlStatement := "SELECT id, email, password, created_at FROM users WHERE email=$1"
	err := u.DB.QueryRow(sqlStatement, user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err == nil && core.CheckPasswordHash(clearPassword, user.Password) == nil {
		return user, true
	}
	return User{}, false
}
