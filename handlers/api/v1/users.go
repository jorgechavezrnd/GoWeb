package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/jorgechavezrnd/go_rest/models"
)

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.SetPassword(user.Password)
	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, user)
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	request := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(request); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	user.Username = request.Username
	user.Email = request.Email
	user.SetPassword(request.Password)

	if err := user.Save(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	models.SendData(w, user)
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		user.Delete()
		models.SendNoContent(w)
	}
}

func getUserByRequest(r *http.Request) (*models.User, error) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	user := models.GetUserByID(userID)

	if user.ID == 0 {
		return user, errors.New("El usuario no existe en la base de datos")
	}
	return user, nil
}
