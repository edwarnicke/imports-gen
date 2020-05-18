package main_test

import (
	"os"
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
	importsFilename := "./testdata/imports.go"
	err = os.Remove(importsFilename)
	if err != nil {
		t.Errorf("failed to remove %q: %+v", importsFilename, err)
	}
}
