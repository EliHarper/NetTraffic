package main

import (
	"fmt"
	"os"
)

func main() {
	cliArg := os.Args[:]
	var malIpSlice []string
	var logFile []string
	var verifiedMalIps []string

	// Check if log file is provided
	if isLogProvided(len(cliArg)) {
		// If it is, use that file path to access the log
		// files cataloging malicious IP addresses
		if len(cliArg) == 2 {
			logFile = extractIPs(cliArg[1])
			fmt.Println("Log file path you provided:", cliArg[1])
			fmt.Println()
		} else {
			fmt.Println("Please provide only 1 argument to be used as your log file.")
			fmt.Println()
		}
	} else {
		// If logFile not provided, collect logs in [hardcoded as of rn] new file
		collectLogs()
		logFile = extractIPs("../logs/lsofLog.txt")
	}
	malIpSlice = append(malIpSlice, extractIPs("../blacklists/myIPLatestBL.txt")...)

	// Nested for loop to compare each IP to each item in malIP.txt
	for _, ip := range logFile {
		for _, malIp := range malIpSlice {
			if compare(ip, malIp) {
				verifiedMalIps = append(verifiedMalIps, ip)
			}
		}
	}

	printAnswer(verifiedMalIps)
	ticker()
}

func isLogProvided(argLength int) bool {
	if argLength < 2 {
		return false
	}
	return true
}

func printAnswer(verifiedMalIps []string) {
	if len(verifiedMalIps) == 0 {
		fmt.Println("None of the traffic checked is being sent to sources that are currently blacklisted.\n\n")
		return
	}
	if len(verifiedMalIps) == 1 {
		fmt.Println("This IP, which your device is communicating with, is blacklisted:\n")
		fmt.Println(verifiedMalIps[0])
		return
	}
	fmt.Printf("These %v IP addresses that your device is communicating with are blacklisted:\n\n", len(verifiedMalIps))

	for i := range verifiedMalIps {
		fmt.Println(verifiedMalIps[i])
	}
}
