package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ent_postgres_crud/ent"
	"ent_postgres_crud/ent/user"

	_ "github.com/lib/pq"
)

func main() {
	constStr := os.Getenv("Database_URL")
	if constStr == "" {
		log.Fatal("Database_URL environment variable is not set")
	}

	client, err := ent.Open("postgres", constStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Auto migrate
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	email := "haarshjaainh@example.com"

	// Check if user already exists
	existingUser, err := client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)

	if err == nil {
		fmt.Println("User already exists:", existingUser)
	} else {
		// CREATE
		newUser, err := client.User.Create().
			SetName("Harsh").
			SetEmail(email).
			Save(ctx)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
		fmt.Println("User created:", newUser)
		existingUser = newUser
	}

	// READ
	readUser, err := client.User.Get(ctx, existingUser.ID)
	if err != nil {
		log.Fatalf("failed reading user: %v", err)
	}
	fmt.Println("User read:", readUser)

	// UPDATE
	updatedUser, err := client.User.
		UpdateOneID(existingUser.ID).
		SetName("Harsh Updated").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	fmt.Println("User updated:", updatedUser)

	// DELETE (optional)
	// err = client.User.DeleteOneID(existingUser.ID).Exec(ctx)
	// if err != nil {
	// 	log.Fatalf("failed deleting user: %v", err)
	// }
	// fmt.Println("User deleted")
}
