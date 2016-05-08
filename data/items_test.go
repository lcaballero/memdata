package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItems(t *testing.T) {

	Convey("Should not be able to add invalid item", t, func() {
		t := NewItem("")
		So(t.IsValid(), ShouldBeFalse)

		items := NewItems()
		So(items.Len(), ShouldEqual, 0)

		err := items.Add(t)
		So(err, ShouldNotBeNil)
	})

	Convey("Adding an item should increase length", t, func() {
		m := NewItems()
		err := m.Add(NewItem("id-1"))
		So(err, ShouldBeNil)
		So(m.Len(), ShouldEqual, 1)
	})

	Convey("NewItems produces empty collection", t, func() {
		m := NewItems()
		So(m, ShouldHaveLength, 0)
	})
}
