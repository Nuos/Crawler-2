package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrorStringSpace = errors.New("string trim error")
)

func HandleCommon(old string) (string, error) {
	newReplacer := strings.NewReplacer("\n", "", "\t", "")
	return strings.TrimSpace(newReplacer.Replace(old)), nil
}
func HandleUrl(old string) (string, error) {
	return "https://github.com" + strings.TrimSpace(old), nil
}
func StringTrim(old string) string {
	return strings.TrimSpace(old)
}

func StringSplit(old string) string {
	if old == "None" {
		return old
	}
	var (
		stringList []string
		length     int
	)
	stringList = strings.Split(old, "/")
	length = len(stringList)
	return stringList[length-1]
}
func StringsReplace(old string) string {
	var newReplacer *strings.Replacer

	newReplacer = strings.NewReplacer("\t", "")
	return newReplacer.Replace(old)
}
func StringSplitByDot(old string) (int, int, int) {
	var stringList []string
	stringList = strings.Split(old, "•")

	var numberList []int
	for index, oneString := range stringList {
		if index > 0 && index < 4 {
			number := StringRegexp(oneString)
			numberList = append(numberList, number)
		}
	}
	return numberList[0], numberList[1], numberList[2]
}
func StringRegexp(old string) int {
	reg := regexp.MustCompile(`\d+`)
	stringContent := reg.FindString(old)
	number, _ := strconv.Atoi(stringContent)
	return number
}
