package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gitlab.com/jorgechavezrnd/go_rest/models"
)

var user *models.User

const (
	id            = 1
	username      = "eduardo_gpg"
	password      = "password"
	passswordHash = "$2a$10$OzpkzxUgk7LTgbpqDXyVYeHyQlDGtXebfo9.oUR7pG3.n2BltBWca"
	email         = "eduardo@gmail.com"
	createdDate   = "2017-08-17"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("No es posible crear el objeto", err)
	}
}

func TestPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60 {
		t.Error("No es posible cifrar el password")
	}
}

func TestValidEmail(t *testing.T) {
	if err := models.ValidEmail(email); err != nil {
		t.Error("Validacion erronea en el email", err)
	}
}

func TestUsernameLength(t *testing.T) {
	newUsername := username
	for i := 0; i < 10; i++ {
		newUsername += newUsername
	}

	_, err := models.NewUser(newUsername, password, email)
	if err == nil || err.Error() != "Username muy largo, maximo 30 caracteres" {
		t.Error("Es posible generar un usuario con un username muy grande")
	}
}

func TestInvalidEmail(t *testing.T) {
	if err := models.ValidEmail("adasdhsjkhdak.com"); err == nil {
		t.Error("Validacion erronea en el email")
	}
}

func TestLogin(t *testing.T) {
	if valid := models.Login(username, password); !valid {
		t.Error("No es posible realizar el login")
	}
}

func TestNoLogin(t *testing.T) {
	if valid := models.Login(randomUsername(), password); valid {
		t.Error("Es posible realizar el login con parametros erroneos")
	}
}

func TestSave(t *testing.T) {
	user, _ := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto", err)
	}
}

func TestUniqueUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err == nil {
		t.Error("Es posible insertar registros con usernames duplicados!")
	}
}

func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'username'", username)
	if err.Error() != message {
		t.Error("Es posible tener un username duplicado en la base de datos!")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUserByID(id)
	if !equalsUser(user) || !equalsCreatedDate(user.GetCreatedDate()) {
		t.Error("No es posible obtener el usuario")
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener a los usuarios")
	}
}

// func TestDeleteUser(t *testing.T) {
// 	if err := user.Delete(); err != nil {
// 		t.Error("No es posible eliminar al usuario")
// 	}
// }

func equalsCreatedDate(date time.Time) bool {
	t, _ := time.Parse("2006-01-02", createdDate)
	return t == date
}

func equalsUser(user *models.User) bool {
	return user.Username == username && user.Email == email
}

func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}
