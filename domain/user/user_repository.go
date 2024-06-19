package user

type UserRepository interface {
	Create(user User) error
	Find(id string) (User, error)
	FindAll() (Users, error)
	ExistsByUserID(userID string) (bool, error)
}
