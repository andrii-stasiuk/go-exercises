/*Package usermodel*/
package usermodel

import (
	"time"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/jinzhu/gorm"
)

// Users structure
type Users struct {
	DB *gorm.DB
}

// User structure
type User struct {
	ID        uint64    `json:"id" gorm:"id"`
	Email     string    `json:"email" gorm:"email"`
	Password  string    `json:"password" gorm:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
}

// NewUser constructor for Users
func NewUser(db *gorm.DB) Users {
	return Users{DB: db}
}

// Register method for user registration logic
func (u Users) Register(user User) (User, error) {
	// hashedPassword, err := core.HashPassword(user.Password)
	// if err != nil {
	// 	return User{}, err
	// }
	// sqlStatement := "INSERT INTO users (email, password) VALUES($1, $2) RETURNING id, created_at"
	// err = u.DB.QueryRow(sqlStatement, user.Email, hashedPassword).Scan(&user.ID, &user.CreatedAt)
	// return user, err
	var err error
	user.Password, err = core.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	return user, u.DB.Create(&user).Error
}

// Login method for user login logic
func (u Users) Login(user User) (User, bool) {
	// clearPassword := user.Password
	// sqlStatement := "SELECT id, email, password, created_at FROM users WHERE email=$1"
	// err := u.DB.QueryRow(sqlStatement, user.Email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	// if err == nil && core.CheckPasswordHash(clearPassword, user.Password) == nil {
	// 	return user, true
	// }
	// return User{}, false
	clearPassword := user.Password
	err := u.DB.Where("email = ?", user.Email).First(&user).Error
	if err == nil && core.CheckPasswordHash(clearPassword, user.Password) == nil {
		return user, true
	}
	return User{}, false
}
