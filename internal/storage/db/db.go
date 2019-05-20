package db

import (
	"github.com/evzpav/documents-crud-refactored/internal/storage/mongo"
	"github.com/evzpav/documents-crud-refactored/internal/storage/postgres"
)

type Database struct {
	Postgres *postgres.Postgres
	Mongo    *mongo.Mongo
}
