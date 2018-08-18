package autofresh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"

	"github.com/TerrenceHo/autofresh/runner"
	"github.com/TerrenceHo/autofresh/watchman"
)

var (
	watchmanCommand  = "watchman"
	directory        = "/Users/kho/go/src/github.com/TerrenceHo/ABFeature"
	subscriptionName = "autofresh_watch"
	buildCommand     = "go build /Users/kho/go/src/github.com/TerrenceHo/ABFeature/cmd/ABFeature/main.go"
	runCommand       = "./main"

	startChannel chan bool
	endChannel   chan bool
)

func Start() {
	startChannel = make(chan bool)
	endChannel = make(chan bool)
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
		log.Println("Start", start)
		buildErr, err := runner.Build(buildCommand)
		if err != nil {
			fmt.Println(buildErr)
			fmt.Println(err.Error())
			continue // skip run if build returned with error
		}
		// kill previous process if necessary
		// run here
	}
}

func read(c net.Conn, startChannel chan bool) {
	d := json.NewDecoder(c)
	m := make(map[string]interface{})
	for {
		if err := d.Decode(&m); err != nil {
			if err != io.EOF {
				log.Fatalf("Error decoding, error: %s\n", err.Error())
			}
			continue
		}
		log.Println("Client got", m)
		startChannel <- true
	}
}

func run(runCommand string) {
	cmd := exec.Command(runCommand)

	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error with StdoutPipe, error: %s\n", err.Error())
	}

	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error with StderrPipe, error: %s\n", err.Error())
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Failed to start command %s, error: %s\n", runCommand, err.Error())
	}

	var errStdout, errStderr error
	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := io.MultiWriter(stdout, &stdoutBuf)
	stderr := io.MultiReader(stderr, &stderrBuf)

	go func() {
		_, stdoutErr := io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, stderrErr := io.Copy(stderr, stderrIn)
	}()

	if err := cmd.Wait(); err != nil {
		log.Fatal("Running command failed, error %s\n", err.Error())
	}

}
