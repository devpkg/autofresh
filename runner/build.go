package runner

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Build(buildString string) (string, error) {
	var cmd *exec.Cmd
	buildArr := strings.Split(buildString, " ")
	buildCmd := buildArr[0]
	buildArgs := buildArr[1:]
	if len(buildArr) > 1 {
		cmd = exec.Command(buildCmd, buildArgs...)
	} else {
		cmd = exec.Command(buildCmd)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error with StdoutPipe, error: %s\n", err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error with StderrPipe, error: %s\n", err.Error())
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Error with starting Command, error: %s\n", err.Error())
	}

	io.Copy(os.Stdout, stdout)
	buildErr, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		return string(buildErr), err
	}

	return "", nil
}
