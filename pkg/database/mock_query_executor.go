package db

import "github.com/stretchr/testify/mock"

type MockQueryExecutor struct {
	mock.Mock
}

func (m *MockQueryExecutor) Execute(query string) (string, error) {
	args := m.Called(query)
	return args.String(0), args.Error(1)
}
