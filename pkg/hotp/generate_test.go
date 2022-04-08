package hotp_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"testing"
	"time"

	"github.com/tongium/totp/pkg/hotp"
)

var hmacResult []byte

func setup() {
	counter := time.Now().Unix() / 30
	secret, err := base32.StdEncoding.DecodeString("3ZBC3KMPUWZMRGYLSKIMKC3FPPSZHU5U")
	if err != nil {
		panic(err)
	}

	h := hmac.New(sha1.New, secret)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(counter))
	h.Write(buf)
	hmacResult = h.Sum(nil)
}

func BenchmarkDynamicTruncateWithDumbImplement(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		hotp.DynamicTruncateExplaination(hmacResult)
	}
}

func BenchmarkDynamicTruncate(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		hotp.DynamicTruncate(hmacResult)
	}
}
