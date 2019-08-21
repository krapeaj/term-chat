package model

type User struct {
	UserId      string
	DateCreated string
	Accessible  []string
	password    string
}

func NewUser(userId, password, dateCreated string, accessible []string) *User {
	return &User{
		UserId:      userId,
		password:    password,
		DateCreated: dateCreated,
		Accessible:  accessible,
	}
}

func (u *User) IsPasswordMatch(password string) bool {
	return u.password == password
}
