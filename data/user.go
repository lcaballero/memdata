package data

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

//go:generate genfront fields --output user.gen.go --template update_and_copy.gen.fm --line $GOLINE
type User struct {
	Id        string
	GroupId   string
	CreatedOn time.Time
	CreatedBy string
	UpdatedOn time.Time
	UpdatedBy string

	Username string
	Password string

	FirstName    string
	LastName     string
	Email        string
	Icon         string
	RecordStatus RecordStatus
	UserState    UserState
	UserRole     UserRole
}

func RootUser() User {
	u := NewUser()
	user := &u
	user.UserRole = RootRole
	user.Password = "manofsteel1"
	user.Username = "kel.el"
	return u
}

func NewUser() User {
	now := time.Now()
	return User{
		Id:           uuid.New(),
		CreatedOn:    now,
		UpdatedOn:    now,
		RecordStatus: Active,
		UserState:    Normal,
	}
}

func (u User) IsValid() bool {
	return u.Id != "" && !u.CreatedOn.IsZero() && !u.UpdatedOn.IsZero()
}
