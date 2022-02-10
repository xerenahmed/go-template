package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Hello World")
}