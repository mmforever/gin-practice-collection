package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var Jwts = make(map[string]string)

func JwtSetup() {
	Jwts["merchant"] = AppConfig.Merchant
	Jwts["rider"] = AppConfig.Rider
}

func Encode(name string, data interface{}) (t string) {
	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ = token.SignedString([]byte(Jwts[name]))
	return
}

func Decode(name string, tokeString string) (*jwt.MapClaims, string) {
	token, err := jwt.ParseWithClaims(tokeString, &jwt.MapClaims{}, Secret(name))
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, "token error"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, "token expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, "token not active yet"
			} else {
				return nil, "can't handle this token"
			}
		}

	}

	if token.Valid {
		return token.Claims.(*jwt.MapClaims), ""
	}

	return nil, "could'd handle this token"

}

func Secret(name string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(Jwts[name]), nil
	}
}
