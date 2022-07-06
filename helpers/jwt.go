package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const secretkey = "S3cReTkEy"

func GenerateToken(username string, id int) string {
	claims := jwt.MapClaims{
		"username": username,
		"id":       id,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretkey))

	return signedToken
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	var err = errors.New("failed to verify token")
	parseToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secretkey), nil
	})
	if _, ok := parseToken.Claims.(jwt.MapClaims); !ok && !parseToken.Valid {
		return nil, err
	}
	return parseToken.Claims.(jwt.MapClaims), nil
}
