package dbclient

import (
        "github.com/stretchr/testify/mock"
        "github.com/hidoudou/go-recipes/goblog/accountservice/model"
)

// MockBoltClient is a mock implementation of a datastore client for testing purposes
type MockBoltClient struct {
        mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
        args := m.Mock.Called(accountId)
        return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
        // Does nothing
}

func (m *MockBoltClient) Seed() {
        // Does nothing
}