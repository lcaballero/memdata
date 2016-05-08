package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItem(t *testing.T) {

	Convey("NewItem without ID should be invalid", t, func() {
		t := NewItem("")
		So(t.IsValid(), ShouldBeFalse)
	})

	Convey("Update method should cause src to reflect dest.", t, func() {
		n := NewItem("test-id")
		n.Title = "Some title"
		n.Description = "Some summary"

		b := n.Copy()
		So(n.Title, ShouldEqual, b.Title)
		So(n.Description, ShouldEqual, b.Description)

		c := b.Copy()
		c.Title = "c title"
		So(b.Title, ShouldEqual, n.Title)
		So(c.Title, ShouldEqual, "c title")
	})

	Convey("Update method should cause src to reflect dest.", t, func() {
		n := NewItem("test-id")
		n.Title = "Some title"

		b := Item{
			Id:           n.Id,
			CreatedOn:    n.CreatedOn,
			ItemState:    n.ItemState,
			RecordStatus: n.RecordStatus,
			Title:        "Other title",
		}

		(&b).FullUpdateDest(n)
		So(n.Title, ShouldEqual, b.Title)
	})
}
