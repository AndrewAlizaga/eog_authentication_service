package access

import (
	"errors"
	"fmt"
	"os"
	"time"

	Claim "github.com/AndrewAlizaga/eog_authentication_service/app/models/claim"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte(os.Getenv("EOG_JWT_KEY"))
var jwtKey = []byte(key)

func NewToken(payload Claim.ClaimUserObj, expDate time.Time) (string, error) {

	expirationTime := time.Now().Add(50 * time.Minute)

	claims := &Claim.Claim{
		Data: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//declaring the token with alg and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println("error")
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenData string) (bool, interface{}, error) {

	claims := &Claim.Claim{}
	println(tokenData)

	//error is in tokenData decoding
	tkn, err := jwt.ParseWithClaims(tokenData, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	fmt.Println("finished checking token parsing and signature")
	fmt.Println(tkn)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("signature error")
		}
		fmt.Println("error here, ", err)
		return false, nil, err
	}

	if !tkn.Valid {
		println("invalid token, token validation: ", tkn.Valid)
		return false, nil, errors.New("invalid token")
	}

	fmt.Println("JSON CLAIMS: ", tkn.Claims)

	fmt.Println("[DEBUG] PASS TKN VALIDATIONS")
	fmt.Println("CLAIMS: ", claims)
	fmt.Println(claims)

	fmt.Println(claims.Data)
	return true, claims, nil
}
