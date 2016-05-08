package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("FullUpdate should over-write properties", t, func() {
		src := NewUser()
		src.FirstName = "Bruce"
		dest := NewUser()
		So(src.Id, ShouldNotEqual, dest.Id)
		So(src.FirstName, ShouldNotEqual, dest.FirstName)

		(&dest).FullUpdateDest(src)
		So(src.Id, ShouldEqual, dest.Id)
		So(src.FirstName, ShouldEqual, dest.FirstName)
	})
}
