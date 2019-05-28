package inc

import (
	"strconv"
	"sync"
)

var (
	innerDB  *Database
	lastElem uint64
	once     sync.Once
)

// User structure of the user data
type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

// Database structure of the users database
type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

// GetDBInstance singleton is used to get instance of the Database object
func GetDBInstance() *Database {
	once.Do(func() {
		// mutex := new(sync.RWMutex)
		// innerDB = &Database{*mutex, make(map[uint64]User)}
		innerDB = &Database{users: make(map[uint64]User)}
	})
	return innerDB
}

// Set method is used to add new users or modify existing ones
func (db *Database) Set(id, name, surname, email string) (User, bool) {
	db.m.Lock()
	defer db.m.Unlock()
	if id != "" {
		digitID, err := strconv.ParseUint(id, 10, 64)
		if err == nil {
			db.users[digitID] = User{digitID, name, surname, email}
			if usrget, ok := db.users[digitID]; ok {
				return usrget, true
			}
			return User{}, false
		}
		panic(err)
	} else {
		lastElem++
		db.users[lastElem] = User{lastElem, name, surname, email}
		if usrget, ok := db.users[lastElem]; ok {
			return usrget, true
		}
		return User{}, false
	}
}

// Get method is used to retrieve user data
func (db *Database) Get(usrID uint64) (User, bool) {
	db.m.Lock()
	defer db.m.Unlock()
	if usrget, ok := db.users[usrID]; ok {
		return usrget, true
	}
	return User{}, false
}

// Delete method is used to delete users
func (db *Database) Delete(usrID uint64) (string, bool) {
	db.m.Lock()
	defer db.m.Unlock()
	if usrdel, exists := db.users[usrID]; exists {
		delete(db.users, usrID)
		return usrdel.Name + " " + usrdel.Surname, true
	}
	return "Error: user doesn't exists", false
}
