package main_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/edwarnicke/exechelper"
)

func TestImportsGen(t *testing.T) {
	cmdStr := "go generate ./testdata"
	options := []*exechelper.Option{
		exechelper.WithStderr(os.Stderr),
		exechelper.WithStdout(os.Stdout),
		exechelper.WithEnvirons(os.Environ()...),
	}
	err := exechelper.Run(cmdStr, options...)
	if err != nil {
		t.Errorf("failed to run %q: %+v", cmdStr, err)
	}
	cmdStr = "go build ./testdata"
	err = exechelper.Run(cmdStr, options...)
	if err != nil {
		t.Errorf("failed to run %q: %+v", cmdStr, err)
	}

	// Extract the GOOS
	goos, ok := os.LookupEnv("GOOS")
	if !ok {
		cmdStr := "go env GOOS"
		options := []*exechelper.Option{
			exechelper.WithStderr(os.Stderr),
			exechelper.WithStdout(os.Stdout),
			exechelper.WithEnvirons(os.Environ()...),
		}
		goosBytes, GoosErr := exechelper.Output(cmdStr, options...)
		if GoosErr != nil {
			log.Fatalf("error extracting GOOS: %+v", GoosErr)
		}
		goos = strings.TrimSpace(string(goosBytes))
	}

	importsFilename := fmt.Sprintf("./testdata/%s", fmt.Sprintf("imports_%s.go", goos))
	err = os.Remove(importsFilename)
	if err != nil {
		t.Errorf("failed to remove %q: %+v", importsFilename, err)
	}
}
