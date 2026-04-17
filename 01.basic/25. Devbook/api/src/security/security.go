package security

import "golang.org/x/crypto/bcrypt"

// Hash will generate a hash from a password
func Hash(userpass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userpass), bcrypt.DefaultCost)
}

// VerifyPassword will compare a hashed password with a password
func VerifyPassword(hashedPassword, userpass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userpass))
}
