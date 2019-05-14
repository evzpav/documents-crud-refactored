package storage

import (
	"github.com/evzpav/documents-crud-refactored/domain"
	"github.com/evzpav/documents-crud-refactored/internal/storage/mongo"
)

type logger interface {
	Printf(string, ...interface{})
}

type documentStorage struct {
	session        *mongo.Mongo
	databaseName   string
	collectionName string
	debugLog       logger
}

var queriesCounter int

func NewDocumentStorage(databaseURL, databaseName, collectionName string, debugLog logger) (*documentStorage, error) {
	mongoSession, err := mongo.NewMongo(databaseURL, databaseName, collectionName, debugLog)
	if err != nil {
		return nil, err
	}
	return &documentStorage{
		session:        mongoSession,
		databaseName:   databaseName,
		collectionName: collectionName,
		debugLog:       debugLog,
	}, nil

}

func (ds *documentStorage) Insert(document *domain.Document) (string, error) {
	queriesCounter++
	return ds.session.Insert(document)
}

func (ds *documentStorage) Update(documentID string, doc *domain.Document) error {
	queriesCounter++
	return ds.session.Update(documentID, doc)
}

func (ds *documentStorage) FindOne(documentID string) (*domain.Document, error) {
	queriesCounter++
	return ds.session.FindOne(documentID)
}

func (ds *documentStorage) FindAll() (*[]domain.Document, error) {
	queriesCounter++
	return ds.session.FindAll()
}

func (ds *documentStorage) RemoveOne(documentID string) error {
	queriesCounter++
	return ds.session.RemoveOne(documentID)
}

func (ds *documentStorage) SessionQueries() int {
	return queriesCounter
}
