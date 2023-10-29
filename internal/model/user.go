package entity

type User struct {
	ID       int
	Login    string
	Password string
	Email    string
	Verified bool
}
