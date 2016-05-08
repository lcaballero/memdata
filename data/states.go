package data

//go:generate stringer -type=ItemState
type ItemState int

const (
	Inception  ItemState = 1
	InArchived ItemState = 2
	InProgress ItemState = 3
	InBacklog  ItemState = 4
	Completed  ItemState = 5
)

//go:generate stringer -type=RecordStatus
type RecordStatus int

const (
	Active  RecordStatus = 1
	Deleted RecordStatus = 2
)

//go:generate stringer -type=UserState
type UserState int

const (
	Normal   UserState = 1
	Disabled UserState = 2
	Locked   UserState = 3
)

//go:generate stringer -type=UserRole
type UserRole int

const (
	RootRole   UserRole = 1
	NormalRole UserRole = 2
)
