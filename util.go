package main

import (
	"os"
	"os/exec"
	"strings"
)

func getLocalDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func getExistingProjectNameAndParentDirectory() (string, string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	name := strings.Split(dir, "\\")

	parentDirectory := strings.Join(name[:len(name)-1], "\\")
	return name[len(name)-1], parentDirectory, nil
}

func createDirectoryIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func RunCommand(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
