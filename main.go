package main

import (
	"log"
	"os"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/spf13/viper"
)

func main() {
	// Load Config
	viper.SetConfigFile("./.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error Loading Config file .env")
		log.Fatal(err)
	}

	// Start SMTP server
	s := smtp.NewServer(&SMTPHandlers{NewFTPClient()})

	s.Addr = ":25"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.Debug = os.Stdout
	s.AllowInsecureAuth = true

	log.Printf("Starting server at: %s", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
