package whitelist

import "github.com/golang-jwt/jwt"

type Claim struct {
	MaxBody  int64 `json:"max_body"`
	RandomID int   `json:"random_id"`
	jwt.StandardClaims
}
