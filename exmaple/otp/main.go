package main

import (
	"fmt"

	"github.com/pquerna/otp/totp"
)

func main() {
	// Generate a new key
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "company",
		AccountName: "username@gmail.com",
		Period:      30,
		Digits:      6,
	})

	// Setup
	// otpauth://totp/company:username@gmail.com?algorithm=SHA1&digits=6&issuer=company&period=30&secret=B4VWH...
	fmt.Println(key.URL())

	// Code from user
	passcode := "123456"
	// Validate user's code
	fmt.Println(totp.Validate(passcode, key.Secret()))

	// Generate QR code
	// qr, _ := key.Image(40, 40)
}
