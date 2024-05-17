package types

type User struct {
	Username     string       `json:"username"`
	PasswordHash PasswordHash `json:"passwordHash"`
}
