package inc

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

// DataFile variable is used to specify the name of the database file
var DataFile = "database.json"

// CheckIP function is used to verify the correctness of the IP-address
// and setting default value in case of error
func CheckIP(addrFlag string) string {
	if net.ParseIP(addrFlag) != nil {
		return addrFlag
	}
	return "127.0.0.1"
}

// SaveToFile function is used to save the database file
func SaveToFile(filename string) error {
	dbExport := make(map[uint64]map[uint64]User)
	dbExport[lastElem] = GetDBInstance().users
	content, err := json.Marshal(dbExport)
	if err != nil {
		return err
	} else {
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
}

// LoadFromFile function is used to load the database file
func LoadFromFile(filename string) ([]byte, bool) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	} else {
		dbImport := make(map[uint64]map[uint64]User)
		err := json.Unmarshal([]byte(content), &dbImport)
		if err != nil {
			panic(err)
		} else {
			for i := range dbImport {
				lastElem = i
				GetDBInstance().users = dbImport[i]
			}
			return content, true
		}
	}
}
