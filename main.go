package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	// Load service account key
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	// Initialize Firebase app
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	// Get Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}

	// Get all users
	iter := authClient.Users(context.Background(), "")
	var emails []string

	for {
		user, err := iter.Next()
		if err != nil {
			break
		}

		if user.Email != "" {
			emails = append(emails, user.Email)
			fmt.Println("Email:", user.Email)
		}
	}

	// Save emails to JSON file in required format
	output := map[string][]string{
		"emails": emails,
	}

	file, err := os.Create("firebase_emails.json")
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(output)
	if err != nil {
		log.Fatalf("Error writing to file: %v\n", err)
	}

	fmt.Println("✅ Emails saved to firebase_emails.json")
}
