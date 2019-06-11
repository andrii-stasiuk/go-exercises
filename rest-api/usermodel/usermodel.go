/*Package usermodel*/
package usermodel

import (
	"database/sql"
	"log"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
)

type Users struct {
	Db *sql.DB
}

type User struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func New(db *sql.DB) Users {
	return Users{Db: db}
}

func (u Users) Register(user *User) (*User, error) {
	sqlStatement := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id, created_at"
	err := u.Db.QueryRow(sqlStatement, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	return user, err
}

func (u Users) Login(user *User) bool {
	pass := user.Password
	row := u.Db.QueryRow("SELECT id, email, password, created_at FROM users WHERE email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return false
	}
	log.Println(user.Password)
	return core.CheckPasswordHash(pass, user.Password)
	//hash https://gowebexamples.com/password-hashing
}
