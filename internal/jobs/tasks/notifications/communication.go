package notifications

import (
	"fmt"

	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/app/services/text_notifications.go"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
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
		// TODO: Create a New Service for Sending OTP
		otp, err := text_notifications.LoginOTPGenerator(c.Phones[0])

		if err != nil {
			logger.LogrusLogger.Println("OTP Generator Error: ", err)
			return err
		}

		args := map[string]interface{}{"Otp": otp}
		otpText, err := text_notifications.GetContentFromTemplate(c.TemplateId, args)
		if err != nil {
			logger.LogrusLogger.Println("Get Content Error: ", err)
			return err
		}

		err = text_notifications.SmsSender.SendSms(c.Phones, otpText)

		if err != nil {
			logger.LogrusLogger.Println("SMS Sender Error: ", err)
			return err
		}
	}

	fmt.Println(c)
	return nil
}
