package types

import "golang.org/x/crypto/bcrypt"

type Password string

func (password Password) Hash() (PasswordHash, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return PasswordHash(passwordHash), nil
}

type PasswordHash string

type PasswordValidator struct {
	Password     Password
	PasswordHash PasswordHash
}

func NewPasswordValidator(password Password, passwordHash PasswordHash) PasswordValidator {
	return PasswordValidator{
		Password:     password,
		PasswordHash: passwordHash,
	}
}

// Validate compares the password with the password hash and returns an error if they don't match
func (passwordValidator PasswordValidator) Validate() error {
	return bcrypt.CompareHashAndPassword([]byte(passwordValidator.PasswordHash), []byte(passwordValidator.Password))
}
