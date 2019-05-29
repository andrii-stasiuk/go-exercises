package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// GetRequest function to get user data
func GetRequest(addr, id string) (map[string]interface{}, error) {
	resp, err := http.Get("http://" + addr + "/users/" + id + "/")
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// PostRequest function to add user or change user data
func PostRequest(addr, id, name, surname, email string) (map[string]interface{}, error) {
	formData := url.Values{
		"id":      {id},
		"name":    {name},
		"surname": {surname},
		"email":   {email},
	}
	resp, err := http.PostForm("http://"+addr+"/users/", formData)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// DeleteRequest function to delete user
func DeleteRequest(addr, id string) (map[string]interface{}, error) {
	// Create client
	client := http.Client{}
	// Create request
	request, err := http.NewRequest("DELETE", "http://"+addr+"/users/"+id+"/", nil)
	if err != nil {
		return nil, err
	}
	// Fetch Request
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	// Read Response Body
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
