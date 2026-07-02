package main

import (
	"bufio"
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

	// Intercept as run if command is 'run' WITHOUT an '-s' or '--submit' flag
	if isRun && !isSubmit {
		fmt.Printf("🐻 Official 'bootdev run <id>' command detected.\n")
		fmt.Printf("🔎 Verify intercepting...\n\n")
		handleRun(realBootDevPath, args)
		return
	}

	// Otherwise, let the official CLI handle it
	fmt.Printf("🐻 Official bootdev <id> run command not detected.\n")
	fmt.Printf("🐻 Passing through to the official bootdev cli...\n\n")
	passThrough(realBootDevPath, args)
}

func handleRun(bin string, args []string) {

	cmd := exec.Command(bin, args...)
	cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1") // Force colors even when capturing output
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

	// After the loop finishes printing everything:
	fmt.Printf("\n🔎 Verify: Tests complete. Would you like to submit for grading? (y/any_other_key): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" || input == "yes" {
		fmt.Printf("🔎 Verify: Submitting to official bootdev cli 🐻...\n\n")

		// Add the -s flag to the current arguments
		submitArgs := append(args, "-s")

		// Use passThrough logic to ensure colors and interactivity work
		passThrough(bin, submitArgs)
	} else {
		fmt.Println("🔎 Verify: Skipping submission. Happy coding!")
	}

}

func passThrough(bin string, args []string) {

	cmd := exec.Command(bin, args...)
	// Connect the subprocess directly to the user's terminal
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
