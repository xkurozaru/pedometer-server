package friend_request

import "fmt"

type FriendRequestStatus string

const (
	FriendRequestStatusUnknown  FriendRequestStatus = "Unknown"
	FriendRequestStatusPending  FriendRequestStatus = "Pending"
	FriendRequestStatusAccepted FriendRequestStatus = "Accepted"
	FriendRequestStatusRejected FriendRequestStatus = "Rejected"
)

func FriendRequestStatusFromString(s string) FriendRequestStatus {
	switch s {
	case "Pending":
		return FriendRequestStatusPending
	case "Accepted":
		return FriendRequestStatusAccepted
	case "Rejected":
		return FriendRequestStatusRejected
	default:
		return FriendRequestStatusUnknown
	}
}

func (s FriendRequestStatus) ToString() string {
	return string(s)
}

func (s FriendRequestStatus) IsPending() bool {
	return s == FriendRequestStatusPending
}
func (s FriendRequestStatus) IsAccepted() bool {
	return s == FriendRequestStatusAccepted
}
func (s FriendRequestStatus) IsRejected() bool {
	return s == FriendRequestStatusRejected
}

func (s FriendRequestStatus) toAccept() (FriendRequestStatus, error) {
	switch s {
	case FriendRequestStatusPending:
		return FriendRequestStatusAccepted, nil
	default:
		return FriendRequestStatusUnknown, fmt.Errorf("cannot accept friend request with status %s", s)
	}
}

func (s FriendRequestStatus) toReject() (FriendRequestStatus, error) {
	switch s {
	case FriendRequestStatusPending:
		return FriendRequestStatusRejected, nil
	default:
		return FriendRequestStatusUnknown, fmt.Errorf("cannot reject friend request with status %s", s)
	}
}
