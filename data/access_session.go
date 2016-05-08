package data

import (
	"fmt"
	"github.com/pborman/uuid"
)

type AccessSession struct {
	accessor LockerAccess
}

func (a *AccessSession) Find(key string) (user User, err error) {
	a.accessor(func(store *Data) bool {
		user, err = store.Session.Find(key)
		return err != nil
	})
	return
}

func (a *AccessSession) Login(username, password string) (user User, sid string, err error) {
	a.accessor(func(store *Data) bool {
		sid := uuid.New()
		user, err = store.Users.FindByAuth(username, password)
		if err != nil {
			return false
		}
		fmt.Println("found user: ", username, user.IsValid())
		store.Session.Set(sid, user)
		return true
	})
	return
}

func (a *AccessSession) Logout(key string) (found bool) {
	a.accessor(func(store *Data) bool {
		_, found = store.Session[key]
		if found {
			delete(store.Session, key)
		}
		return false
	})
	return
}