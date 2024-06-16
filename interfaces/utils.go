package interfaces

import (
	"errors"
	"net/http"
	"strings"
)

func GetToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authField := strings.Split(authHeader, " ")
	if len(authField) != 2 || authField[0] != "Bearer" {
		return "", errors.New("invalid authorization header")
	}
	return authField[1], nil
}
