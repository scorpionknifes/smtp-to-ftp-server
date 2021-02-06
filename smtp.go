package main

import (
	"io"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/spf13/viper"
)

// SMTPHandlers implements SMTP server methods.
type SMTPHandlers struct {
	ftpClient FTPClient
}

// Login handles a login command with username and password.
func (handlers *SMTPHandlers) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	return &Session{handlers.ftpClient}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (handlers *SMTPHandlers) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{handlers.ftpClient}, nil
}

// A Session is returned after successful login.
type Session struct {
	ftpClient FTPClient
}

// Mail mail handler
func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	return nil
}

// Rcpt rcpt handler
func (s *Session) Rcpt(to string) error {
	if !strings.HasSuffix(to, viper.GetString("EMAIL_SUFFIX")) {
		return smtp.ErrAuthRequired
	}
	return nil
}

// Data data handler
func (s *Session) Data(r io.Reader) error {
	email, err := Parse(r) // returns Email struct and error
	if err != nil {
		return err
	}

	currentTime := time.Now().Format("20060102150405")

	if len(email.Attachments) == 0 {
		return s.ftpClient.Store(currentTime+"_"+email.Subject, email.Content)
	}

	for _, a := range email.Attachments {
		err := s.ftpClient.Store(currentTime+"_"+a.Filename, a.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Reset reset
func (s *Session) Reset() {}

// Logout logout
func (s *Session) Logout() error {
	return nil
}
