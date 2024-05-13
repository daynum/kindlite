package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

// main just for testing
func sendEmail() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env file: %v", err)
		return
	}
	sender := os.Getenv("SENDER_EMAIL")
	pword := os.Getenv("PWORD")
	recipient := os.Getenv("RECIPIENT")
	auth := smtp.PlainAuth("", sender, pword, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:25", auth, sender, []string{recipient}, []byte("Hi! from Dinank on smtp"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Sending Email to: %s", recipient)

}
