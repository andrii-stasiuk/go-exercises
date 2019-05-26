package inc

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// GetRequest function to get user data
func GetRequest(addr, id string) map[string]interface{} {
	resp, err := http.Get("http://" + addr + "/users/" + id + "/")
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

// PostRequest function to add user or change user data
func PostRequest(addr, id, name, surname, email string) map[string]interface{} {
	formData := url.Values{
		"id":      {id},
		"name":    {name},
		"surname": {surname},
		"email":   {email},
	}
	resp, err := http.PostForm("http://"+addr+"/users/", formData)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

// DeleteRequest function to delete user
func DeleteRequest(addr, id string) map[string]interface{} {
	// Create client
	client := http.Client{}
	// Create request
	request, err := http.NewRequest("DELETE", "http://"+addr+"/users/"+id+"/", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// Fetch Request
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	// Read Response Body
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}
