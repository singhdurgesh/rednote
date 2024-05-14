package text_notifications

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
)

const (
	loginOTP = "{{.Otp}} is Your Login OTP for the AppName. Use this to Login into your Account."
)

func LoginOTPGenerator(phone string) (string, error) {
	// Generate OTP with Phone and TimeStamp()
	// Cache the OTP for 1 Minutes to be used for Resender Serivce
	return phone[1:7], nil
}

func GetContentFromTemplate(templateId string, args map[string]interface{}) (string, error) {
	if templateId != "login_otp" {
		return "", errors.New("invalid Template Id")
	}

	tmpl, err := template.New("loginOtp").Parse(loginOTP)
	if err != nil {
		panic(err)
	}

	var result bytes.Buffer
	tmpl.Execute(&result, args)
	return result.String(), nil
}

var SmsSender *SmsVendor

// it can be an insterface as well
type SmsVendor struct {
	Provider string
	UserName string
	Password string
}

func (s *SmsVendor) SendSms(phones []string, content string) error {
	// Call the third party API
	fmt.Printf("OTP Processed Succefully with Phone: %v, content: %v", phones, content)
	return nil
}
