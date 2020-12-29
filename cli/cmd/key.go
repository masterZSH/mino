package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/masterZSH/mino"
	"github.com/spf13/cobra"
)

func getKey(keyName string) (k mino.Key, err error) {
	var keyStr string
	key, err := ioutil.ReadFile(keyName)
	if err != nil {
		keyStr = keyName
	} else {
		keyStr = string(key)
	}
	tempArr := strings.Split(keyStr, "-")
	if len(tempArr) != 2 {
		err = errors.New("error key")
		return
	}
	k, err = mino.NewKey([]byte(tempArr[0]), []byte(tempArr[1]))
	return
}

// CmdKey key command
var CmdKey = &cobra.Command{
	Use:     "key [genearate random key file]",
	Short:   "genearate random key file",
	Args:    cobra.MinimumNArgs(1),
	Example: "minoctl key my.key",
	Run: func(cmd *cobra.Command, args []string) {
		pass := getUUID()
		salt := getUUID()
		err := ioutil.WriteFile(args[0], []byte(pass+"-"+salt), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("create key file success")
	},
}

func getUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
