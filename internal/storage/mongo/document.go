package mongo

import (
	"errors"
	"fmt"
	"github.com/evzpav/documents-crud-refactored/domain"
	"github.com/evzpav/documents-crud-refactored/internal"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"strings"
)

var queriesCounter int

type documentStorage struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
	debugLog       logger
}

func NewDocumentStorage(mongoURL, databaseName string, debugLog logger) (*documentStorage, error) {
	log.Println(mongoURL) //TODO remove
	if strings.TrimSpace(mongoURL) == "" {
		return nil, errors.New("mongoURL is required")
	}
	mongoSession, err := mgo.Dial(mongoURL)
	if err != nil {
		return nil, fmt.Errorf("error on MongoDB connection: %q", err)
	}

	//mongoSession.SetMode(mgo.Monotonic, true)

	//index := mgo.Index{
	//	Key:        []string{"document", "documentId"},
	//	Unique:     true,
	//	Background: true,
	//}

	collectionName := "documents"
	cc := mongoSession.DB(databaseName).C(collectionName)

	log.Printf("collection: %+v ", cc)
	//if err = mongoSession.DB(databaseName).C(collectionName); err != nil {
	//	return nil, fmt.Errorf("error on MongoDB ensure document unique index: %q", err)
	//}

	return &documentStorage{
		session:        mongoSession,
		databaseName:   databaseName,
		collectionName: collectionName,
		debugLog:       debugLog,
	}, nil

}

func (ds *documentStorage) Insert(document *domain.Document) (string, error) {
	queriesCounter++
	ds.debugLog.Printf("storage: insert document [%s] and value [%s]", document.ID, document.Value)

	session := ds.session.Copy()
	defer session.Close()

	document.ID = bson.NewObjectId()

	if err := session.DB(ds.databaseName).C(ds.collectionName).Insert(document); err != nil {
		//if mgo.IsDup(err) {
		//	return "", internal.NewDuplicatedRecordError("document")
		//}
		return "", fmt.Errorf("mongo: cannot insert document: %q", err)
	}

	ds.debugLog.Printf("storage: successfully insert document for doc ID [%s] and value [%s]", document.ID, document.Value)

	return document.ID.Hex(), nil
}

func (ds *documentStorage) Update(documentId string, doc *domain.Document) error {
	queriesCounter++
	ds.debugLog.Printf("storage: update document for doc ID [%s] and ID [%s]", documentId, doc.Value)

	session := ds.session.Copy()
	defer session.Close()

	key := bson.ObjectIdHex(documentId)

	if err := session.DB(ds.databaseName).C(ds.collectionName).UpdateId(key, doc); err != nil {
		if err == mgo.ErrNotFound {
			return internal.NewNotFoundError("document")
		}
		return fmt.Errorf("mongo: cannot update document: %q", err)
	}

	ds.debugLog.Printf("storage: successfully update document for doc ID [%s] and value [%s]", documentId, doc.Value)

	return nil
}

func (ds *documentStorage) FindOne(documentID string) (*domain.Document, error) {
	queriesCounter++
	ds.debugLog.Printf("storage: getting document for doc ID [%s]", documentID)

	session := ds.session.Copy()
	defer session.Close()

	key := bson.ObjectIdHex(documentID)
	var document *domain.Document
	if err := session.DB(ds.databaseName).C(ds.collectionName).FindId(key).One(&document); err != nil {
		if err == mgo.ErrNotFound {
			return nil, internal.NewNotFoundError("document")
		}
		return nil, fmt.Errorf("mongo: cannot get the document: %q", err)
	}

	ds.debugLog.Printf("storage: document retrieved with success doc ID [%s] and value [%s]", documentID, document.Value)

	return document, nil
}

func (ds *documentStorage) FindAll() (*[]domain.Document, error) {
	queriesCounter++
	ds.debugLog.Printf("storage: getting all documents")

	session := ds.session.Copy()
	defer session.Close()

	key := bson.M{}

	var docs = make([]domain.Document, 0)
	if err := session.DB(ds.databaseName).C(ds.collectionName).Find(key).All(&docs); err != nil {

		if err == mgo.ErrNotFound {
			return nil, internal.NewNotFoundError("document")
		}
		return nil, fmt.Errorf("mongo: cannot get the document: %q", err)
	}

	ds.debugLog.Printf("storage: documents retrieved with success")

	return &docs, nil
}

func (ds *documentStorage) RemoveOne(documentID string) error {
	queriesCounter++
	ds.debugLog.Printf("storage: removing document for doc ID [%s]", documentID)

	session := ds.session.Copy()
	defer session.Close()

	key := bson.ObjectIdHex(documentID)

	if err := session.DB(ds.databaseName).C(ds.collectionName).RemoveId(key); err != nil {
		if err == mgo.ErrNotFound {
			return internal.NewNotFoundError("document")
		}
		return fmt.Errorf("mongo: cannot remove the document: %q", err)
	}

	ds.debugLog.Printf("storage: document removed with success doc ID [%s]", documentID)

	return nil
}

func (ds *documentStorage) SessionQueries() int {
	return queriesCounter

}