package user

type User struct {
	id       string
	userID   string
	username string
}

func NewUser(id string, userID string, username string) User {
	return newUser(id, userID, username)
}
func RecreateUser(id string, userID string, username string) User {
	return newUser(id, userID, username)
}
func newUser(id string, userID string, username string) User {
	return User{
		id:       id,
		userID:   userID,
		username: username,
	}
}

func (u User) ID() string {
	return u.id
}
func (u User) UserID() string {
	return u.userID
}
func (u User) Username() string {
	return u.username
}
