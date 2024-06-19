package user

type User struct {
	userID   string
	username string
	authID   string
}

func NewUser(userID string, username string, authID string) User {
	return newUser(userID, username, authID)
}
func RecreateUser(userID string, username string, authID string) User {
	return newUser(userID, username, authID)
}
func newUser(userID string, username string, authID string) User {
	return User{
		userID:   userID,
		username: username,
		authID:   authID,
	}
}

func (u User) UserID() string {
	return u.userID
}
func (u User) Username() string {
	return u.username
}
func (u User) AuthID() string {
	return u.authID
}

type Users []User
