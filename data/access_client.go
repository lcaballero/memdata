package data

const DefaultDbName = "items.db.json"

type DataAccess func(*Data) bool
type LockerAccess func(DataAccess)

type AccessClient struct {
	Store     *Locker
	Users     *AccessUsers
	Items     *AccessItems
	Session   *AccessSession
	DataStore LockerAccess
}

func NewAccessClient(dbname string) (*AccessClient, error) {
	store, err := NewDataStore(
		dbname,
		func() interface{} { return NewData().Add(RootUser()) },
		func() interface{} { return NewData() },
	)
	if err != nil {
		return nil, err
	}
	accessor := func(fn DataAccess) {
		var usage Access = func(locked interface{}) bool {
			data, ok := locked.(*Data)
			if !ok {
				panic("Should have been working with a Data instance all along.")
			}
			return fn(data)
		}
		store.DataStore(usage)
	}
	a := &AccessClient{
		Store:   store,
		Users:   &AccessUsers{accessor: accessor},
		Items:   &AccessItems{accessor: accessor},
		Session: &AccessSession{accessor: accessor},
	}

	return a, nil
}
