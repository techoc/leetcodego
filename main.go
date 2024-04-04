package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// exec 执行所有的测试用例
	cmd := exec.Command("go", "test", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		return
	}
	fmt.Println(string(output))
}
