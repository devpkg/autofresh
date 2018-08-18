package watchman

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type WatchmanSockName struct {
	Version  string `json:"version"`
	Sockname string `json:"sockname"`
}

func GetSockName(watchmanPath string) string {
	cmd := exec.Command("watchman", "get-sockname")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Watchman check failed, error\n%s", err.Error())
	}

	var wsn WatchmanSockName
	if err := json.NewDecoder(&stdout).Decode(&wsn); err != nil {
		log.Fatalf("Socket name did not decode correctly. Error: %s\n", err.Error())
	}

	return wsn.Sockname
}
