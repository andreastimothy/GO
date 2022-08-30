package socialgraph

type FollowersNotifier interface {
	NotifyLike(*User, *User)
	NotifyUpload(*User)
}