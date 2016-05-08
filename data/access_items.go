package data

type AccessItems struct {
	accessor LockerAccess
}

func (a *AccessItems) FindById(id string) (item Item, err error) {
	a.accessor(func(store *Data) bool {
		item, err = store.Items.FindById(id)
		return err != nil
	})
	return
}

func (a *AccessItems) Update(id string, updated Item) (item Item, err error) {
	a.accessor(func(storage *Data) bool {
		item, err = storage.Items.Update(item)
		return err != nil
	})
	return
}

func (a *AccessItems) UpdateState(id string, state ItemState) (item Item, err error) {
	a.accessor(func(stored *Data) bool {
		item, err = stored.Items.UpdateState(id, state)
		return err != nil
	})
	return
}

func (a *AccessItems) Add(item Item) error {
	a.accessor(func(stored *Data) bool {
		return stored.Items.Add(item) != nil
	})
	return nil
}

func (d *AccessItems) GetAll() (res []Item) {
	d.accessor(func(stored *Data) bool {
		res = stored.Items.Copy()
		return false
	})
	return
}
