package data

import (
	"errors"
)

var BadCredentialsError = errors.New("Bad Credentials Error")

type Users []User

func NewUsers(users ...User) Users {
	return users
}

func (users Users) FindByAuth(username, password string) (User, error) {
	for _, u := range users {
		if u.Password == password && u.Username == username {
			return u, nil
		}
	}
	return User{}, BadCredentialsError
}

func (users Users) FindUser(id string) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, errors.New("Did not find User")
}

func (users *Users) Add(u User) error {
	*users = append(*users, u)
	return nil
}

func (users Users) Len() int {
	return len(users)
}

func (users Users) Update(updated User) (User, error) {
	for i := 0; i < len(users); i++ {
		t := &users[i]
		if t.Id == updated.Id {
			t.Username = updated.Username
			t.Email = updated.Email
			t.FirstName = updated.FirstName
			t.LastName = updated.LastName
			return *t, nil
		}
	}
	return User{}, errors.New("Couldn't find item to update")
}

func (users Users) Copy() Users {
	other := make([]User, len(users))
	copy(other, users)
	return other
}
