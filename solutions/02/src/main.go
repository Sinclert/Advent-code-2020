package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type PasswordRule struct {
	minOccurrences int
	maxOccurrences int
	referencedChar string
}


func main() {

	rules, passwords, err := parsePasswordFile("solutions/02/files/pass.txt")
	if err != nil {
		fmt.Printf("Invalid password rules\n")
	}

	validPasswords := 0

	for i := range rules {
		occurrences := strings.Count(passwords[i], rules[i].referencedChar)
		if (rules[i].minOccurrences <= occurrences) &&
			(rules[i].maxOccurrences >= occurrences) {
			validPasswords += 1
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", validPasswords)
}


func parsePasswordFile(filePath string) ([]PasswordRule, []string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	var policies  []PasswordRule
	var passwords []string

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		rule, pass, err := parsePasswordInfo(scanner.Text())
		if err != nil {
			return nil, nil, err
		}
		policies  = append(policies, rule)
		passwords = append(passwords, pass)
	}

	return policies, passwords, nil
}


func parsePasswordInfo (passwordSpec string) (PasswordRule, string, error) {

	splitResults := strings.Split(passwordSpec, ":")
	if len(splitResults) != 2 {
		return PasswordRule{}, "", errors.New("invalid password specification")
	}

	rule := strings.TrimSpace(splitResults[0])
	pass := strings.TrimSpace(splitResults[1])

	splitResults = strings.Split(rule, " ")
	if len(splitResults) != 2 {
		return PasswordRule{}, "", errors.New("invalid password rule")
	}

	occurrences := strings.TrimSpace(splitResults[0])
	character   := strings.TrimSpace(splitResults[1])

	splitResults = strings.Split(occurrences, "-")
	if len(splitResults) != 2 {
		return PasswordRule{}, "", errors.New("invalid password rule occurrences")
	}

	minOccurrencesStr := strings.TrimSpace(splitResults[0])
	maxOccurrencesStr := strings.TrimSpace(splitResults[1])
	minOccurrencesNum, _ := strconv.ParseInt(minOccurrencesStr, 10, 32)
	maxOccurrencesNum, _ := strconv.ParseInt(maxOccurrencesStr, 10, 32)

	parsedRule := PasswordRule{
		referencedChar: character,
		minOccurrences: int(minOccurrencesNum),
		maxOccurrences: int(maxOccurrencesNum),
	}

	return parsedRule, pass, nil
}
