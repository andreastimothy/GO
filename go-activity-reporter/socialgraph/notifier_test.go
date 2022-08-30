package socialgraph_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/assignment-activity-reporter/socialgraph"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/assignment-activity-reporter/socialgraph/mocks"
)

func TestNotifyUpload(t *testing.T) {
	t.Run("should notify when others like your photo", func(t *testing.T) {
		test := mocks.NewFollowersNotifier(t)

		system := socialgraph.NewSystem()
		system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")
		user2.UploadPhoto()
		test.On("NotifyUpload", user2).Return()

		user2.AddFollowers(test)
		user2.UploadPhoto()
		user2.NotifyUploadFollowers()

		test.AssertExpectations(t)
		test.AssertNumberOfCalls(t, "NotifyUpload", 1)
	})
}
func TestNotifyLike(t *testing.T) {
	t.Run("should notify when others like your photo", func(t *testing.T) {
		test := mocks.NewFollowersNotifier(t)

		system := socialgraph.NewSystem()
		user1 := system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")
		user2.UploadPhoto()
		test.On("NotifyLike", user1, user2).Return()
		user1.AddFollowers(test)
		system.Like("Alice", "Bob")
		// user2.NotifyLikeFollowers(user1)

		test.AssertExpectations(t)
		test.AssertNumberOfCalls(t, "NotifyLike", 1)
	})
}
