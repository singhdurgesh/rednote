package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the secret key (should be kept secure)
var secretKey = "LF2FSZBZMVZFGVDXO54DM2KOKRWWQQKLJJSS6VSOLFUVEVSQMYXVEWSOOVJWIYSH"

// Define the length of the OTP
const otpLength = 6

// Function to generate the HOTP
func generateHOTP(secret string, counter uint64) (string, error) {
	fmt.Println(secret)
	key, err := base32.StdEncoding.DecodeString(secret)
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

// Function to validate the OTP
func validateOTP(otp string, secret string, counter uint64) bool {
	generatedOTP, err := generateHOTP(secret, counter)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return false
	}
	return otp == generatedOTP
}

// Function to get OTP from cache or generate a new one
func getOTPFromCache(phoneNumber string) (string, uint64, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := rdb.Context()

	// Fetching the counter and OTP from cache
	counter, err := rdb.Get(ctx, phoneNumber+"_counter").Uint64()
	if err == redis.Nil {
		// Counter not found in cache, initializing to 0
		counter = 0
	} else if err != nil {
		return "", 0, err
	}

	otp, err := rdb.Get(ctx, phoneNumber+"_otp").Result()
	if err == redis.Nil {
		// OTP not found in cache, generate a new one
		newOTP, err := generateHOTP(secretKey, counter)
		if err != nil {
			return "", 0, err
		}
		// Cache the OTP and counter
		err = rdb.Set(ctx, phoneNumber+"_otp", newOTP, 10*time.Minute).Err()
		if err != nil {
			return "", 0, err
		}
		err = rdb.Set(ctx, phoneNumber+"_counter", counter, 10*time.Minute).Err()
		if err != nil {
			return "", 0, err
		}
		return newOTP, counter, nil
	} else if err != nil {
		return "", 0, err
	}
	return otp, counter, nil
}

func main() {
	phoneNumber := "user_phone_number_here"

	otp, counter, err := getOTPFromCache(phoneNumber)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Generated OTP:", otp)

	// Simulating OTP verification
	// Assuming user inputs the OTP for verification
	// We can directly validate the OTP without fetching it from cache again
	isValid := validateOTP(otp, secretKey, counter)
	if isValid {
		fmt.Println("OTP is valid!")
	} else {
		fmt.Println("OTP is invalid!")
	}
}

// pow10 computes 10 raised to the power of n.
func pow10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
