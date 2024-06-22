package user

type UserRepository interface {
	Create(u User) error
	FindByUserID(userID UserID) (User, error)
	FindByAuthID(authID string) (User, error)
	FindFollows(userID UserID) (Users, error)
	FindAll() (Users, error)
	ExistsByUserID(userID UserID) (bool, error)
	Delete(u User) error
}
