package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func extractIPs(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log the error and entirely quit the program
		fmt.Println("Error >> ", err)
		os.Exit(1)
	}

	var loggedIPs []string
	regexPattern := `[0-9][0-9]?[0-9]?[.][0-9][0-9]?[0-9]?[.][0-9][0-9]?[0-9]?[.][0-9][0-9]?[0-9]?`
	s := strings.Fields(string(bs))
	for _, item := range s {
		matched, _ := regexp.MatchString(regexPattern, item)
		if matched {
			if strings.Contains(item, "->") {
				localAndRemote := strings.Split(item, "->")
				// Only append the remote TCP socket
				loggedIPs = append(loggedIPs, localAndRemote[1])
			} else if strings.Contains(item, ">") {
				postGt := strings.Split(item, ">")
				loggedIPs = append(loggedIPs, postGt[1])
			} else {
				loggedIPs = append(loggedIPs, item)
			}
		}
	}

	return loggedIPs
}
