package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItems(t *testing.T) {

	Convey("Adding an item should increase length", t, func() {
		m := NewItems()
		err := m.Add(Item{Id: "1"})
		So(err, ShouldBeNil)
		So(m.Len(), ShouldEqual, 1)
	})

	Convey("NewItems produces empty collection", t, func() {
		m := NewItems()
		So(m, ShouldHaveLength, 0)
	})
}
