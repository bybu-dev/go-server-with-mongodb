package emailRepo

import (
	"fmt"
	"net/smtp"
)

func sendEmail() {
	from := "from@gmail.com"
	password := "<Email Password>"
  
	to := []string{
	  "sender@example.com",
	}
  
	smtpHost := "smtp.gmail.com";
	smtpPort := "587";
  
	message := []byte("This is a test email message.");
	
	auth := smtp.PlainAuth("", from, password, smtpHost);
	
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message);
	if err != nil {
	  fmt.Println(err)
	  return
	}
	fmt.Println("Email Sent Successfully!")
  }
