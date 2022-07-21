package datastore

import (
	"github.com/thisdougb/cleango/pkg/usecase/enablething"
)

// DatastoreInterface methods are implemented by any concrete datastore
type DatastoreInterface interface {
	Connect() bool
	Disconnect()

	enablething.Repository
}
