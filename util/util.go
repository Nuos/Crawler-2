package util

import (
	"errors"
	"strings"
)

var (
	//ErrorStringSpace e
	ErrorStringSpace = errors.New("string trim error")
)

//HandleCommon a
func HandleCommon(oldString string) (string, error) {
	newReplacer := strings.NewReplacer("\n", "", "\t", "")
	return strings.TrimSpace(newReplacer.Replace(oldString)), nil
}

//HandleUrl a
func HandleUrl(oldString string) (string, error) {
	return "https://github.com" + strings.TrimSpace(oldString), nil
}
