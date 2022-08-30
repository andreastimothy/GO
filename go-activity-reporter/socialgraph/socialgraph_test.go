package socialgraph_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/assignment-activity-reporter/socialgraph"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	t.Run("should add new user when initialized", func(t *testing.T) {
		system := socialgraph.NewSystem()

		system.Follow("Alice", "Bob")

		assert.NotNil(t, system.UserList)
		assert.Equal(t, 2, len(system.UserList))
	})

	t.Run("should add followers when followed", func(t *testing.T) {
		system := socialgraph.NewSystem()

		system.Follow("Alice", "Bob")

		user1 := system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")

		expected := true
		actual := system.IsAlreadyFollowed(user1, user2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should not add followers when already followed", func(t *testing.T) {
		system := socialgraph.NewSystem()

		system.Follow("Alice", "Bob")

		user1 := system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")

		expected := true
		actual := system.IsAlreadyFollowed(user1, user2)

		assert.Equal(t, expected, actual)

		system.Follow("Alice", "Bob")

		expectedFollowersCount := 1
		actualFollowersCount := len(system.Followers[user2])

		assert.Equal(t, actualFollowersCount, expectedFollowersCount)
	})
}

func TestLikeAndUpload(t *testing.T) {
	t.Run("should not like when photo not uploaded", func(t *testing.T) {
		system := socialgraph.NewSystem()

		system.Follow("Alice", "Bob")
		system.Like("Alice", "Bob")
		user := system.GetOrCreateUserByName("Bob")
		expected := 0
		actual := user.Likes

		assert.Equal(t, expected, actual)
	})

	t.Run("should increment like count when liked an uploaded photo", func(t *testing.T) {
		system := socialgraph.NewSystem()

		system.Follow("Alice", "Bob")
		system.Upload("Bob")
		system.Like("Alice", "Bob")
		user := system.GetOrCreateUserByName("Bob")
		expected := 1
		actual := user.Likes

		assert.Equal(t, expected, actual)
	})
}

func TestGetRanking(t *testing.T) {
	t.Run("should return sorted slice", func(t *testing.T) {
		system := socialgraph.NewSystem()

		user1 := system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")
		user3 := system.GetOrCreateUserByName("John")
		system.GetOrCreateUserByName("Bill")

		system.Upload("Alice")
		system.Upload("Bob")
		system.Upload("John")

		system.Like("Alice", "Alice")
		system.Like("Bob", "Alice")
		system.Like("John", "Alice")
		system.Like("Bill", "Alice")
		system.Like("Bill", "Bob")
		system.Like("John", "Bob")
		system.Like("Bill", "John")

		expected := []*socialgraph.User{user1, user2, user3}
		actual := system.GetRanking()
		assert.Equal(t, expected, actual)
	})

	t.Run("should return only 2 items when only 2 users upload photo", func(t *testing.T) {
		system := socialgraph.NewSystem()

		user1 := system.GetOrCreateUserByName("Alice")
		user2 := system.GetOrCreateUserByName("Bob")
		system.GetOrCreateUserByName("John")
		system.GetOrCreateUserByName("Bill")

		system.Upload("Alice")
		system.Upload("Bob")

		system.Like("Alice", "Alice")
		system.Like("Bob", "Alice")
		system.Like("John", "Alice")
		system.Like("Bill", "Alice")
		system.Like("Bill", "Bob")
		system.Like("John", "Bob")

		expected := []*socialgraph.User{user1, user2}
		actual := system.GetRanking()
		assert.Equal(t, expected, actual)
	})
}
