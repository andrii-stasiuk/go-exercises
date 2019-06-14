package auth

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// Token struct declaration
type Token struct {
	UserID uint64
	Email  string
	*jwt.StandardClaims
}

// Set up a global string for signing key
var signingKey = []byte("secret_key")

// Auth - middleware function for Authentication process
func Auth(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		// Grab the token from the header
		header := strings.TrimSpace(r.Header.Get("x-access-token"))
		tokenCookie, err := r.Cookie("x-access-token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				log.Println("Unauthorized")
				responses.WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			// For any other type of error, return a bad request status
			log.Println("Bad Request")
			responses.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request")
			return
		}
		if !core.CheckStr(header) && !core.CheckStr(tokenCookie.Value) {
			// Token is missing, returns with error code 403 Unauthorized
			log.Println("Missing auth token")
			responses.WriteErrorResponse(w, http.StatusForbidden, "Missing auth token")
			return
		}
		var tknStr string
		if core.CheckStr(header) {
			tknStr = header
		} else {
			tknStr = tokenCookie.Value
		}
		tk := &Token{}
		_, err = jwt.ParseWithClaims(tknStr, tk, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			log.Println(err)
			responses.WriteErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}
		userKey := "user"
		ctx := context.WithValue(r.Context(), &userKey, tk)
		fn(w, r.WithContext(ctx), param)
	}
}

// GetToken - function that creates new token for a logged in user
func GetToken(us usermodel.User) (string, time.Time) /*map[string]interface{}*/ {
	expiresAt := time.Now().Add(time.Minute * 100000)
	tk := &Token{
		UserID: us.ID,
		Email:  us.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Println(err)
		return "", time.Time{}
	}
	return tokenString, expiresAt
}
