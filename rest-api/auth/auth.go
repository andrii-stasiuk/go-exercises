package auth

import (
	"context"
	"encoding/json"
	"io"
	"log"
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

// Auth - middleware function for Authentication process
func Auth(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		var header = r.Header.Get("x-access-token") // Grab the token from the header
		header = strings.TrimSpace(header)
		if header == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, `{"Error": "Missing auth token"}`)
			return
		}
		tk := &Token{}
		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(err)
			return
		}
		userKey := "user"
		ctx := context.WithValue(r.Context(), &userKey, tk)
		fn(w, r.WithContext(ctx), param)
	}
}

// GetToken - function that creates new token for a logged in user
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
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		log.Println(err)
	}
	resp := make(map[string]interface{})
	resp["token"] = tokenString // Store the token in the response
	resp["user"] = map[string]interface{}{"id": us.ID, "email": us.Email, "created_at": us.CreatedAt}
	return resp
}
