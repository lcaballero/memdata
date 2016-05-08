package data

import (
	"errors"
)

type Items []Item

func NewItems(items ...Item) Items {
	return items
}

func (items Items) FindById(id string) (item Item, err error) {
	for _, t := range items {
		if t.Id == id {
			return t, nil
		}
	}
	return Item{}, errors.New("Couln't find item")
}

func (items Items) FindItem(id string) (Item, error) {
	for _, t := range items {
		if t.Id == id {
			return t, nil
		}
	}
	return Item{}, errors.New("Did not find Item")
}

func (items *Items) Add(more ...Item) error {
	for _, t := range more {
		if !t.IsValid() {
			return errors.New("Cannot add invalid item")
		}
	}
	*items = append(*items, more...)
	return nil
}

func (items Items) Len() int {
	return len(items)
}
func (items Items) UpdateState(id string, state ItemState) (item Item, err error) {
	for i := 0; i < len(items); i++ {
		t := &items[i]
		if t.Id == id {
			t.ItemState = state
			return *t, nil
		}
	}
	return Item{}, errors.New("Couldn't find item to update")
}

func (items Items) Update(update Item) (Item, error) {
	for i := 0; i < len(items); i++ {
		t := &items[i]
		if t.Id == update.Id {
			t.Title = update.Title
			t.Description = update.Description
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
