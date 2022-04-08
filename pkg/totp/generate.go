package totp

import (
	"encoding/binary"
	"time"

	"github.com/tongium/totp/pkg/hotp"
)

func TOTP(secret []byte, timestamp time.Time, periodSecond int64, digit int) string {
	t := timestamp.Unix() / periodSecond

	counter := make([]byte, 8)
	binary.BigEndian.PutUint64(counter, uint64(t))

	return hotp.HOTP(secret, counter, digit)
}
