package main

import (
    "fmt"
    "example.com/check"
)

func main() {
    // List of emails to check
    emails := []string{"example@domain.com", "notanemail", "another@example.com"}

    // Loop through each email
    for _, email := range emails {
        // Check if the email is valid
        if check.IsEmailValid(email) {
            fmt.Printf("%s is a valid email.\n", email)
        } else {
            fmt.Printf("%s is an invalid email.\n", email)
        }
    }
}
