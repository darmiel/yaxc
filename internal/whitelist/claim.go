package whitelist

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	MaxBody  int64 `json:"max_body"`
	RandomID int   `json:"random_id"`
	jwt.StandardClaims
}
