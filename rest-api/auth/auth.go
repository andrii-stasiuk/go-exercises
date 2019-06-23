package auth

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

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
		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")
		tokenCookie, err := r.Cookie("x-access-token")
		var tknStr string
		if len(authArr) == 2 {
			tknStr = authArr[1]
		} else if err == nil {
			tknStr = tokenCookie.Value
		} else {
			// Token is missing, returns with error code 401 Unauthorized
			log.Println("Missing auth token")
			responses.WriteErrorResponse(w, http.StatusUnauthorized, "Missing auth token")
			return
		}
		claims := &Token{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})
		if err != nil {
			// Malformed token, returns with http code 403 as usual
			log.Println("Malformed authentication token")
			responses.WriteErrorResponse(w, http.StatusForbidden, "Malformed authentication token")
			return
		}
		if !tkn.Valid {
			// Token is invalid, maybe not signed on this server
			log.Println("Token is not valid")
			responses.WriteErrorResponse(w, http.StatusUnauthorized, "Token is not valid")
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		fn(w, r.WithContext(ctx), param)
	}
}

// GetToken - function that creates new token for a logged in user
func GetToken(us usermodel.User) (string, time.Time) {
	expiresAt := time.Now().Add(time.Minute * 60 * 24) // expires in 24 hours
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
