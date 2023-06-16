package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const (
	appName  = "test build"
	fileName = "test file(need to be removed)"
)

func TestMain(m *testing.M) {
	fmt.Println("-> Building...")

	build := exec.Command("go", "build", "-o", appName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building %s: %s", appName, err)
		os.Exit(1)
	}
	fmt.Println("-> Running...")
	// Running the test
	result := m.Run()
	fmt.Println("-> Getting done...", result)
	os.Remove(appName)
	os.Remove(fileName)
	os.Exit(result)
}
