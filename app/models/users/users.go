package users

import (
	"github.com/jinzhu/gorm"

	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User is the data model for users
type User struct {
	gorm.Model
	Username string
	Password string
}

var db *gorm.DB

// Init handles migration of data models
func Init(gormdb *gorm.DB) {
	db = gormdb
	// Migrate the schema
	db.AutoMigrate(&User{})
}

// CreateUser create a new user in the database
func CreateUser() (user User) {
	// Create
	user = User{Username: "L1212", Password: "1234"}
	db.Create(&user)
	return
}

// ListUsers lists all users in database
func ListUsers() (users []User) {
	db.Find(&users)
	return
}
