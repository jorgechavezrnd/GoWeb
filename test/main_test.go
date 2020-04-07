package test

import (
	"fmt"
	"os"
	"testing"

	"gitlab.com/jorgechavezrnd/go_rest/models"
)

func TestMain(m *testing.M) {
	beforeTest()
	result := m.Run()
	afterTest()
	os.Exit(result)
}

func beforeTest() {
	fmt.Println(">> Antes de las pruebas")
	models.CreateConnection()
	models.CreateTables()
	createDefaultUser()
}

func createDefaultUser() {
	sql := fmt.Sprintf("INSERT users SET id='%d', username='%s', password='%s', email='%s', created_date='%s'", id, username, passswordHash, email, createdDate)
	_, err := models.Exec(sql)
	if err != nil {
		panic(err)
	}
	user = &models.User{ID: id, Username: username, Password: password, Email: email}
}

func afterTest() {
	fmt.Println(">> Despues de las pruebas")
	models.CloseConnection()
}
