package http

import (
	"github.com/evzpav/documents-crud-refactored/domain"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type logger interface {
	Printf(string, ...interface{})
}

type handler struct {
	documentService domain.DocumentService
	debugLog        logger
	errorLog        logger
}

func NewHandler(documentService domain.DocumentService, debugLog logger, errorLog logger) http.Handler {
	handler := &handler{
		documentService: documentService,
		debugLog:        debugLog,
		errorLog:        errorLog,
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))


	e.GET("/", handler.getIndex)
	e.GET("/documents", handler.getDocuments)
	e.GET("/document/:id", handler.getDocument)
	e.POST("/document", handler.createDocument)
	e.PUT("/document/:id", handler.updateDocument)
	e.DELETE("/document/:id", handler.deleteDocument)
	//e.GET("/status", handler.serverStatus)

	return e
}
