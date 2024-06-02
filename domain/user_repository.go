package domain

type UserRepository interface {
	Create(user *User) error
	FindUserById(id int) (*User, error)
	DeleteUser(id int) error
	UpdateUser(user *User) error
}
