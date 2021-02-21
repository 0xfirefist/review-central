package users

import (
	"log"

	"github.com/kalradev/review-central/internal/db"
	"gorm.io/gorm"
)

type User struct {
	// Username   string `json:"username"`
	gorm.Model
	FirstName  string `json:"first-name"`
	MiddleName string `json:"middle-name"`
	LastName   string `json:"last-name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Number     string `json:"number"`
}

type IPFSHash struct {
	gorm.Model
	Hash   string
	UserID string
	User   User
}

var users = []*User{}

// Create - add user to the database
func (user *User) Create() error {
	database, err := db.GetDBInstance()
	if err != nil {
		log.Println("Error getting DB instance")
		return err
	}

	res := database.Create(user)
	log.Println("New User ID %s", user.ID)
	return res.Error
}

//GetUserByUsername check if a user exists in database by given username
func GetUserByEmail(email string) (*User, error) {
	database, err := db.GetDBInstance()
	if err != nil {
		log.Println("Error getting DB instance")
		return nil, err
	}
	var user User
	res := database.Where("email = ?", email).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// Authenticate user while login
func (u *User) Authenticate() bool {
	database, err := db.GetDBInstance()
	if err != nil {
		log.Println("Error getting DB instance")
		return false
	}
	var user User
	res := database.Where("email = ? AND password = ?", u.Email, u.Password).First(&user)

	if res.Error != nil {
		return false
	}
	return true
}

// AddReviewHash to map hash to users
func (u *User) AddReviewHash(hash string) error {
	database, err := db.GetDBInstance()
	if err != nil {
		log.Println("Error getting DB instance")
		return err
	}

	res := database.Create(&IPFSHash{
		Hash: hash,
		User: *u,
	})

	return res.Error
}
