package whitelist

import (
	"errors"
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var ErrParse = errors.New("error parsing claims (cast)")

func GenerateToken(secret, audience, issuer string, maxBody int64) (string, error) {
	randomId := rand.Intn(9999999)
	claims := &Claim{
		// attributes
		MaxBody: maxBody,
		// generate random id
		RandomID: randomId,
		StandardClaims: jwt.StandardClaims{
			Audience: audience,
			IssuedAt: time.Now().Unix(),
			Issuer:   issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(secret []byte, signed string) (claims *Claim, err error) {
	fmt.Println(common.StyleDebug(), "checking token", signed, "with secret:", string(secret))
	var token *jwt.Token
	if token, err = jwt.ParseWithClaims(signed, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}); err != nil {
		return
	}
	var ok bool
	if claims, ok = token.Claims.(*Claim); !ok {
		return nil, ErrParse
	}
	return
}
