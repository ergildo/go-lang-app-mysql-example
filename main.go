package main

import (
	"fmt"
	"github.com/ergildo/go-lang-app-mysql-example/service"
	"github.com/ergildo/go-lang-app-mysql-example/setup"
	"log"
	"math/rand"
	"time"
)

func nextRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	log.Println("Starting the application...")
	defer log.Println("Finishing the application...")
	setupDB()
	InsertAll()
	listAll()
	updateAll()
	listAll()
	deleteAll()

}

func setupDB() {
	log.Println("Starting database setup...")
	defer log.Println("Finishing setup")
	setup.SetUpDB()
}



func deleteAll() {
	log.Println("Deleting all users...")
	users := service.ListAll()
	for _, user := range users {
		service.Delete(user.Id)
	}
}

func updateAll() {
	log.Println("Updating all users...")
	users := service.ListAll()
	for i, user := range users {
		newName := fmt.Sprintf("User_update_%d", i)
		newUser := service.User{
			Id:   user.Id,
			Name: newName,
		}
		service.Update(newUser)
	}
}

func listAll() {
	log.Println("Listing all users...")
	users := service.ListAll()

	for _, user := range users {
		fmt.Printf("{id: %d, name: %s}\n", user.Id, user.Name)
	}

}

func InsertAll() {
	log.Println("Inserting all users...")
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("User_%d", nextRandom())
		service.Save(service.User{Name: name})
	}
}
