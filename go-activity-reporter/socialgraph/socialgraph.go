package socialgraph

import (
	"fmt"
	"sort"
)

type System struct {
	UserList   []*User
	Followers  map[*User][]*User
	Activities map[*User][]string
}

func NewSystem() *System {
	return &System{
		UserList:   make([]*User, 0),
		Followers:  make(map[*User][]*User),
		Activities: make(map[*User][]string),
	}
}

func (system *System) GetOrCreateUserByName(name string) *User {
	for _, value := range system.UserList {
		if value.GetName() == name {
			return value
		}
	}
	newUser := NewUser(name)
	system.UserList = append(system.UserList, newUser)
	newUser.AddFollowers(system)
	return newUser
}

func (system *System) NotifyLike(user1 *User, user2 *User) {
	activity := fmt.Sprint(user1.GetName(), " liked ", user2.GetName(), "'s photo")
	if user1.GetName() == user2.GetName() {
		selfactivity := "You liked your photo"
		system.Activities[user1] = append(system.Activities[user1], selfactivity)
		for _, followers := range system.Followers[user1] {
			system.Activities[followers] = append(system.Activities[followers], activity)
		}
	} else {
		selfactivity := fmt.Sprint("You liked ", user2.GetName(), "'s photo")
		system.Activities[user1] = append(system.Activities[user1], selfactivity)
		flag := false
		for _, followers := range system.Followers[user1] {
			if user2 == followers {
				flag = true
				system.Activities[user2] = append(system.Activities[user2], user1.GetName()+" liked your photo")
			} else {
				system.Activities[followers] = append(system.Activities[followers], activity)
			}
		}
		if !flag {
			system.Activities[user2] = append(system.Activities[user2], user1.GetName()+" liked your photo")
		}
	}
}

func (system *System) NotifyUpload(user *User) {
	activity := fmt.Sprint(user.GetName(), " uploaded photo")
	selfactivity := "You uploaded photo"
	system.Activities[user] = append(system.Activities[user], selfactivity)
	for _, followers := range system.Followers[user] {
		system.Activities[followers] = append(system.Activities[followers], activity)
	}
}

func (system *System) IsAlreadyFollowed(user1, user2 *User) bool {
	for _, value := range system.Followers[user2] {
		if value == user1 {
			return true
		}
	}
	return false
}

func (system *System) Follow(user1, user2 string) {
	userfollow := system.GetOrCreateUserByName(user1)
	userfollowed := system.GetOrCreateUserByName(user2)
	if system.IsAlreadyFollowed(userfollow, userfollowed) {
		return
	}
	system.Followers[userfollowed] = append(system.Followers[userfollowed], userfollow)
}

func (system *System) IsUserExist(user string) (*User, error) {
	for _, userexist := range system.UserList {
		if userexist.GetName() == user {
			return userexist, nil
		}
	}
	return nil, UserDidNotExist{}
}

func IsAlreadyLiked(user1, user2 *User) bool {
	for _, value := range user2.LikedBy {
		if value == user1 {
			return true
		}
	}
	return false
}

func (system *System) Like(user1, user2 string) {
	userlike, err1 := system.IsUserExist(user1)
	userliked, err2 := system.IsUserExist(user2)
	if err1 == nil && err2 == nil && !IsAlreadyLiked(userlike, userliked) && userliked.Photos {
		userlike.LikePhoto(userliked)
		userlike.LikedBy = append(userlike.LikedBy, userlike)
		userliked.Likes++
	}
}

func (system *System) Upload(user string) {
	for _, i := range system.UserList {
		if i.GetName() == user {
			i.UploadPhoto()
		}
	}
}

func (system *System) GetRanking() []*User {
	var unsorted []*User
	for _, value := range system.UserList {
		if value.Photos {
			unsorted = append(unsorted, value)
		}
	}
	sort.SliceStable(unsorted, func(i, j int) bool {
		return unsorted[i].Likes > unsorted[j].Likes
	})
	if len(unsorted) >= 3 {
		return unsorted[0:3]
	} else {
		return unsorted
	}
}
