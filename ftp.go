package main

import (
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/spf13/viper"
)

type FTPClient interface {
	Store(filename string, data io.Reader) error
}

// Create New FTPClient
func NewFTPClient() FTPClient {

	return &ftpClient{}
}

type ftpClient struct {
	client *ftp.ServerConn
}

func (*ftpClient) Store(filename string, data io.Reader) error {
	c, err := ftp.Dial(viper.GetString("FTP_SERVER"), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Println("Error Dialing in FTP")
		return err
	}

	err = c.Login(
		viper.GetString("FTP_USERNAME"),
		viper.GetString("FTP_PASSWORD"),
	)
	if err != nil {
		log.Println("Error Logging in FTP")
		return err
	}

	return c.Stor(filename, data)
}
