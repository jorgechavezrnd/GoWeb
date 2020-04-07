package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response ...
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	writer      http.ResponseWriter
}

// CreateDefaultResponse ...
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

// NotFound ...
func (response *Response) NotFound() {
	response.Status = http.StatusNotFound
	response.Message = "Resource Not Found"
}

// SendNotFound ...
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

// SendUnprocessableEntity ...
func SendUnprocessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnprocessableEntity()
	response.Send()
}

// UnprocessableEntity ...
func (response *Response) UnprocessableEntity() {
	response.Status = http.StatusUnprocessableEntity
	response.Message = "Unprocessable Entity"
}

// SendNoContent ...
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}

// NoContent ...
func (response *Response) NoContent() {
	response.Status = http.StatusNoContent
	response.Message = "No Content"
}

// SendData ...
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

// Send ...
func (response *Response) Send() {
	response.writer.Header().Set("Content-Type", response.contentType)
	response.writer.WriteHeader(response.Status)

	output, _ := json.Marshal(&response)
	fmt.Fprintf(response.writer, string(output))
}
