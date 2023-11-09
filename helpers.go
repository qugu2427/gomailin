package main

// package mailin

import "strings"

func findCmd(msg string) (cmd string) {
	for _, char := range msg {
		if char == ' ' {
			break
		} else {
			cmd += string(char)
		}
	}
	return strings.ToUpper(cmd)
}

func findEmailInLine(line string) (isFound bool, email string) {
	lessThanIndex := -1
	greaterThanIndex := -1

	for i, char := range line {
		if char == '>' {
			greaterThanIndex = i
		} else if char == '<' && lessThanIndex == -1 {
			lessThanIndex = i
		}
	}

	if lessThanIndex == -1 ||
		greaterThanIndex == -1 ||
		lessThanIndex >= greaterThanIndex {
		return false, ""
	} else {
		return true, line[lessThanIndex+1 : greaterThanIndex]
	}

}
