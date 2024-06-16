package user

type UserRepository interface {
	Create(user User) error
	Get(id string) (User, error)
	ExistsByUserID(userID string) (bool, error)
}
