package data

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDataStore(t *testing.T) {
	forNew := func() interface{} { return NewData() }
	forExisting := func() interface{} { return &Data{} }
	filename := "not-file.json"

	Convey("DataStore should flush file to disk", t, func() {
		dbname := ".db.json"
		d, _ := NewDataStore(dbname, forNew, forExisting)
		d.modified = true
		isFlushed, err := d.Flush()
		So(isFlushed, ShouldBeTrue)
		So(err, ShouldBeNil)
		_, err = os.Stat(dbname)
		So(os.IsNotExist(err), ShouldBeFalse)
	})

	Convey("New DataStore has min values", t, func() {
		d, nil := NewDataStore(filename, forNew, forExisting)
		So(nil, ShouldBeNil)
		So(d.modified, ShouldBeFalse)
		So(d.filename, ShouldEqual, filename)
		So(d.lock, ShouldNotBeNil)
	})
}
