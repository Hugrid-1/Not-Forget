package entity

type User struct {
	id       int
	login    string
	password string
	email    string
	verified bool
	premium  bool
}

// TODO: Обработка ошибок при создании пользователя
func NewUser(
	id int,
	login string,
	password string,
	email string,
	verified bool,
	premium bool,
) (User, error) {
	return User{
		id:       id,
		login:    login,
		password: password,
		email:    email,
		verified: verified,
		premium:  premium,
	}, nil
}

func (u *User) GetID() int {
	return u.id
}

func (u *User) GetLogin() string {
	return u.login
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) IsVerified() bool {
	return u.verified
}

func (u *User) IsPremium() bool {
	return u.premium
}
