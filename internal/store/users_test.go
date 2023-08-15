package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.NotEmpty(t, user.Salt)
	assert.NotEmpty(t, user.HashedPassword)
}

// test that we can add for user account creation is when user tries to create account with already existing username
func TestAddUserWithExistingUsername(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)

	user, err = addTestUser()
	assert.Error(t, err)
	assert.Equal(t, "Username already exists.", err.Error())
}

// To test Authenticate() function, we will create 3 tests: successful authentication, authenticating with invalid username, and authenticating with invalid password:
func TestAuthenticateUser(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	authUser, err := Authenticate(user.Username, user.Password)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, authUser.ID)
	assert.Equal(t, user.Username, authUser.Username)
	assert.Equal(t, user.Salt, authUser.Salt)
	assert.Equal(t, user.HashedPassword, authUser.HashedPassword)
	assert.Empty(t, authUser.Password)
}

func TestAuthenticateUserInvalidUsername(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	authUser, err := Authenticate("invalid", user.Password)
	assert.Error(t, err)
	assert.Nil(t, authUser)
}

func TestAuthenticateUserInvalidPassword(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	authUser, err := Authenticate(user.Username, "invalid")
	assert.Error(t, err)
	assert.Nil(t, authUser)
}

// we will test FetchUser() function with 2 tests: successfull fetch, and fetching not existing user:
func TestFetchUser(t *testing.T) {
	testSetup()
	user, err := addTestUser()
	assert.NoError(t, err)

	fetchedUser, err := FetchUser(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Username, fetchedUser.Username)
	assert.Empty(t, fetchedUser.Password)
	assert.Equal(t, user.Salt, fetchedUser.Salt)
	assert.Equal(t, user.HashedPassword, fetchedUser.HashedPassword)
}

func TestFetchNotExistingUser(t *testing.T) {
	testSetup()

	fetchedUser, err := FetchUser(1)
	assert.Error(t, err)
	assert.Nil(t, fetchedUser)
	assert.Equal(t, "Not found.", err.Error())
}
