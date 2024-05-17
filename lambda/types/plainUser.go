package types

type PlainUser struct {
	Username string `json:"username"`
	Password
}

func (plainUser PlainUser) ConvertToUser() (*User, error) {
	passwordHash, err := plainUser.Password.Hash()

	if err != nil {
		return nil, err
	}

	return &User{
		Username:     plainUser.Username,
		PasswordHash: *passwordHash,
	}, nil
}
