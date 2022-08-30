package socialgraph

type User struct {
	Name      string
	Followers []FollowersNotifier
	Photos    bool
	Likes     int
	LikedBy   []*User
}

func NewUser(newName string) *User {
	return &User{Name: newName}
}

func (user *User) AddFollowers(follower FollowersNotifier) {
	user.Followers = append(user.Followers, follower)
}

func (user *User) NotifyLikeFollowers(u *User) {
	for _, notify := range user.Followers {
		notify.NotifyLike(user, u)
	}
}

func (user *User) LikePhoto(u *User) {
	user.NotifyLikeFollowers(u)
}

func (user *User) NotifyUploadFollowers() {
	for _, notify := range user.Followers {
		notify.NotifyUpload(user)
	}
}

func (user *User) UploadPhoto() {
	if !user.Photos {
		user.Photos = true
		user.NotifyUploadFollowers()
	}
}

func (user *User) GetName() string {
	return user.Name
}
