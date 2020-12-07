package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var Questions = [...]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}


func main() {

	formResponses, err := parseResponsesFile("solutions/06/files/form_responses.txt")
	if err != nil {
		fmt.Println(err)
	}

	totalUniqueResponses := 0

	for i, groupResponses := range *formResponses {
		groupUniqueResponses := checkGroupResponses(groupResponses)
		totalUniqueResponses += groupUniqueResponses
		fmt.Printf("Unique 'YES' answers in group %d is: %d\n", i, groupUniqueResponses)
	}

	fmt.Printf("Unique 'YES' answers within groups is: %d\n", totalUniqueResponses)
}


func parseResponsesFile(filePath string) (*[][]string, error) {

	var allResponses [][]string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	var currentGroup []string

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			allResponses = append(allResponses, currentGroup)
			currentGroup = nil
		} else {
			currentGroup = append(currentGroup, line)
		}
	}

	allResponses = append(allResponses, currentGroup)
	return &allResponses, nil
}


func checkGroupResponses(groupResponses []string) int {

	groupUniqueResponsesMap := buildGroupResponsesMap()
	groupUniqueResponsesNum := 0

	for _, userResponses := range groupResponses {
		groupUniqueResponsesNum += checkUserResponses(userResponses, &groupUniqueResponsesMap)
	}

	return groupUniqueResponsesNum
}


func checkUserResponses(userResponse string, responsesMap *map[string]bool) int {

	userUniqueResponseList := strings.Split(userResponse, "")
	userUniqueResponseNum  := 0

	for _, response := range userUniqueResponseList {
		if (*responsesMap)[response] == false {
			(*responsesMap)[response] = true
			userUniqueResponseNum += 1
		}
	}

	return userUniqueResponseNum
}


func buildGroupResponsesMap() map[string]bool {

	yesResponses := make(map[string]bool)
	for _, question := range Questions {
		yesResponses[question] = false
	}

	return yesResponses
}
