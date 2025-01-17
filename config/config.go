package config

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}

func GetDatabaseConnectionString() string {
	return os.Getenv("DB_CONNECTION")
}
