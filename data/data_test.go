package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestData(t *testing.T) {
	Convey("New Data has User and Items", t, func() {
		d := NewData()
		So(d.Users, ShouldNotBeNil)
		So(d.Items, ShouldNotBeNil)
	})
}
