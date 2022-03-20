package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	upperCustomer := strings.ToUpper(customer)
	message := "Welcome to the Tech Palace, " + upperCustomer
	
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsBorder := strings.Repeat("*", numStarsPerLine)
	message := fmt.Sprintf(
		`%s
%s
%s`, starsBorder, welcomeMsg, starsBorder)

	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	messageArray := strings.Split(oldMsg, "\n")
	newMessage := strings.TrimSpace(strings.ReplaceAll(messageArray[1], "*", ""))

	return newMessage
}
