package autofresh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"time"
)

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

type WatchmanSockName struct {
	Version  string `json:"version"`
	Sockname string `json:"sockname"`
}

func toJSON(obj interface{}) string {
	str, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(str)
}

func fromJSON(str string) interface{} {
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	var obj interface{}
	d.Decode(&obj)
	return obj
}

func read(c net.Conn) {
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
	}
}

func Start() {
	sockname := GetSockName("watchman")

	conn, err := net.Dial("unix", sockname)
	if err != nil {
		log.Fatalf("Dialing unix socket %s failed, error: %s\n",
			sockname, err.Error())
	}
	defer conn.Close()

	go read(conn) // read in separate go routine

	cmd := `["watch", "/Users/kho/go/src/github.com/TerrenceHo/ABFeature"]`
	bytes, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytes)
	time.Sleep(2 * time.Second)

	cmd = `["subscribe", "/Users/kho/go/src/github.com/TerrenceHo/ABFeature", "autofresh_watch", {"expression": ["allof", ["type", "f"], ["not", "empty"], ["suffix", "go"] ],"fields": ["name"]}]`

	bytes, err = conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytes)

	log.Println("Spinning forever")
	for {
	}
}
