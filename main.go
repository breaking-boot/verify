package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Printf("🔎 Verify started.\n")

	realBootDevPath, _ := exec.LookPath("bootdev")
	myVerifyPath, _ := os.Executable()

	if realBootDevPath == myVerifyPath {
		// Only happens if user renamed 'verify' to 'bootdev'
		// and put it earlier in the PATH than the official bootdev cli.
		fmt.Println("Danger! 🔎 Verify is trying to call itself. Check your PATH.")
		return
	}

	args := os.Args[1:]

	isRun := false
	isSubmit := false

	if len(args) > 0 && args[0] == "run" {
		isRun = true
		// Check if any of the other arguments are "-s" or "--submit"
		for _, arg := range args {
			if arg == "-s" || arg == "--submit" {
				isSubmit = true
				break
			}
		}
	}

	// Only intercept if command is 'run' WITHOUT an '-s' or '--submit' flag
	if isRun && !isSubmit {
		handleRun(realBootDevPath, args)
		return
	}

	// Otherwise, let the official CLI handle it
	passThrough(realBootDevPath, args)
}

func handleRun(bin string, args []string) {

	fmt.Printf("🐻 Official bootdev cli <run> command detected.\n")
	fmt.Printf("🔎 Verify intercepting...\n\n")

	cmd := exec.Command(bin, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// If the bootdev command fails (e.g., code has a bug)
		fmt.Printf("The 🐻 official bootdev cli reported an error: %v\n", err)
	}

	lines := strings.Split(string(output), "\n")

	testCounter := 0

	for _, line := range lines {
		//fmt.Printf("*** DEBUG Line %d: %q\n", i, line) // DEBUGGING LINE, must add index i back to for loop definition

		if strings.HasSuffix(line, "╮") {
			// lengthen the line
			line = strings.Replace(line, "╮", "───╮", 1)
		}
		if strings.HasPrefix(line, "│ ") && strings.HasSuffix(line, " │") {
			// add the test case #
			testCounter++
			line = strings.Replace(line, "│", fmt.Sprintf("│ %d.", testCounter), 1)
		}
		if strings.HasSuffix(line, "╯") {
			// lengthen the line
			line = strings.Replace(line, "╯", "───╯", 1)
		}

		fmt.Println(line) // COMMENT THIS OUT WHEN USING DEBUGGING LINE AT TOP (after the for loop)

	}
}

func passThrough(bin string, args []string) {

	fmt.Printf("🐻 Official bootdev cli <run> command not detected.\n")
	fmt.Printf("🐻 Passing through to the official bootdev cli...\n\n")

	cmd := exec.Command(bin, args...)
	// Connect the subprocess directly to the user's terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
