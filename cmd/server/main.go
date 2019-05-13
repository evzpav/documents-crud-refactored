package main

import (
	"fmt"
	"github.com/evzpav/documents-crud-refactored/domain/document"
	"github.com/evzpav/documents-crud-refactored/internal/server/http"
	"github.com/evzpav/documents-crud-refactored/internal/storage/mongo"
	"log"
	"os"
	"strconv"
)

const (
	ENVVAR_DEBUG               = "DEBUG"
	DEFAULT_DATABASE_NAME      = "crmin"
	DEFAULT_DEBUG              = true
	ENVVAR_SERVICE_PORT        = "SERVICE_PORT"
	ENVVAR_DATABASE_NAME       = "DATABASE_NAME"
	ENVVAR_MONGO_URL           = "MONGO_URL"
	ENVVAR_CONFIG_SERVICE_HOST = "CONFIG_SERVICE_HOST"
	ENVVAR_CONFIG_SERVICE_PORT = "CONFIG_SERVICE_PORT"
	ENVVAR_PLATFORM_URL        = "PLATFORM_URL"
)

func main() {

	errorLog, debugLog, err := createLoggers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create loggers: %s\n", err.Error())
		return
	}

	//platformClient := platform.NewPlatformClient(getPlatformURL(), getConfigURL())

	//mongoURL := getMongoURL()
	mongoURL := fmt.Sprintf("mongodb://%s:%s", "localhost","27017")

	//dbName := getDatabaseName()
	dbName := "documents-crud"
	documentStorage, err := mongo.NewDocumentStorage(mongoURL, dbName, debugLog)
	if err != nil {
		errorLog.Printf("error on creating a document storage instance: %q", err)
		return
	}

	//servicePort := getServicePort()
	servicePort := "1323"
	documentService := document.NewService(documentStorage)
	server := http.New(servicePort, documentService, debugLog, errorLog)
	server.ListenAndServe()

}

func getMongoURL() string {
	return getEnvVar(ENVVAR_MONGO_URL)
}

func getDatabaseName() string {
	return getEnvVar(ENVVAR_DATABASE_NAME, DEFAULT_DATABASE_NAME)
}

func getPlatformURL() string {
	return getEnvVar(ENVVAR_PLATFORM_URL)
}

func getConfigURL() string {
	configHost := os.Getenv(ENVVAR_CONFIG_SERVICE_HOST)
	configPort := os.Getenv(ENVVAR_CONFIG_SERVICE_PORT)

	return "http://" + configHost + ":" + configPort + "/config"
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


