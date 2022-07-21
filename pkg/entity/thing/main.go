package thing

// Just a simple type

type Thing struct {
	Identifier int
	IsEnabled  bool
	Name       string
}

func NewThing(id int) *Thing {
	return &Thing{Identifier: id}
}
