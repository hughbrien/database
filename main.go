package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"sync"
)

type User struct {
	UUID        string
	FirstName   string
	LastName    string
	DateOfBirth int64
	Email       string
}

var userMap map[string]User
var mu sync.RWMutex

func main() {
	fmt.Println("Starting data of user program")
	initialize_database()
	newUUID, err := uuid.NewRandom()
	if err != nil {
		os.Exit(1)
	}

	hugh := User{
		UUID:        newUUID.String(),
		FirstName:   "Hugh",
		LastName:    "Brien",
		Email:       "hugh@hughbrien.com",
		DateOfBirth: 100000000,
	}
	hugh.Greet()
	userMap[hugh.UUID] = hugh
}

func addUserSafe(luser User) {
	mu.Lock() // lock for writing
	var newname = luser.LastName + " " + luser.LastName
	userMap[newname] = luser
	mu.Unlock()
}

func getUserAgeSafe(keyName string) User {
	mu.RLock() // lock for reading
	age := userMap[keyName]
	mu.RUnlock()
	return age
}

func initialize_database() {
	fmt.Println("Initializing Database")
	userMap = make(map[string]User)

}

func (p *User) UpdateEmail(newEmail string) {
	p.Email = newEmail
}

func (u User) Greet() {
	fmt.Printf("Hello, my name is %s %s.\n", u.FirstName, u.LastName)
}

func get_guid() string {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("Failed to generate UUID: %v", err)
		return "Error"
	}
	fmt.Printf("Generated UUID: %s\n", newUUID.String())
	return concatenateStrings("Generated UUID", " ", newUUID.String())
}

func concatenateStrings(parts ...string) string {
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part)
	}
	return sb.String()
}
