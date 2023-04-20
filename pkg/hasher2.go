package hasher2

import (
	"golang.org/x/crypto/bcrypt"
)

/*
 * Uses a string as an input parameter, returns a hash of this string and error if any
 */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
 * Requires two string parameters: a password and a hash to compare it against password. Returns true if password
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
