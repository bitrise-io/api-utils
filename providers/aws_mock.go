package providers

import "time"

// AWSMock ...
type AWSMock struct {
	Config                    AWSConfig
	GeneratePresignedGETURLFn func(string, time.Duration) (string, error)
	GeneratePresignedPUTURLFn func(string, time.Duration, int64) (string, error)
	GetObjectFn               func(string) (string, error)
	MoveObjectFn              func(string, string) error
}

// GetConfig ...
func (m *AWSMock) GetConfig() AWSConfig {
	return m.Config
}

// GeneratePresignedGETURL ...
func (m *AWSMock) GeneratePresignedGETURL(key string, expiresIn time.Duration) (string, error) {
	if m.GeneratePresignedGETURLFn == nil {
		panic("You have to override GeneratePresignedGETURL function in tests")
	}
	return m.GeneratePresignedGETURLFn(key, expiresIn)
}

// GeneratePresignedPUTURL ...
func (m *AWSMock) GeneratePresignedPUTURL(key string, expiresIn time.Duration, fileSize int64) (string, error) {
	if m.GeneratePresignedPUTURLFn == nil {
		panic("You have to override GeneratePresignedPUTURL function in tests")
	}
	return m.GeneratePresignedPUTURLFn(key, expiresIn, fileSize)
}

// GetObject ...
func (m *AWSMock) GetObject(key string) (string, error) {
	if m.GetObjectFn == nil {
		panic("You have to override GetObject function in tests")
	}
	return m.GetObjectFn(key)
}

// MoveObject ...
func (m *AWSMock) MoveObject(from string, to string) error {
	if m.MoveObjectFn == nil {
		panic("You have to override MoveObject function in tests")
	}
	return m.MoveObjectFn(from, to)
}