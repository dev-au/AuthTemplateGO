package resources

import (
	"AuthTemplate/src/models"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func createSuperUser() {
	var email string
	var password string
	fmt.Println("Creating SuperUser")
	fmt.Print("Email: ")
	_, err := fmt.Scan(&email)
	if err != nil {
		log.Fatal("Error reading email:", err)
	}

	fmt.Print("Password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Fatal("Error reading password:", err)
	}

	superAdmin := models.User{
		Email:    strings.TrimSpace(email),
		Name:     "SuperAdmin",
		IsActive: true,
		Password: strings.TrimSpace(password),
		RoleID:   nil,
	}

	result := DB.Create(&superAdmin)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value") {
			return
		}
		log.Fatalf("Unexpected error while creating SuperAdmin: %v", err)
	} else {
		fmt.Println("SuperAdmin created successfully with ID:", superAdmin.ID)
	}
}

func makeRand32ByteKeyword() {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	fmt.Println("Your 32 byte key: ", base64.StdEncoding.EncodeToString(key))
}

func InitCommands(command string) {
	if command == "createsuperuser" {
		createSuperUser()
	} else if command == "randomtoken" {
		makeRand32ByteKeyword()
	}
}
