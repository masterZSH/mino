package main

import (
	"github.com/masterZSH/mino/cli/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mino",
	Short: "mino is a very Fast&Secure encrypt&decrypt data tool",
}

func init() {
	rootCmd.AddCommand(cmd.CmdEncrypt, cmd.CmdDecrypt, cmd.CmdKey)
}

func main() {
	rootCmd.Execute()
}
