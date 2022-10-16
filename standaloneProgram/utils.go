package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type command struct {
	Custom,
	Default string
}

var (
	Gifsicle = command{Default: "gifsicle"}
	Blender  = command{Default: "blender"}
)

func (c command) Run(args ...string) error {
	cmd := exec.Command(c.GetCommand(), args...)
	var out bytes.Buffer
	cmd.Stderr = &out
	e := cmd.Run()
	fmt.Println(out.String())
	return e
}

func (c command) GetCommand() string {
	if c.Custom != "" {
		return c.Custom
	}
	//if runtime.GOOS == "windows" {
	//	return c.Default + ".exe"
	//}
	return c.Default
}

func (c command) Exists() bool {
	return commandExists(c.GetCommand())
}

func commandExists(cmd string) bool {
	if s, err := os.Stat(cmd); errors.Is(err, os.ErrNotExist) {
		return s.Mode()|0111 != 0
		//return true
	}
	_, err := exec.LookPath(cmd)
	return err == nil
}
