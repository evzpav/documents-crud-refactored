package internal

import (
"fmt"
)

type notFoundError struct {
	message string
}

func (err *notFoundError) Error() string {
	return err.message
}

func (err *notFoundError) NotFound() {}

func NewNotFoundError(entity string) *notFoundError {
	return &notFoundError{fmt.Sprintf("error.%s.not.found", entity)}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	type notFound interface {
		NotFound()
	}

	_, ok := err.(notFound)
	return ok
}

type duplicatedRecord struct {
	message string
}

func (err *duplicatedRecord) Error() string {
	return err.message
}

func (err *duplicatedRecord) DuplicatedRecord() {}

func NewDuplicatedRecordError(entity string) *duplicatedRecord {
	return &duplicatedRecord{fmt.Sprintf("error.%s.already.exists", entity)}
}

func IsDuplicatedRecord(err error) bool {
	if err == nil {
		return false
	}

	type duplicatedRecord interface {
		DuplicatedRecord()
	}

	_, ok := err.(duplicatedRecord)
	return ok
}

type unauthorized struct {
	message string
}

func (err *unauthorized) Error() string {
	return err.message
}

func (err *unauthorized) Unauthorized() {}

func NewUnauthorizedError(message string) *unauthorized {
	return &unauthorized{message}
}

type excedeedLimit struct {
	message string
}

func (err *excedeedLimit) Error() string {
	return err.message
}

func (err *excedeedLimit) ExcedeedLimit() {}

func NewExcedeedLimitError(message string) *excedeedLimit {
	return &excedeedLimit{message}
}
