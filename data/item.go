package data

import (
	"time"
)

//go:generate genfront fields --output item.gen.go --template update_and_copy.gen.fm --line $GOLINE
type Item struct {
	Id        string
	GroupId   string
	CreatedOn time.Time
	CreatedBy string
	UpdatedOn time.Time
	UpdatedBy string

	Url          string
	Title        string
	Description  string
	Tags         []string
	FavoritedBy  []string
	Position     int
	Security     int // Unix like privacy
	ItemState    ItemState
	RecordStatus RecordStatus
}

// NewItem provides an item with a new ID and with CreatedOn as Now().
// The Item's ItemState and RecordStatus are both set to defaults of
// Inception and Active (respectively).
func NewItem(creatorId string) Item {
	now := time.Now()
	return Item{
		Id:           creatorId,
		CreatedOn:    now,
		UpdatedOn:    now,
		ItemState:    Inception,
		RecordStatus: Active,
	}
}

// IsValid return true if the Item has an ID and a title, else false.
func (e Item) IsValid() bool {
	return e.Id != "" && !e.UpdatedOn.IsZero() && !e.CreatedOn.IsZero()
}
