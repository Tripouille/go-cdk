package types

type PlainUser struct {
	Username string   `json:"username"`
	Password Password `json:"password"`
}

func (plainUser PlainUser) ConvertToUser() (User, error) {
	user := User{}

	passwordHash, err := plainUser.Password.Hash()

	if err != nil {
		return user, err
	}

	user.Username = plainUser.Username
	user.PasswordHash = passwordHash
	return user, nil
}
