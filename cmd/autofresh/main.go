package main

import (
	"fmt"
	"os"
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
	flags.BoolP("hidebanner", "h", false, "hide autofresh banner")
	flags.StringP("watchman", "w", "watchman", "path to watchman executable")
	flags.StringP("watchdir", "d", "./", "path to directory for autofresh to watch")
	flags.StringP("port", "p", "31337", "set HTTP port of autofresh will run on, between 1024-65535.")

}

func main() {
	mainCmd.Execute()
}

func serve(cmd *cobra.Command, args []string) {
	conf := config.LoadConfig(cmd)

	if conf.Version {
		fmt.Printf("autofresh %s %s %s/%s\n", Version, GoVersion, GOOS, GOARCH)
		fmt.Printf("git hash: %s\n", GitHash)
		fmt.Printf("built at: %s\n", BuildTime)
		os.Exit(0)
	}
	fmt.Println("Port", conf.Port)

	// autoFresh.Start(conf)
}
