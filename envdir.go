package envdir

import (
	"io/ioutil"
	"os"
	"os/exec"
	"fmt"
)

func ReadDir(dir string) (map[string]string, error) {
	envVariables := make(map[string]string{})
	directoryFiles, error = ioutil.ReadDir(dir)
	if error != nil {
		return envVariables, error
	}
	for _, file := range directoryFiles {
		fileInner, error := ioutil.ReadFile(file.Name)
		if error != nil {
			return envVariables, error
		}
		envVariables[file.Name()] = string(fileInner)
	}
	return envVariables, nil
}

func RunCmd(cmd []string, env map[string]string) int {
	commandParameters := make([]string{})
	commandParameters = append(commandParameters, "env")
	for paramName, paramValue := range env {
		commandParameters = append(commandParameters, string(paramName)+"="+string(paramValue))
	}
	for _, commandPiece := range cmd {
		commandParameters = append(commandParameters, string(commandPiece))
	}
	fmt.Printf("%v\n", commandParameters)
	command := exec.Command(commandParameters)
	return command.Run()
}
