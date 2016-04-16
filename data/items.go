package data

import (
	"errors"
)

type Items []Item

func NewItems(items ...Item) Items {
	return items
}

func (items Items) FindItem(id string) (Item, error) {
	for _, t := range items {
		if t.Id == id {
			return t, nil
		}
	}
	return Item{}, errors.New("Did not find Item")
}

func (items *Items) Add(t ...Item) error {
	*items = append(*items, t...)
	return nil
}

func (items Items) Len() int {
	return len(items)
}

func (items Items) Update(update Item) (Item, error) {
	for i := 0; i < len(items); i++ {
		t := &items[i]
		if t.Id == update.Id {
			t.Title = update.Title
			t.Summary = update.Summary
			t.ItemState = update.ItemState
			return *t, nil
		}
	}
	return Item{}, errors.New("Couldn't find Item to update")
}

func (items Items) Copy() []Item {
	res := make([]Item, items.Len())
	copy(res, items)
	return res
}
