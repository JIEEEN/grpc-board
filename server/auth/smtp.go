package auth

import (
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

func SendMail(_to string) (string, error) {
	from := os.Getenv("SMTP_SENDER_EMAIL")
	password := os.Getenv("SMTP_SENDER_PASSWORD")
	host := os.Getenv("SMTP_HOST_SERVER")
	port := os.Getenv("SMTP_HOST_PORT")
	to := []string{_to}
	subject := "[GRPC-BOARD] Verification Code"

	msg, randomString := mailBody(from, _to, subject)
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)

	endpoint := host + ":" + port
	err := smtp.SendMail(endpoint, auth, from, to, body)
	if err != nil {
		return "", err
	}

	return randomString, nil
}

func mailBody(from string, to string, subject string) (string, string) {
	randomString := generateRandomString()
	msg := `
	<html>
		<head>
			<title>Verification Code</title>
		</head>
		<body>
			<h1>Verification Code</h1>
			<p>Your verification code is: ` + randomString + `</p>
		</body>
	</html>
	`

	msg = "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=\"utf-8\"\n" +
		"\n" +
		msg

	return msg, randomString
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
