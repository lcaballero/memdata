package data

type Data struct {
	Users Users
	Items []Item
}

func NewData() *Data {
	return &Data{
		Users: make([]User, 0),
		Items: make([]Item, 0),
	}
}
func (d *Data) Add(u User) *Data {
	d.Users = append(d.Users, u)
	return d
}
