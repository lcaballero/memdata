package data

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Item struct {
	Id           string
	CreatedOn    time.Time
	Title        string
	Summary      string
	OwnersId     string
	ItemState    ItemState
	RecordStatus RecordStatus
	Position     int
}

// NewItem provides an item with a new ID and with CreatedOn as Now().
// The Item's ItemState and RecordStatus are both set to defaults of
// Inception and Active (respectively).
func NewItem() Item {
	return Item{
		Id:           uuid.New(),
		CreatedOn:    time.Now(),
		ItemState:    Inception,
		RecordStatus: Active,
	}
}

// IsValid return true if the Item has an ID and a title, else false.
func (e Item) IsValid() bool {
	return e.Id != "" && e.Title != ""
}

// Update copies all fields from b (the source) to the receiver
// (the dest).
func (e *Item) Update(b Item) {
	e.Id = b.Id
	e.CreatedOn = b.CreatedOn
	e.Title = b.Title
	e.Summary = b.Summary
	e.OwnersId = b.OwnersId
	e.ItemState = b.ItemState
	e.RecordStatus = b.RecordStatus
	e.Position = b.Position
}

// Copy provides a new value with fields from the receiver.
func (e Item) Copy() Item {
	a := Item{}
	a.Update(e)
	return a
}
