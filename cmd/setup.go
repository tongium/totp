package cmd

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/tongium/totp/pkg/totp"
)

var issuer string
var accountName string

var SetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Generate setup URI",
	Run:   generateSetupURI,
}

func init() {
	SetupCmd.Flags().StringVarP(&issuer, "issuer", "i", "", "specify issuer")
	SetupCmd.Flags().StringVarP(&accountName, "account", "a", "", "specify account name")
}

func generateSetupURI(cmd *cobra.Command, args []string) {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, 20)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	fmt.Println(totp.CreateSetupURI(key, issuer, accountName, 6, 30))
}

func generateKey() string {
	rand.Seed(time.Now().UnixNano())
	key := make([]byte, 20)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	return base32.StdEncoding.EncodeToString(key)
}
