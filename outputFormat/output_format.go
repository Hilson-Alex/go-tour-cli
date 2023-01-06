package outputFormat

// AsError just turns a string red for the console
func AsError(message string) string {
	return "\033[91m" + message + "\033[0m"
}
