package gqlrelay

import "fmt"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = map[string]*User{
	"1": {
		ID:   "1",
		Name: "Kevin",
	},
	"2": {
		ID:   "2",
		Name: "Marla",
	},
	"3": {
		ID:   "3",
		Name: "Jose",
	},
}

var nextUser = 4

func CreateUser(userName string) *User {
	nextUser = nextUser + 1
	newUser := &User{
		fmt.Sprintf("%v", nextUser),
		userName,
	}
	users[newUser.ID] = newUser

	return newUser
}

func GetUser(id string) *User {
	if user, ok := users[id]; ok {
		return user
	}
	return nil
}

func GetUsers() map[string]*User {
	return users
}
