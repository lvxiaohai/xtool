package fileserver

import (
	"github.com/spf13/cobra"
	"xtool/pkg/fileserver"
)

var (
	dir  string
	port int
)

var Cmd = &cobra.Command{
	Use:     "fileserver",
	Aliases: []string{"fs"},
	Short:   "静态服务器",
	Run:     run,
}

func init() {
	Cmd.Flags().StringVarP(&dir, "dir", "d", ".", "文件目录")
	Cmd.Flags().IntVarP(&port, "port", "p", 9001, "监听端口")
}

func run(_ *cobra.Command, _ []string) {
	fileserver.FileServer(dir, port)
}
