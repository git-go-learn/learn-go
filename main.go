package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter commit message: ")
	scanner.Scan()
	commitMessage := scanner.Text()

	fmt.Print("Enter branch name to push (e.g., main): ")
	scanner.Scan()
	branchName := scanner.Text()

	// Run git add .
	exec.Command("git", "add", ".").CombinedOutput()

	// Run git commit -m "commitMessage"
	exec.Command("git", "commit", "-m", commitMessage).CombinedOutput()

	// Run git push origin branchName
	exec.Command("git", "push", "origin", branchName).CombinedOutput()
	
	fmt.Println("Commands executed.")
}
