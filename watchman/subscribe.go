package watchman

import (
	"fmt"
	"log"
	"net"
)

// Subscribes to a directory's filesystem changes
func Subscribe(conn net.Conn, directory, subname string) {
	var cmd string

	cmd = fmt.Sprintf(`["subscribe", "%s", "%s", {"expression": ["allof", ["type", "f"], ["not", "empty"], ["suffix", "go"] ], "fields": ["name"] }]`,
		directory, subname)

	bytes, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytes)
}

// Watches a project directory and all it's subdirectories
func WatchProject(conn net.Conn, directory string) {
	var cmd string

	cmd = fmt.Sprintf(`["watch", "%s"]`, directory)
	bytes, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytes)
}
