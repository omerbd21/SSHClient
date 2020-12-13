package linux

import (
	"fmt"
	"os/exec"
	"strings"
)

func viewPrograms(os string) ([]string, error) {
	var allPrograms string
	switch os {
	case "RedHat":
		command := exec.Command("yum", "list",  "installed")
		output, err := command.CombinedOutput()
		allPrograms = fmt.Sprintln(string(output))
		if err == nil {
			return nil,err
		}
	case "Debian":
		command := exec.Command("apt","list", "--installed")
		output, err := command.CombinedOutput()
		allPrograms = fmt.Sprintln(string(output))
		if err != nil {
			return nil, err
		}
	}
	programs := strings.Split(allPrograms, "\n")
	return programs[1:], nil
}