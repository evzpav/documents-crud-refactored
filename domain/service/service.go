package service

import "github.com/evzpav/documents-crud-refactored/domain"

type documentService struct {
	documentStorage domain.DocumentStorage
}

func NewService(documentStorage domain.DocumentStorage) *documentService {
	return &documentService{
		documentStorage: documentStorage,
	}
}

func (ds *documentService) Get(documentID string) (*domain.Document, error) {
	doc, err := ds.documentStorage.FindOne(documentID)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (ds *documentService) GetAll() (docs *[]domain.Document, err error) {
	docs, err = ds.documentStorage.FindAll()
	if err != nil {
		return docs, err
	}
	return docs, nil
}

func (ds *documentService) Insert(doc *domain.Document) (string, error) {
	documentID, err := ds.documentStorage.Insert(doc)

	if err != nil {
		return "", err
	}

	return documentID, nil
}

func (ds *documentService) Update(documentID string, doc *domain.Document) error {
	err := ds.documentStorage.Update(documentID, doc)

	if err != nil {
		return err
	}

	return nil
}

func (ds *documentService) Remove(documentId string) error {
	err := ds.documentStorage.RemoveOne(documentId)

	if err != nil {
		return err
	}

	return nil
}

func (ds *documentService) SessionQueries() int {
	return ds.documentStorage.SessionQueries()

}
