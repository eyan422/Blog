package main

import (
	"github.com/eyan422/Blog/CommonStruct"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateArticle(article *CommonStruct.Article) (int64, error) {
	rets := m.Called(article)
	return 1, rets.Error(0)
}

func (m *MockStore) GetArticle(id uint64) (article *CommonStruct.Article, err error) {
	rets := m.Called()
	return rets.Get(0).(*CommonStruct.Article), rets.Error(1)
}

func (m *MockStore) GetArticles() ([]*CommonStruct.Article, error) {
	rets := m.Called()
	return rets.Get(0).([]*CommonStruct.Article), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
