package data

import (
	"errors"

	"github.com/lcaballero/memdata/da"
)

const DefaultDbName = "items.db.json"
type access func(*Data) bool

type AccessClient struct {
	Store *da.Locker
}

func NewAccessClient(dbname string) (*AccessClient, error) {
	store, err := da.NewDataStore(
		dbname,
		func() interface{} { return NewData().Add(NewUser()) },
		func() interface{} { return &Data{} },
	)
	if err != nil {
		return nil, err
	}
	a := &AccessClient{
		Store: store,
	}
	return a, nil
}

func (d *AccessClient) DataStore(fn access) {
	var usage da.Access = func(locked interface{}) bool {
		data, ok := locked.(*Data)
		if !ok {
			panic("Should have been working with a Data instance all along.")
		}
		return fn(data)
	}
	d.Store.DataStore(usage)
}

func (a *AccessClient) Users() (users []User) {
	a.DataStore(func(stored *Data) bool {
		users := stored.Users.Copy()
		return users != nil
	})
	return
}

func (a *AccessClient) AddUser(u User) error {
	a.DataStore(func(stored *Data) bool {
		err := stored.Users.Add(u)
		return err != nil
	})
	return nil
}

func (a *AccessClient) UpdateUser(updated User) (user User, err error) {
	a.DataStore(func(stored *Data) bool {
		user, err = stored.Users.Update(updated)
		return err != nil
	})
	return user, err
}

func (a *AccessClient) FindUser(id string) (user User, err error) {
	a.DataStore(func(store *Data) bool {
		user, err = store.Users.FindUser(id)
		return err != nil
	})
	return user, err
}

func (a *AccessClient) FindItem(id string) (item Item, err error) {
	a.DataStore(func(store *Data) bool {
		for _, t := range store.Items {
			if t.Id == id {
				item, err = t, nil
				return false
			}
		}
		item, err = Item{}, errors.New("Couln't find item")
		return false
	})
	return
}

func (a *AccessClient) UpdateItem(id string, updated Item) (item Item, err error) {
	a.DataStore(func(storage *Data) bool {
		for i := 0; i < len(storage.Items); i++ {
			t := &storage.Items[i]
			if t.Id == id {
				t.Title = updated.Title
				t.Summary = updated.Summary
				item, err = *t, nil
				return true
			}
		}
		item, err = Item{}, errors.New("Couldn't find item to update")
		return false
	})
	return
}

func (a *AccessClient) UpdateItemState(id string, state ItemState) (item Item, err error) {
	a.DataStore(func(stored *Data) bool {
		for i := 0; i < len(stored.Items); i++ {
			t := &stored.Items[i]
			if t.Id == id {
				t.ItemState = state
				item, err = *t, nil
				return true
			}
		}
		item, err = Item{}, errors.New("Couldn't find item to update")
		return false
	})
	return
}

func (a *AccessClient) AddItem(item Item) error {
	a.DataStore(func(stored *Data) bool {
		stored.Items = append(stored.Items, item)
		return true
	})
	return nil
}

func (d *AccessClient) Items() (res []Item) {
	d.DataStore(func(stored *Data) bool {
		res = make([]Item, len(stored.Items))
		copy(res, stored.Items)
		return false
	})
	return
}
