package users

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	ListUsersByPagination(limit, offset *int) ([]User, error)
	GetUserByUsername(username string) (*User, error)
}
