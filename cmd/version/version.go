package version

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"xtool/pkg/version"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "显示版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		s := version.Get()
		bytes, err := json.MarshalIndent(&s, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(bytes))
	},
}
