package model

type User struct {
	UserId      string
	DateCreated string
	Password    string
}

func NewUser(userId, password, dateCreated string) *User {
	return &User{
		UserId:      userId,
		Password:    password,
		DateCreated: dateCreated,
	}
}

func (u *User) IsPasswordMatch(password string) bool {
	return u.Password == password
}
