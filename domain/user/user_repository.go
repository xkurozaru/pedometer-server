package user

import "github.com/google/uuid"

type UserRepository interface {
	Create(u User) error
	FindByUserID(userID UserID) (User, error)
	FindByAuthID(authID uuid.UUID) (User, error)
	FindAll() (Users, error)
	ExistsByUserID(userID UserID) (bool, error)
	Delete(u User) error
}
