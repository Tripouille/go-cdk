package types

import "golang.org/x/crypto/bcrypt"

type Password struct {
	Password string `json:"password"`
}

func (password Password) Hash() (*PasswordHash, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password.Password), 13)

	if err != nil {
		return nil, err
	}

	return &PasswordHash{
		PasswordHash: string(passwordHash),
	}, nil
}

type PasswordHash struct {
	PasswordHash string `json:"passwordHash"`
}

type PasswordValidator struct {
	Password
	PasswordHash
}

func NewPasswordValidator(password Password, passwordHash PasswordHash) PasswordValidator {
	return PasswordValidator{
		Password:     password,
		PasswordHash: passwordHash,
	}
}

func (passwordValidator PasswordValidator) Validate() error {
	return bcrypt.CompareHashAndPassword([]byte(passwordValidator.PasswordHash.PasswordHash), []byte(passwordValidator.Password.Password))
}
