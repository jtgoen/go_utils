package go_utils

import "os"

func Bind_To_Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
