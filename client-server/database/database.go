/*Package database Model*/
package database

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"strconv"
	"time"

	"github.com/andrii-stasiuk/go-exercises/client-server/errors"
	"github.com/andrii-stasiuk/go-exercises/client-server/synchro"
)

// User structure of the user data
type User struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// Database structure of the users database
type Database struct {
	sync       synchro.Synchronizer
	Users      map[uint64]User `json:"users"`
	datafile   string
	LastUserID uint64 `json:"lastUserID"`
}

// CreateDB function is used to get instance of the Database object
func CreateDB(filename string) *Database {
	innerDB := &Database{sync: make(synchro.Synchronizer, 1), Users: make(map[uint64]User), datafile: filename} // "database.json"
	return innerDB
}

// Set method is used to add new users or modify existing ones
func (db *Database) Set(id, name, surname, email string) (User, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if id != "" {
		digitID, err := strconv.ParseUint(id, 10, 64)
		if err == nil {
			db.Users[digitID] = User{digitID, name, surname, email}
			if usrget, ok := db.Users[digitID]; ok {
				return usrget, nil
			}
			return User{}, errors.DatabaseErrors{ErrorID: 1, ErrorText: "Database.Set: can't change user data"}
		}
		return User{}, err
	}
	db.LastUserID++
	db.Users[db.LastUserID] = User{db.LastUserID, name, surname, email}
	if usrget, ok := db.Users[db.LastUserID]; ok {
		return usrget, nil
	}
	return User{}, errors.DatabaseErrors{ErrorID: 2, ErrorText: "Database.Set: can't create new user"}
}

// Get method is used to retrieve user data
func (db Database) Get(usrID uint64) (User, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if usrget, ok := db.Users[usrID]; ok {
		return usrget, nil
	}
	return User{}, errors.DatabaseErrors{ErrorID: 3, ErrorText: "Database.Get: user doesn't exists"}
}

// Delete method is used to delete users
func (db *Database) Delete(usrID uint64) (string, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if usrdel, exists := db.Users[usrID]; exists {
		delete(db.Users, usrID)
		return usrdel.Name + " " + usrdel.Surname, nil
	}
	return "", errors.DatabaseErrors{ErrorID: 4, ErrorText: "Database.Delete: user doesn't exists"}
}

// SaveToFile method is used to save the database file
func (db Database) SaveToFile() error {
	db.sync.Stop()
	defer db.sync.Resume()
	// dbExport := make(map[uint64]map[uint64]User)
	// dbExport[db.lastUserID] = db.users
	content, err := json.Marshal(db)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(db.datafile, []byte(content), 0644); err != nil {
		return err
	}
	fmt.Println("The database was backed up at", time.Now())
	return nil
}

// LoadFromFile method is used to load the database file
func (db *Database) LoadFromFile() ([]byte, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	content, err := ioutil.ReadFile(db.datafile)
	if err != nil {
		return nil, err
	}
	// dbImport := make(map[uint64]map[uint64]User)
	// err = json.Unmarshal([]byte(content), &dbImport)
	if err = json.Unmarshal([]byte(content), &db); err != nil {
		return nil, err
	}
	// for i := range dbImport {
	// 	db.LastUserID = i
	// 	db.Users = dbImport[i]
	// }
	return content, nil
}

// RepeatSaving method is used to save the database at specific time intervals
func (db Database) RepeatSaving(safeDelay int) error {
	for {
		time.Sleep(time.Second * time.Duration(safeDelay))
		if err := db.SaveToFile(); err != nil {
			return err
		}
	}
}
