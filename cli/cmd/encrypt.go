package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// CmdEncrypt encrypt command
var CmdEncrypt = &cobra.Command{
	Use:     "encrypt [string to encrypt]",
	Short:   "Print ciphertext",
	Example: "minoctl encrypt my.key foo",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		keyName := args[0]
		plainText := args[1]
		k, err := getKey(keyName)
		if err != nil {
			log.Fatal(err)
		}
		cipherText, err := k.Encrypt([]byte(plainText))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x", cipherText)
	},
}
