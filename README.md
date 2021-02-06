# SMTP to FTP server

Got annoyed that some companies feeds use email and doesn't have a ftp option so decided to create a SMTP server just to convert their feeds (attachments) to FTP.

## Setup

Rename `example.env` to `.env`

```bash
EMAIL_SUFFIX="myemail@email.com" # only matching email suffix can be received

FTP_SERVER="localhost:21" # ftp server (recommend using local server)
FTP_USERNAME="username"
FTP_PASSWORD="password"
```

Setup SMTP server on 25 port and MX for domain DNS.
