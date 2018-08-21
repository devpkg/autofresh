package autofresh

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/TerrenceHo/autofresh/config"
	"github.com/TerrenceHo/autofresh/runner"
	"github.com/TerrenceHo/autofresh/watchman"
)

const logo = `
    ___         __        ______               __  
   /   | __  __/ /_____  / ____/_______  _____/ /_ 
  / /| |/ / / / __/ __ \/ /_  / ___/ _ \/ ___/ __ \
 / ___ / /_/ / /_/ /_/ / __/ / /  /  __(__  ) / / /
/_/  |_\__,_/\__/\____/_/   /_/   \___/____/_/ /_/ 
`

const (
	watchmanCommand  = "watchman"
	directory        = "/Users/kho/go/src/github.com/TerrenceHo/autofresh/cmd"
	subscriptionName = "test_autofresh"
	buildCommand     = "go build /Users/kho/go/src/github.com/TerrenceHo/autofresh/cmd/main.go"
	runCommand       = "./main"
)

var (
	startChannel chan bool
	stopChannel  chan bool
)

// Main application start point. Will check if watchman exists at that path,
// retrieve the socket name, instantiate a connection to watchman using Unix
// Sockets, subscribe to the directory, and begin reading the subscription
// messages and building the executable.
func Start(conf config.Config) {
	if !conf.Hidebanner {
		fmt.Println(logo)
	}
	startChannel = make(chan bool)
	stopChannel = make(chan bool)
	isRunning := false

	watchman.Check(conf.Watchman)

	sockname := watchman.GetSockName(conf.Watchman)

	conn, err := net.Dial("unix", sockname)
	if err != nil {
		log.Fatalf("Dialing unix socket %s failed, error: %s\n",
			sockname, err.Error())
	}
	defer conn.Close()

	go read(conn, startChannel) // read in separate go routine

	// watchman.WatchProject(conn, directory)
	watchman.Subscribe(conn, directory, subscriptionName)

	for {
		start := <-startChannel
		if start {
			// kill previous process if necessary
			if isRunning {
				stopChannel <- true
			}

			buildErr, err := runner.Build(buildCommand)
			if err != nil {
				fmt.Println(buildErr)
				fmt.Println(err.Error())
				continue // skip run if build returned with error
			}

			// run here
			isRunning = runner.Run(runCommand, stopChannel)
		}
	}
}

func read(c net.Conn, startChannel chan bool) {
	d := json.NewDecoder(c)
	m := make(map[string]interface{})
	var start bool = true
	for {
		if err := d.Decode(&m); err != nil {
			if err != io.EOF {
				log.Fatalf("Error decoding, error: %s\n", err.Error())
			}
		}
		// fmt.Println(m)
		fmt.Println("Start", start)
		startChannel <- start
		start = !start
	}

}
