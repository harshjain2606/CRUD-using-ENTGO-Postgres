package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"ent_postgres_crud/ent"

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

	// CREATE
	user, err := client.User.Create().
		SetName("Harsh").
		SetEmail("hhh@example.com").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	fmt.Println("User created:", user)

	// READ
	readUser, err := client.User.Get(ctx, user.ID)
	if err != nil {
		log.Fatalf("failed reading user: %v", err)
	}
	fmt.Println("User read:", readUser)

	// UPDATE
	updatedUser, err := client.User.
		UpdateOneID(user.ID).
		SetName("Harsh Updated").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	fmt.Println("User updated:", updatedUser)

	//DELETE
	// err = client.User.DeleteOneID(user.ID).Exec(ctx)
	// if err != nil {
	// 	log.Fatalf("failed deleting user: %v", err)
	// }
	// fmt.Println("User deleted")
}
