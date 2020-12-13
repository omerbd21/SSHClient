package linux

import (
	"fmt"
	"os/exec"
)

func viewPrograms(os string) string {
	var programs string
	switch os {
	case "RedHat":
		command := exec.Command("yum", "list",  "installed")
		output, err := command.CombinedOutput()
		programs = fmt.Sprintln(string(output))
		if err == nil {
			return err.Error()
		}
	case "Debian":
		command := exec.Command("apt","list", "--installed")
		output, err := command.CombinedOutput()
		programs = fmt.Sprintln(string(output))
		if err != nil {
			return err.Error()
		}
	}
	return programs
}