package totp

import (
	"encoding/base32"
	"fmt"
)

func CreateSetupURI(key []byte, issuer, accountName string, digits int, periodSecond int) string {
	secret := base32.StdEncoding.EncodeToString(key)
	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s&algorithm=sha-1&digits=%d&period=%d", issuer, accountName, secret, issuer, digits, periodSecond)
}
