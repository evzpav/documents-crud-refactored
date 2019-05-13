package domain

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Document struct {
	ID            bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`
	DocType       string    `json:"docType" bson:"doc_type"`
	IsBlacklisted bool      `json:"isBlacklisted" bson:"is_blacklisted"`
	Value         string    `json:"value" bson:"value"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
}

type ServerStatus struct {
	UpTime         float64 `json:"uptime"`
	SessionQueries int     `json:"sessionQueries"`
}

type DocumentService interface {
	Get(documentID string) (*Document, error)

	GetAll() (*[]Document, error)

	Insert(doc *Document) (string, error)

	Update(documentId string, doc *Document) error

	Remove(documentId string) error

	SessionQueries() int

}

type DocumentStorage interface {
	FindOne(documentID string) (*Document, error)

	FindAll() (*[]Document, error)

	Insert(doc *Document) (string, error)

	Update(documentId string, doc *Document) error

	RemoveOne(documentId string) error

	SessionQueries() int
}
