package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Initialize flags
	var (
		to      = flag.String("to", "", "Recipient email address")
		subject = flag.String("subject", "", "Email subject")
		body    = flag.String("body", "", "Email body")
	)

	flag.Parse()

	if *to == "" {
		fmt.Println("Error: recipient email address is required")
		os.Exit(1)
	}

	// TODO: Implement email sending logic
	fmt.Println("Email CLI initialized")
}
