package friend

type FriendStatus string

const (
	FriendStatusRequesting  FriendStatus = "requesting"
	FriendStatusRequested   FriendStatus = "requested"
	FriendStatusEstablished FriendStatus = "established"
	FriendStatusUnknown     FriendStatus = "unknown"
)

func FriendStatusFromString(s string) FriendStatus {
	switch s {
	case FriendStatusRequesting.ToString():
		return FriendStatusRequesting
	case FriendStatusRequested.ToString():
		return FriendStatusRequested
	case FriendStatusEstablished.ToString():
		return FriendStatusEstablished
	default:
		return FriendStatusUnknown
	}
}

func (s FriendStatus) ToString() string {
	return string(s)
}

func (s FriendStatus) isRequesting() bool {
	return s == FriendStatusRequesting
}
func (s FriendStatus) isRequested() bool {
	return s == FriendStatusRequested
}
func (s FriendStatus) isEstablished() bool {
	return s == FriendStatusEstablished
}
