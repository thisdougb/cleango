//go:build dev || test
// +build dev test

package ourpurpose

import (
	"errors"
)

type MockWriter struct{}

// Mock methods here, with conditionals enabling testing of return values
func (m *MockWriter) SetThingStatus(thingID int, status bool) error {

	if thingID == 2 {
		return errors.New("thing not found")
	}

	if thingID == 3 {
		return errors.New("datastore error")
	}

	return nil
}
