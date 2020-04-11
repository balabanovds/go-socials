package data

import (
	"testing"
)

var user User

func TestUser_Create(t *testing.T) {
	setup()
	id, err := users[0].Create()
	handleErr(err, t)
	user, err = UserById(id)
	handleErr(err, t)
	if user.Email != users[0].Email {
		t.Errorf("email exp: %v, got %v", users[0].Email, user.Email)
	}
}

func TestUser_CreateSession(t *testing.T) {
	TestUser_Create(t)

	session, err := user.CreateSession()
	handleErr(err, t)
	if session.Id == 0 {
		t.Errorf("wrong session %#v", session)
	}
	if session.UserId != user.Id {
		t.Errorf(`exp %v, got %v`, session.UserId, user.Id)
	}

}

func TestUser_CreateProfile(t *testing.T) {
	TestUser_Create(t)
	pr, err := user.CreateProfile(profile)
	handleErr(err, t)
	if pr.UserId != user.Id {
		t.Errorf(`exp %v, got %v`, pr.UserId, user.Id)
	}
}

func TestUser_CreateProfileTwiceWithError(t *testing.T) {
	TestUser_CreateProfile(t)
	_, err := user.CreateProfile(profile2)
	if err == nil {
		t.Error("expected error")
	}

}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
