package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenJWT(iss, sub, aud string, jwtKey []byte) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"sub": sub,
		"aud": aud,
		"iat": now.Unix(),
		"exp": now.Add(5 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string, jwtKey []byte) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
