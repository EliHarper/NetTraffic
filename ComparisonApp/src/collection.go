package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func collectLogs() {
	cmd := exec.Command("", "")
	if _, err := os.Stat("../logs/lsofLog.txt"); os.IsNotExist(err) {
		cmd = exec.Command("bash", "../scripts/no_prev_log.sh")
	} else {
		fmt.Println("I see you've been here before.. Thanks for using my app!")
		fmt.Println()
		cmd = exec.Command("bash", "../scripts/has_prev_log.sh")
	}
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished logging your network traffic!")
	fmt.Println()
	if strings.Contains(cmd.Args[1], "no_prev_log") {
		fmt.Println("New log file is named lsofLog.txt")
		fmt.Println()
	} else {
		fmt.Println("Appended to lsofLog.txt")
		fmt.Println()
	}
	fmt.Println("Now comparing logs to blacklist...")
	fmt.Println()
}
