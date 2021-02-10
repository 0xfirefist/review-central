package users

import "errors"

type User struct {
	// Username   string `json:"username"`
	FirstName  string `json:"first-name"`
	MiddleName string `json:"middle-name"`
	LastName   string `json:"last-name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Number     string `json:"number"`
}

var users = []*User{
	&User{
		Email:    "asdf@asdf.com",
		Password: "asdf@asdf.com",
	},
}

// Create - add user to the database
func (user *User) Create() {
	users = append(users, user)
}

//GetUserByUsername check if a user exists in database by given username
func GetUserByEmail(email string) (*User, error) {

	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("User not found")
}

// Authenticate user while login
func (u *User) Authenticate() bool {
	for _, user := range users {
		if u.Email == user.Email && u.Password == user.Password {
			return true
		}
	}
	return false
}
