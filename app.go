package autofresh

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"

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

func Start(conf config.Config) {
	fmt.Println(logo)
	startChannel = make(chan bool)
	stopChannel = make(chan bool)
	isRunning := false

	_, err := exec.LookPath(watchmanCommand)
	if err != nil {
		log.Fatalf("Failed to find watchman executable. Error: %s\n", err.Error())
	}

	sockname := watchman.GetSockName("watchman")

	conn, err := net.Dial("unix", sockname)
	if err != nil {
		log.Fatalf("Dialing unix socket %s failed, error: %s\n",
			sockname, err.Error())
	}
	defer conn.Close()

	// watchman.WatchProject(conn, directory)
	watchman.Subscribe(conn, directory, subscriptionName)

	go read(conn, startChannel) // read in separate go routine

	log.Println("Spinning forever")
	for {
		start := <-startChannel
		if start {
			buildErr, err := runner.Build(buildCommand)
			if err != nil {
				fmt.Println(buildErr)
				fmt.Println(err.Error())
				continue // skip run if build returned with error
			}

			// kill previous process if necessary
			if isRunning {
				stopChannel <- true
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
