package main

import "strings"

func compare(logIP string, malIP string) bool {
	if strings.Contains(logIP, malIP) {
		return true
	}
	return false
}
