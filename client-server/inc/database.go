package inc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
)

var (
	innerDB  *Database
	lastElem uint64
	once     sync.Once
	// DataFile variable is used to specify the name of the database file
	DataFile = "database.json"
)

// DatabaseErrors string type
type DatabaseErrors string

// Error: application error handling method
func (e DatabaseErrors) Error() string {
	return fmt.Sprintf(string(e))
}

// Synchronizer boolean channel type for synchronization of work
type Synchronizer chan bool

// Stop method of Synchronizer type
func (s Synchronizer) Stop() {
	s <- bool(true)
}

// Resume method of Synchronizer type
func (s Synchronizer) Resume() {
	<-s
}

// User structure of the user data
type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

// Database structure of the users database
type Database struct {
	sync  Synchronizer
	users map[uint64]User
}

// GetDBInstance singleton is used to get instance of the Database object
func GetDBInstance() *Database {
	once.Do(func() {
		innerDB = &Database{sync: make(Synchronizer, 1), users: make(map[uint64]User)}
	})
	return innerDB
}

// Set method is used to add new users or modify existing ones
func (db *Database) Set(id, name, surname, email string) (User, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if id != "" {
		digitID, err := strconv.ParseUint(id, 10, 64)
		if err == nil {
			db.users[digitID] = User{digitID, name, surname, email}
			if usrget, ok := db.users[digitID]; ok {
				return usrget, nil
			}
			return User{}, DatabaseErrors("Database.Set: can't change user data")
		}
		return User{}, err
	}
	lastElem++
	db.users[lastElem] = User{lastElem, name, surname, email}
	if usrget, ok := db.users[lastElem]; ok {
		return usrget, nil
	}
	return User{}, DatabaseErrors("Database.Set: can't create new user")
}

// Get method is used to retrieve user data
func (db *Database) Get(usrID uint64) (User, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if usrget, ok := db.users[usrID]; ok {
		return usrget, nil
	}
	return User{}, DatabaseErrors("Database.Get: user doesn't exists")
}

// Delete method is used to delete users
func (db *Database) Delete(usrID uint64) (string, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	if usrdel, exists := db.users[usrID]; exists {
		delete(db.users, usrID)
		return usrdel.Name + " " + usrdel.Surname, nil
	}
	return "", DatabaseErrors("Database.Delete: user doesn't exists")
}

// SaveToFile method is used to save the database file
func (db *Database) SaveToFile(filename string) error {
	db.sync.Stop()
	defer db.sync.Resume()
	dbExport := make(map[uint64]map[uint64]User)
	dbExport[lastElem] = GetDBInstance().users
	content, err := json.Marshal(dbExport)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadFromFile method is used to load the database file
func (db *Database) LoadFromFile(filename string) ([]byte, error) {
	db.sync.Stop()
	defer db.sync.Resume()
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	dbImport := make(map[uint64]map[uint64]User)
	err = json.Unmarshal([]byte(content), &dbImport)
	if err != nil {
		return nil, err
	}
	for i := range dbImport {
		lastElem = i
		GetDBInstance().users = dbImport[i]
	}
	return content, nil
}
