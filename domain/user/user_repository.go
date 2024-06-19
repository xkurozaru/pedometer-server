package user

type UserRepository interface {
	Create(user User) error
	FindByUserID(userID string) (User, error)
	FindByAuthID(authID string) (User, error)
	FindFollows(userID string) (Users, error)
	FindAll() (Users, error)
	ExistsByUserID(userID string) (bool, error)
}
