package errors

import (
	"encoding/json"
	"fmt"
	"io"
)

type DefaultError struct {
	TimeStamp string    `json:"timestamp"`
	Status    int       `json:"status"`
	Err       string    `json:"error"`
	Details   []Details `json:"details"`
	Path      string    `json:"path"`
}

type Details struct {
	Field     string `json:"field"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

func (r *DefaultError) Error() string {

	queryString := "Error: %v\nDateTime: %v\nStatus: %d\n\nDetails: [\n"

	for _, detail := range r.Details {
		queryString += fmt.Sprintf("\t{\n\t\tField: %v\n\t\tMessage: %v\n\t\tErrorCode: %v\n\t}\n", detail.Field, detail.Message, detail.ErrorCode)
	}

	queryString += "]\n\nPath: %d\n\n"

	return fmt.Sprintf(queryString, r.Err, r.TimeStamp, r.Status, r.Path)
}

func ParseDefaultError(body io.ReadCloser) error {
	defaultError := DefaultError{}
	err := json.NewDecoder(body).Decode(&defaultError)

	if err != nil {
		return err
	}

	return &defaultError
}
