package friend_application

type FriendListDTO struct {
	Friends    []FriendDTO
	Requested  []FriendDTO
	Requesting []FriendDTO
}

type FriendDTO struct {
	FriendUserID   string
	FriendUsername string
}
