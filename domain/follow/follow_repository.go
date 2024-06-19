package follow

type FollowRepository interface {
	Upsert(follow Follow) error
	Delete(follow Follow) error
}
