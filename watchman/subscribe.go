package watchman

import (
	"fmt"
	"log"
	"net"
)

// Subscribes to watchman's messages about a directory's filesystem changes,
// using an appropriate Unix Socket to watchman. Can take in the directory and
// a subscription name to configure the subscription.
func Subscribe(conn net.Conn, directory, subname string) {
	var cmd string

	cmd = fmt.Sprintf(`["subscribe", "%s", "%s", {"expression": ["allof", ["type", "f"], ["not", "empty"], ["suffix", "go"] ], "fields": ["name"] }]`,
		directory, subname)

	_, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
}

// Watches a project directory and all it's subdirectories. Currectly, this
// function is deprecated in favor of Subscribe, until further notice.
func WatchProject(conn net.Conn, directory string) {
	var cmd string

	cmd = fmt.Sprintf(`["watch", "%s"]`, directory)
	bytes, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
		log.Fatalf("Error writing to socket, error: %s\n", err.Error())
	}
	log.Printf("Bytes written: %d\n", bytes)
}
