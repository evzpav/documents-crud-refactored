package main

import (
	"fmt"
	"github.com/evzpav/documents-crud-refactored/domain/document"
	"github.com/evzpav/documents-crud-refactored/internal/server/http"
	"github.com/evzpav/documents-crud-refactored/internal/storage/mongo"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	ENVVAR_DEBUG               = "DEBUG"
	DEFAULT_DATABASE_NAME      = "documents-crud"
	DEFAULT_DEBUG              = true
	ENVVAR_SERVICE_PORT        = "SERVICE_PORT"
	ENVVAR_DATABASE_NAME       = "DATABASE_NAME"
	ENVVAR_MONGO_HOST          = "MONGO_HOST"
	ENVVAR_MONGO_PORT          = "MONGO_PORT"
)
var serverUp time.Time

func main() {
	serverUp = time.Now()
	errorLog, debugLog, err := createLoggers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create loggers: %s\n", err.Error())
		return
	}

	mongoURL := fmt.Sprintf("mongodb://%s:%s", getMongoHost(), getMongoPort())

	dbName := getDatabaseName()
	documentStorage, err := mongo.NewDocumentStorage(mongoURL, dbName, debugLog)
	if err != nil {
		errorLog.Printf("error on creating a document storage instance: %q", err)
		return
	}

	servicePort := getServicePort()
	documentService := document.NewService(documentStorage)
	server := http.New(servicePort, documentService, debugLog, errorLog, serverUp)
	server.ListenAndServe()

}

func getMongoHost() string {
	return getEnvVar(ENVVAR_MONGO_HOST)
}

func getMongoPort() string {
	return getEnvVar(ENVVAR_MONGO_PORT)
}

func getDatabaseName() string {
	return getEnvVar(ENVVAR_DATABASE_NAME, DEFAULT_DATABASE_NAME)
}

func getServicePort() string {
	return getEnvVar(ENVVAR_SERVICE_PORT)
}

func getEnvVar(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}

type logger interface {
	Printf(string, ...interface{})
}

type discardLog struct{}

func (*discardLog) Printf(string, ...interface{}) {}

func createLoggers() (logger, logger, error) {
	var errorLog logger
	var debugLog logger

	errorLog = log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)

	if !debugEnabled() {
		return errorLog, &discardLog{}, nil
	}

	debugLog = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)

	return errorLog, debugLog, nil
}

func debugEnabled() bool {
	debug, err := strconv.ParseBool(os.Getenv(ENVVAR_DEBUG))
	if err != nil {
		debug = DEFAULT_DEBUG
	}

	return debug
}
