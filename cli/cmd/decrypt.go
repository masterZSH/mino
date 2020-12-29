package cmd

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// CmdDecrypt decrypt command
var CmdDecrypt = &cobra.Command{
	Use:     "decrypt [string to decrypt]",
	Short:   "ciphertext to plaintext",
	Args:    cobra.MinimumNArgs(2),
	Example: "minoctl decrypt my.key foo",
	Run: func(cmd *cobra.Command, args []string) {
		keyName := args[0]
		cipherText := args[1]
		k, err := getKey(keyName)
		if err != nil {
			log.Fatal(err)
		}
		cipherBytes, err := hex.DecodeString(cipherText)
		if err != nil {
			log.Fatal(err)
		}

		plainText, err := k.Decrypt(cipherBytes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", plainText)
	},
}
