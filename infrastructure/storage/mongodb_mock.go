package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/mock"
)

//DataAccessLayerMock is a mock for db connection
type DataAccessLayerMock struct {
	mock.Mock
}

//WithTransaction is a mock for db WithTransaction
func (m *DataAccessLayerMock) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}

//Initialize is a mock for db Initialize
func (m *DataAccessLayerMock) Initialize(ctx context.Context, credential options.Credential, dbURI, dbName string) error {
	GetInstance()
	mongoInstance = m
	return nil
}

//Insert is a mock for db Insert
func (m *DataAccessLayerMock) Insert(ctx context.Context, collName string, doc interface{}) (interface{}, error) {
	args := m.Called(ctx, collName, doc)
	return args.Get(0), args.Error(1)
}

//FindOne is a mock for db FindOne
func (m *DataAccessLayerMock) FindOne(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) error {
	args := m.Called(ctx, collName, query, doc)
	return args.Error(0)
}

//Find is a mock for db Find
func (m *DataAccessLayerMock) Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) error {
	args := m.Called(ctx, collName, query, doc)
	return args.Error(0)
}

//Count is a mock for db Count
func (m *DataAccessLayerMock) Count(ctx context.Context, collName string, query map[string]interface{}) (int64, error) {
	args := m.Called(ctx, collName, query)
	return int64(args.Int(0)), args.Error(1)
}

//UpdateOne is a mock for UpdateOne
func (m *DataAccessLayerMock) UpdateOne(ctx context.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, collName, selector, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

//Remove is a mock for Remove
func (m *DataAccessLayerMock) Remove(ctx context.Context, collName string, selector map[string]interface{}) error {
	args := m.Called(ctx, collName, selector)
	return args.Error(0)
}

//Disconnect is a mock for Disconnect
func (m *DataAccessLayerMock) Disconnect() {}
