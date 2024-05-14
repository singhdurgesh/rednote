package notifications

import (
	"fmt"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/app/services/otp_handler"
	"github.com/singhdurgesh/rednote/internal/tasks"
)

type Recepients struct {
	Emails []string
	Phones []string
}

type Communication struct {
	Content     string
	ContentType string
	TemplateId  string
	Recepients
}

const (
	loginOTP              = "login_otp"
	loginResendOtp        = "login_resend_otp"
	CommunicationTaskName = "communication"
)

func NewCommunication(emails []string, phones []string, template string) *Communication {
	return &Communication{
		TemplateId: template,
		Recepients: Recepients{
			Emails: emails,
			Phones: phones,
		},
	}
}

func NewLoginOtpCommunication(user models.User) *Communication {
	phones := []string{user.Phone.String}
	emails := []string{user.Email.String}

	return NewCommunication(emails, phones, loginOTP)
}

func NewResendLoginOtpCommunication(user models.User) *Communication {
	phones := []string{user.Phone.String}
	emails := []string{user.Email.String}

	return NewCommunication(emails, phones, loginResendOtp)
}

func ProcessCommunication(data string) (bool, error) {
	task := &Communication{}
	return tasks.ProcessTask(task, data)
}

func (n *Communication) Name() string {
	return CommunicationTaskName
}

func (c *Communication) Run() error {
	// Generate or Fetch Content
	// Call the SMS Vendor Service and Email Service to Send the Email

	if c.TemplateId == loginOTP {
		otp_handler.SendLoginOtp(c.Phones[0])
	} else if c.TemplateId == loginResendOtp {
		otp_handler.ResendLoginOtp(c.Phones[0])
	} else {
		app.Logger.Error("Invalid Template Id", c.TemplateId)
	}

	fmt.Println(c)
	return nil
}
