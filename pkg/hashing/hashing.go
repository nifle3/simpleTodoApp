package hashing

import "golang.org/x/crypto/bcrypt"

const (
	cost = 14
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}

func CheckHashPassword(hashing string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashing), []byte(password))
}
