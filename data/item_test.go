package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItem(t *testing.T) {

	Convey("Update method should cause src to reflect dest.", t, func() {
		n := NewItem()
		n.Title = "Some title"
		n.Summary = "Some summary"

		b := n.Copy()
		So(n.Title, ShouldEqual, b.Title)
		So(n.Summary, ShouldEqual, b.Summary)

		c := b.Copy()
		c.Title = "c title"
		So(b.Title, ShouldEqual, n.Title)
		So(c.Title, ShouldEqual, "c title")
	})

	Convey("Update method should cause src to reflect dest.", t, func() {
		n := NewItem()
		n.Title = "Some title"

		b := Item{
			Id:           n.Id,
			CreatedOn:    n.CreatedOn,
			ItemState:    n.ItemState,
			RecordStatus: n.RecordStatus,
			Title:        "Other title",
		}

		n.Update(b)
		So(n.Title, ShouldEqual, b.Title)
	})
}
