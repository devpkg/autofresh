package main

import (
	"fmt"
	"runtime"

	"github.com/TerrenceHo/autofresh/config"
	"github.com/spf13/cobra"
)

const (
	GOARCH string = runtime.GOARCH
	GOOS   string = runtime.GOOS
)

var (
	Version   string
	GitHash   string
	BuildTime string
	GoVersion string = runtime.Version()
)

var mainCmd = &cobra.Command{
	Use:   "autofresh",
	Short: "Autofresh: live reloading server",
	// Run: func(cmd *cobra.Command, args []string) {
	// serve()
	// },
	Run: serve,
}

func init() {
	flags := mainCmd.Flags()

	flags.BoolP("version", "v", false, "print version of autofresh")
	flags.StringP("file", "f", "autofresh-config.json", "set path to configuration file.")
	flags.StringP("port", "p", "31337", "set HTTP port of autofresh will run on, between 1024-65535.")

}

func main() {
	mainCmd.Execute()
}

func serve(cmd *cobra.Command, args []string) {
	conf := config.LoadConfig(cmd)

	fmt.Println("Port", conf.Port)

	// version check
	// autoFresh.Start(conf)
}
