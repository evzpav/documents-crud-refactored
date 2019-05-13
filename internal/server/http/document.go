package http

import (
	"fmt"
	"github.com/evzpav/documents-crud-refactored/domain"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

var queriesCounter int

func (h *handler) getIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "Document index")
}

func (h *handler) getDocuments(c echo.Context) error {
	docs, err := h.documentService.GetAll()
	if err != nil {
		log.Printf("Could get docs: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, docs)
}

func (h *handler) getDocument(c echo.Context) error {
	id, err := resolveID(c)
	if err != nil {
		log.Print(err)
		return err
	}

	docs, err := h.documentService.Get(id)
	if err != nil {
		log.Printf("Could get docs: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, docs)
}


func (h *handler) createDocument(c echo.Context) (err error) {
	var doc domain.Document
	if err = c.Bind(&doc); err != nil {
		log.Printf("Could not unmarshal doc: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()
	doc.ID = bson.NewObjectId()

	docs, err := h.documentService.Insert(&doc)
	if err != nil {
		log.Printf("Could get docs: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, docs)
}

func (h *handler) updateDocument(c echo.Context) (err error) {
	id, err := resolveID(c)
	if err != nil {
		log.Print(err)
		return err
	}

	var doc domain.Document
	if err = c.Bind(&doc); err != nil {
		log.Printf("Could not unmarshal doc: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	doc.UpdatedAt = time.Now()
	doc.ID = bson.ObjectIdHex(id)

	err = h.documentService.Update(doc.ID.Hex(), &doc)
	if err != nil {
		log.Printf("Could not update doc: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, doc)

}

func (h *handler) deleteDocument(c echo.Context) (err error) {
	id, err := resolveID(c)
	if err != nil {
		log.Print(err)
		return err
	}

	err = h.documentService.Remove(id)
	if err != nil {
		log.Printf("Could not update doc: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, id)

}

func resolveID(c echo.Context) (string, error) {
	id := c.Param("id")

	if id == "" {
		err := fmt.Errorf("invalid id")
		log.Print(err)
		return "", c.JSON(http.StatusBadRequest, err.Error())
	}
	return id, nil
}

func (h *handler) serverStatus(c echo.Context) (err error) {
	var status domain.ServerStatus
	uptime := time.Since(h.serverUptime)
	status.UpTime = uptime.Seconds()
	status.SessionQueries = h.documentService.SessionQueries()
	return c.JSON(http.StatusOK, status)


}
