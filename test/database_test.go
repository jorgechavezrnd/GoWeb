package test

import (
	"testing"

	"gitlab.com/jorgechavezrnd/go_rest/models"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("No es posible realizar la conexion", nil)
	}
}
