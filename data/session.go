package data

import (
	"fmt"
	"errors"
)

type Session map[string]User

func (s Session) Find(key string) (User, error) {
	user, ok := s[key]
	if ok {
		return user, nil
	}
	return User{}, errors.New("no session found")
}

func (s Session) Set(key string, user User) {
	if s == nil {
		s = make(Session)
	}
	if key == "" {
		fmt.Println("cannot save session user with empty key")
		return
	}
	s[key] = user
}