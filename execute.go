package main

import (
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(print bool, name string, arguments string) string {
	pipeCommandOutput := ""
	if strings.Contains(arguments, " | ") {
		splittedCommandsArguments := strings.SplitN(arguments, " | ", 2)
		pipeCommandOutput = executeCommand(print, name, splittedCommandsArguments[0])
		mainCommand := strings.SplitN(splittedCommandsArguments[1], " ", 2)
		name = mainCommand[0]
		arguments = mainCommand[1]
	}
	arg, err := shlex.Split(arguments)
	if err != nil {
		fmt.Println(err)
	}
	if HelmChartsManagerArgs.Debug {
		debugExecuteCommand(name, arg...)
	}
	if name == "echo" {
		return os.ExpandEnv(strings.Join(arg, " "))
	}

	command := exec.Command(name, arg...)
	command.Stdin = strings.NewReader(pipeCommandOutput)
	var barr, _ = command.CombinedOutput()
	var ret string
	if len(barr) > 0 && barr[len(barr)-1] == byte(10) {
		ret = string(barr[:len(barr)-1])
	} else {
		ret = string(barr)
	}
	if print && len(ret) > 0 {
		fmt.Println(ret)
	}
	return ret
}

func debugExecuteCommand(name string, arg ...string) {
	fmt.Println("DEBUG: Executing command:")
	fmt.Println("DEBUG:", name, strings.Join(arg, " "))
	fmt.Println(" ")
}
