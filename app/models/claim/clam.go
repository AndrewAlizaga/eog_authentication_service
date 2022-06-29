package claim

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Data ClaimUserObj `json:"data"`
	jwt.StandardClaims
}

type ClaimUserObj struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
}
