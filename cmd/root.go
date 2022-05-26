package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"xtool/cmd/version"
)

var root = &cobra.Command{
	Use: "xtool",
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	root.AddCommand(version.Cmd)
}
