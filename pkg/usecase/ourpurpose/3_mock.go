//go:build dev || test
// +build dev test

package ourpurpose

type MockRepository struct {
	MockReader
	MockWriter
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
