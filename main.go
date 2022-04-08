package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
	"github.com/tongium/totp/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "totp",
	Short: "Time-based one-time password",
}

func init() {
	rootCmd.AddCommand(cmd.CodeCmd)
}

func main() {
	gotenv.Load()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
