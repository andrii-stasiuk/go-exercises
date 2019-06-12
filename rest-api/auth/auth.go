package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

//Token struct declaration
type Token struct {
	UserID uint64
	Email  string
	*jwt.StandardClaims
}

// Auth - reserved for Authentication
func Auth(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		var header = r.Header.Get("x-access-token") //Grab the token from the header
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(`{"Message": "Missing auth token"}`)
			return
		}
		tk := &Token{}
		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk)
		fn(w, r.WithContext(ctx), param)
	}
}

//
func GetToken(us usermodel.User) map[string]interface{} {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := &Token{
		UserID: us.ID,
		Email:  us.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, error := token.SignedString([]byte("secret_key"))
	if error != nil {
		fmt.Println(error)
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = map[string]interface{}{"id": us.ID, "email": us.Email, "created_at": us.CreatedAt}
	return resp
}
