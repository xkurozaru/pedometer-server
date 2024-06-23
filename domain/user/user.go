package user

import "github.com/google/uuid"

type UserID string

type User struct {
	userID   UserID
	username string
	authID   uuid.UUID
}

func NewUser(userID UserID, username string, authID uuid.UUID) User {
	return newUser(userID, username, authID)
}
func RecreateUser(userID UserID, username string, authID uuid.UUID) User {
	return newUser(userID, username, authID)
}
func newUser(userID UserID, username string, authID uuid.UUID) User {
	return User{
		userID:   userID,
		username: username,
		authID:   authID,
	}
}

func (u User) UserID() UserID {
	return u.userID
}
func (u User) Username() string {
	return u.username
}
func (u User) AuthID() uuid.UUID {
	return u.authID
}

type Users []User

func (u Users) UserIDs() []UserID {
	var ids []UserID
	for _, user := range u {
		ids = append(ids, user.UserID())
	}
	return ids
}
