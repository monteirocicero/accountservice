package dbclient

import (
	"accountservice/model"

	"github.com/stretchr/testify/mock"
)

type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountId string) (model.Account, error) {
	args := m.Mock.Called(accountId)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// nothing
}

func (m *MockBoltClient) Seed() {

}

func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}
