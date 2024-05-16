package utils

import (
	"fmt"
	"testing"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/configs"
)

func TestValidateOTP(t *testing.T) {
	app.Config = &configs.Config{Jwt: configs.Jwt{Secret: "Text"}}

	valid := ValidateOTP("", 1)
	if valid == true {
		t.Errorf("Expected false but got %v", valid)
	}

	counter := uint64(1)
	otp, _ := GenerateHOTP(counter)
	// OTP is passed as blank
	fmt.Println(t)
	valid = ValidateOTP(otp, counter+1)

	if valid == true {
		t.Errorf("Expected false but got %v", valid)
	}

	valid = ValidateOTP(otp, counter)
	if valid == false {
		t.Errorf("Expected true but got %v", valid)
	}
}
