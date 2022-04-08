package hotp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

func HOTP(key []byte, counter []byte, digit int) string {
	// Step 1: Generate an HMAC-SHA-1 value Let HS = HMAC-SHA-1(K,C)
	hmacResult := GenerateHMAC(key, counter)

	// Step 2: Generate a 4-byte string (Dynamic Truncation)
	code := DynamicTruncate(hmacResult)

	//  Step 3: Compute an HOTP value, Return D = Snum mod 10^Digit
	value := code % int64(math.Pow10(digit))

	format := fmt.Sprintf("%%0%dd", digit)
	return fmt.Sprintf(format, value)
}

func GenerateHMAC(key []byte, message []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(message)
	return h.Sum(nil)
}

func DynamicTruncateExplaination(hmacResult []byte) int64 {
	offset := hmacResult[len(hmacResult)-1] & 0xf
	binCode := hmacResult[offset : offset+4]
	binCode[0] = binCode[0] & 0x7f

	var code uint32
	buf := bytes.NewReader(binCode)
	err := binary.Read(buf, binary.BigEndian, &code)
	if err != nil {
		panic(err)
	}

	return int64(code)
}

func DynamicTruncate(hmacResult []byte) int64 {
	// https://datatracker.ietf.org/doc/html/rfc4226#section-5.4
	offset := hmacResult[len(hmacResult)-1] & 0xf
	return int64(((int(hmacResult[offset]) & 0x7f) << 24) |
		((int(hmacResult[offset+1] & 0xff)) << 16) |
		((int(hmacResult[offset+2] & 0xff)) << 8) |
		(int(hmacResult[offset+3]) & 0xff))
}
