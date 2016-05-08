package data

type AccessUsers struct {
	accessor LockerAccess
}

func (a *AccessUsers) GetAll() (users []User) {
	a.accessor(func(stored *Data) bool {
		users = stored.Users.Copy()
		return users == nil
	})
	return
}

func (a *AccessUsers) Add(u User) error {
	a.accessor(func(stored *Data) bool {
		err := stored.Users.Add(u)
		return err == nil
	})
	return nil
}

func (a *AccessUsers) Update(updated User) (user User, err error) {
	a.accessor(func(stored *Data) bool {
		user, err = stored.Users.Update(updated)
		return err == nil
	})
	return user, err
}

func (a *AccessUsers) Find(id string) (user User, err error) {
	a.accessor(func(store *Data) bool {
		user, err = store.Users.FindUser(id)
		return err == nil
	})
	return user, err
}

func (a *AccessUsers) FindByAuth(username, password string) (user User, err error) {
	a.accessor(func(store *Data) bool {
		user, err = store.Users.FindByAuth(username, password)
		return err == nil
	})
	return user, err
}
