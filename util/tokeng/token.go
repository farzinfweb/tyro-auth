package tokeng

import "github.com/golang-jwt/jwt/v5"

func Generate(secret string) (string, error) {
	tokenGenerator := jwt.New(jwt.SigningMethodHS256)

	token, err := tokenGenerator.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
