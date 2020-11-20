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
	checkExecutionError(err)
	if HelmChartsManagerArgs.Debug {
		debugExecuteCommand(name, arg...)
	}
	if name == "echo" {
		return os.ExpandEnv(strings.Join(arg, " "))
	}

	command := exec.Command(name, arg...)
	command.Stdin = strings.NewReader(pipeCommandOutput)
	var byteArray []byte
	byteArray, err = command.Output()
	checkExecutionError(err)
	return byteArrayToString(print, byteArray)
}

func debugExecuteCommand(name string, arg ...string) {
	fmt.Println("DEBUG: Executing command:")
	fmt.Println("DEBUG:", name, strings.Join(arg, " "))
	fmt.Println(" ")
}

func checkExecutionError(err error) {
	if err != nil {
		_, ioErr := fmt.Fprintln(os.Stderr, err)
		if ioErr != nil {
			panic(ioErr)
		}
	}
}

func byteArrayToString(print bool, byteArray []byte) string {
	var result string
	if len(byteArray) > 0 && byteArray[len(byteArray)-1] == byte(10) {
		result = string(byteArray[:len(byteArray)-1])
	} else {
		result = string(byteArray)
	}
	if print && len(result) > 0 {
		fmt.Println(result)
	}
	return result
}
