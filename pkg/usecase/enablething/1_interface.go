package enablething

import ()

//Repository interface pattern
type Repository interface {
	Reader
	Writer
}

type Reader interface{}

type Writer interface {
	SetThingStatus(thingID int, status bool) error
}
