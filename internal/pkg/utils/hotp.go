package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"

	"github.com/singhdurgesh/rednote/cmd/app"
)

const otpLength = 6

func ValidateOTP(otp string, counter uint64) bool {
	generatedOtp, err := GenerateHOTP(counter)

	if err != nil {
		return false
	}

	return otp == generatedOtp
}

func GenerateHOTP(counter uint64) (string, error) {
	key, err := base32.StdEncoding.DecodeString(app.Config.App.OtpSecret)
	if err != nil {
		return "", err
	}

	// Convert counter to byte array
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, counter)

	// HMAC SHA-1 hashing
	mac := hmac.New(sha1.New, key)
	mac.Write(counterBytes)
	hash := mac.Sum(nil)

	// Generate the OTP
	offset := hash[len(hash)-1] & 0xf
	binary := binary.BigEndian.Uint32(hash[offset : offset+4])
	otp := binary % uint32(pow10(otpLength))

	return fmt.Sprintf("%06d", otp), nil
}

// pow10 computes 10 raised to the power of n.
func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
