package models

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	createdDate time.Time
}

// var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-][email protected][a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var emailRegexp = regexp.MustCompile("^([a-zA-Z0-9_\\-\\.]+)@([a-zA-Z0-9_\\-\\.]+)\\.([a-zA-Z]{2,5})$")

const userSchema string = `
CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL UNIQUE,
	password VARCHAR(60) NOT NULL,
	email VARCHAR(40),
	created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

// Users ...
type Users []User

// NewUser ...
func NewUser(username, password, email string) (*User, error) {
	user := &User{Username: username, Email: email}
	if err := user.Valid(); err != nil {
		return &User{}, err
	}

	err := user.SetPassword(password)
	return user, err
}

// CreateUser ...
func CreateUser(username, password, email string) (*User, error) {
	user, err := NewUser(username, password, email)
	if err != nil {
		return &User{}, err
	}
	err = user.Save()
	return user, err
}

// GetUserByUsername ...
func GetUserByUsername(username string) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE username=?"
	return GetUser(sql, username)
}

// GetUserByID ...
func GetUserByID(id int) *User {
	sql := "SELECT id, username, password, email, created_date FROM users WHERE id=?"
	return GetUser(sql, id)
}

// GetUser ...
func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.createdDate)
	}

	return user
}

// GetUsers ...
func GetUsers() Users {
	sql := "SELECT id, username, password, email, created_date FROM users"
	users := Users{}
	rows, _ := Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.createdDate)
		users = append(users, user)
	}

	return users
}

// Login ...
func Login(username, password string) (*User, error) {
	user := GetUserByUsername(username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return &User{}, errorLogin
	}

	return user, nil
}

// ValidEmail ...
func ValidEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errorEmail
	}
	return nil
}

// ValidUsername ...
func ValidUsername(username string) error {
	if username == "" {
		return errorUsername
	}

	if len(username) > 30 {
		return errorLargeUsername
	}

	if len(username) < 5 {
		return errorShortUsername
	}

	return nil
}

// Valid ...
func (user *User) Valid() error {
	if err := ValidEmail(user.Email); err != nil {
		return err
	}
	if err := ValidUsername(user.Username); err != nil {
		return err
	}
	return nil
}

// Save ...
func (user *User) Save() error {
	if user.ID == 0 {
		return user.insert()
	} else {
		return user.update()
	}
}

func (user *User) insert() error {
	sql := "INSERT users SET username=?, password=?, email=?"
	id, err := InsertData(sql, user.Username, user.Password, user.Email)
	user.ID = id
	return err
}

func (user *User) update() error {
	sql := "UPDATE users SET username=?, password=?, email=?"
	_, err := Exec(sql, user.Username, user.Password, user.Email)
	return err
}

// Delete ...
func (user *User) Delete() error {
	sql := "DELETE FROM users WHERE id=?"
	_, err := Exec(sql, user.ID)
	return err
}

// SetPassword ...
func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorPasswordEncryption
	}
	user.Password = string(hash)
	return nil
}

// GetCreatedDate ...
func (user *User) GetCreatedDate() time.Time {
	return user.createdDate
}
