package orm

import "time"

// User ...
type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

// Users ...
type Users []User

// CreateUser ...
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// NewUser ...
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// GetUsers ...
func GetUsers() Users {
	users := Users{}
	db.Find(&users)
	return users
}

// GetUser ...
func GetUser(id int) *User {
	user := &User{}
	db.Where("id=?", id).First(user)
	return user
}

// Save ...
func (user *User) Save() {
	if user.ID == 0 {
		db.Create(&user)
	} else {
		user.update()
	}
}

func (user *User) update() {
	userUpdated := User{Username: user.Username, Password: user.Password, Email: user.Email}
	db.Model(&user).UpdateColumns(userUpdated)
}

// Delete ...
func (user *User) Delete() {
	db.Delete(&user)
}
