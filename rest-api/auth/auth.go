package auth

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
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

// Using asymmetric crypto/RSA keys
// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa"     // openssl genrsa -out app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// Verify key and sign key
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Read the key files before starting http handlers
func init() {
	var err error
	signKeyByte, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("Error reading private key: %v\n", err)
		return
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	if err != nil {
		log.Fatalf("Error parsing RSA private key: %v\n", err)
		return
	}
	verifyKeyByte, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("Error reading public key: %v\n", err)
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		log.Fatalf("Error parsing RSA public key: %v\n", err)
		return
	}
}

// Auth - middleware function for Authentication process
func Auth(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		// Grab the token from the header
		authToken := r.Header.Get("Authorization")
		authArr := strings.Split(authToken, " ")
		if len(authArr) != 2 {
			// Token is missing, returns with error code 401 Unauthorized
			log.Println("Missing auth token")
			responses.WriteErrorResponse(w, http.StatusUnauthorized, "Missing auth token")
			return
		}
		tknStr := authArr[1]
		claims := &Token{}
		// Validate the token
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				// Something was wrong during the validation
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					log.Printf("Token Expired: %+v\n", vErr.Errors)
					responses.WriteErrorResponse(w, http.StatusUnauthorized, "Token Expired, get a new one")
					return
				default:
					log.Printf("ValidationError error: %+v\n", vErr.Errors)
					responses.WriteErrorResponse(w, http.StatusForbidden, "Error while Parsing Token")
					return
				}
			default:
				// Something else went wrong
				log.Printf("Token parse error: %v\n", err)
				responses.WriteErrorResponse(w, http.StatusInternalServerError, "Malformed authentication token")
				return
			}
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
func GetToken(us usermodel.User) (string, error) {
	// Set the expire time
	expiresAt := time.Now().Add(time.Minute * 60 * 24) // expires in 24 hours
	// Set our claims
	tk := &Token{
		UserID: us.ID,
		Email:  us.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	// Create a signer for rsa 256
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), tk)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
