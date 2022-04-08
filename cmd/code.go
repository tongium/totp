package cmd

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/tongium/totp/pkg/totp"
)

var CodeCmd = &cobra.Command{
	Use:   "code",
	Short: "Generate TOTP code",
	Args:  cobra.ExactArgs(1),
	Run:   gernerateCode,
}

func gernerateCode(cmd *cobra.Command, args []string) {
	secret, err := base32.StdEncoding.DecodeString(args[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(totp.TOTP(secret, time.Now(), 30, 6))
}
