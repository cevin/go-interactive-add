package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"os/exec"
	"strings"
)

var DIR string

func main() {
	DIR, _ = os.Getwd()
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "iadd" {
		files := getUntrackedFiles()
		if len(files) == 0 {
			fmt.Println("no untracked files found")
			return
		}

		var selected []string

		prompt := &survey.MultiSelect{
			Message: "Select files to track",
			Options: files,
		}
		_ = survey.AskOne(prompt, &selected)

		if len(selected) == 0 {
			return
		}

		addFile(selected...)

	} else {
		git(strings.Join(args, " "))
	}

}

func addFile(files ...string) {
	git(fmt.Sprintf("add %s", strings.Join(files, " ")))
}

func getUntrackedFiles() []string {
	var files []string

	command := exec.Command("git", strings.Split("ls-files -o --exclude-standard", " ")...)
	command.Dir = DIR
	output, err := command.CombinedOutput()

	if err != nil {
		return files
	}

	files = strings.Split(strings.Trim(string(output), "\n"), "\n")

	return files
}

func git(args string) {
	command := exec.Command("git", strings.Split(args, " ")...)
	command.Dir = DIR
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	_ = command.Run()
}
