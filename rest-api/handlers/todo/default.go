/*Package handlers Todo*/
package todo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/julienschmidt/httprouter"
)

// User structure
type User struct {
	gorm.Model
	ID        uint64 `gorm:"id"`
	Email     string `gorm:"email"`
	Password  string `gorm:"password"`
	CreatedAt string `gorm:"created_at"`
}

// Todo main identifier
type Todo struct {
	//gorm.Model
	ID          int       `gorm:"id"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	State       string    `gorm:"state"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

var db *gorm.DB //database

// Default - handler for the root page /
func (h TodoHandlers) Default(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses.WriteOKResponse(w, "Welcome to API server!")
	// TEMPORARY CODE!!!
	// //Define DB connection string
	// dbURI := "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable"
	// //connect to db URI
	// db, err := gorm.Open("postgres", dbURI)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	panic(err)
	// }
	username := "testuser"
	password := "testpass"
	dbName := "testdb"
	dbHost := "localhost"
	dbPort := "5555"
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string
	fmt.Println(dbURI)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}
	// close db when not in use
	defer db.Close()
	db.Debug().AutoMigrate(&Todo{})
	fmt.Println("Successfully connected!", db)
	temp := &Todo{}
	err = db.Table("todos").Where("id = ?", 28).First(temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			println("Information not found")
		}
		println("Connection error. Please retry")
	}
	fmt.Println(*temp)
	// create
	todo := &Todo{Name: "asdf", Description: "ghjk", State: "1"}
	createdTodo := db.Create(todo)
	var errMessage = createdTodo.Error
	if createdTodo.Error != nil {
		fmt.Println(errMessage)
	}
	fmt.Println(createdTodo)
}
