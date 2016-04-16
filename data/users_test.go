package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func SomeUsers() Users {
	users := NewUsers(
		User{Id: "1", FirstName: "First"},
		User{Id: "2", FirstName: "Second"},
		User{Id: "3", FirstName: "Third"},
	)
	return users
}

func TestUsers(t *testing.T) {

	Convey("Copy clones the content wihtout reference", t, func() {
		users := SomeUsers()
		other := users.Copy()

		update := User{Id: "2", FirstName: "2nd"}
		a1, _ := users.Update(update)
		a2, _ := users.FindUser("2")

		b2, _ := other.FindUser("2")

		So(a1.FirstName, ShouldEqual, "2nd")
		So(a2.FirstName, ShouldEqual, "2nd")
		So(b2.FirstName, ShouldNotEqual, "2nd")
	})

	Convey("Update users with a new user by ID", t, func() {
		users := SomeUsers()
		update := User{Id: "2", FirstName: "2nd"}
		u, err := users.Update(update)
		So(err, ShouldBeNil)
		So(u.FirstName, ShouldEqual, "2nd")

		u2, e2 := users.FindUser("2")
		So(e2, ShouldBeNil)
		So(u2.FirstName, ShouldEqual, "2nd")
	})

	Convey("FindUser should by ID should find Name", t, func() {
		users := NewUsers(
			User{Id: "1", FirstName: "First"},
			User{Id: "2", FirstName: "Second"},
			User{Id: "3", FirstName: "Third"},
		)
		u, err := users.FindUser("2")
		So(err, ShouldBeNil)
		So(u.FirstName, ShouldEqual, "Second")
	})

	Convey("Adding User increases length", t, func() {
		users := NewUsers()
		u := User{Id: "id"}
		err := users.Add(u)
		So(err, ShouldBeNil)
		So(users, ShouldHaveLength, 1)
	})

	Convey("Empty Users collection should have length 0", t, func() {
		var users Users
		So(users, ShouldHaveLength, 0)
	})
}
