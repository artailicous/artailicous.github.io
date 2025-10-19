package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	Border := strings.Repeat("*", numStarsPerLine)
    return Border + "\n" + welcomeMsg + "\n" + Border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// 1. Remove all instances of the asterisk character
	msgWithoutAsterisks := strings.ReplaceAll(oldMsg, "*", "")

	// 2. Trim leading and trailing whitespace from the resulting string
	cleanedMsg := strings.TrimSpace(msgWithoutAsterisks)

	return cleanedMsg
}
