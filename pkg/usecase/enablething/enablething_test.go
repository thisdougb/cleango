//go:build dev
// +build dev

package enablething

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// using a test table
var TestItems = []struct {
	comment       string // a comment used to identify test in output
	thingID       int    // in our mock we use this value to affect the return values
	expectedError error
}{
	{
		comment:       "update existing thing",
		thingID:       1,
		expectedError: nil,
	},
	{
		comment:       "thing does not exist",
		thingID:       2,
		expectedError: errors.New("thing not found"),
	},
	{
		comment:       "datastore error",
		thingID:       3,
		expectedError: errors.New("datastore error"),
	},
}

func TestEnableThing(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	for _, item := range TestItems {

		err := s.EnableThing(item.thingID)
		assert.Equal(t, item.expectedError, err, item.comment)
	}
}
