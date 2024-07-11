package pkg

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(recipient, subject string) error {
	// Sender's email address and authentication
	user := os.Getenv("SMTP_USER")
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpServer := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Recipient's email address
	to := []string{recipient}

	body := `<html>
	<body>
		<h1>Order Confirmation</h1>
		<p>Your order has been confirmed</p>
	</body>
	</html>`

	// Message content
	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		from, recipient, subject, body,
	))

	// Authentication configuration
	auth := smtp.PlainAuth("", user, password, smtpServer)

	// Connect to the SMTP server
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent successfully")
	return nil
}
