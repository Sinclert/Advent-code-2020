package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)


var literalRuleRegex = regexp.MustCompile(`"\w+"`)
var nextRuleRegex = regexp.MustCompile(`\d+( \d+)*`)


type Rule struct {
	validRules [][]string
	validChars []string
}


func main() {

	rules, messages, err := parseMessageRulesFile("solutions/19/files/message_rules.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	allValidMessagesList := generateValidMessagesList(rules, "0")
	allValidMessagesMap  := buildValidMessagesMap(allValidMessagesList)
	totalValidMessages   := 0

	for _, message := range messages {
		_, ok := allValidMessagesMap[message]
		if ok {
			totalValidMessages += 1
			fmt.Printf("The message \"%s\" is: valid\n", message)
		} else {
			fmt.Printf("The message \"%s\" is: invalid\n", message)
		}
	}

	fmt.Printf("The number of valid messages is: %d\n", totalValidMessages)
}


func parseMessageRulesFile(filePath string) (map[string]Rule, []string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	scanner  := bufio.NewScanner(f)
	rules    := parseMessageRules(scanner)
	messages := parseMessages(scanner)

	return rules, messages, nil
}


func parseMessageRules(scanner *bufio.Scanner) map[string]Rule {

	var rules = make(map[string]Rule)

	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}

		var messageRule Rule
		var lineParts = strings.Split(line, ": ")
		var ruleNumber = lineParts[0]

		// Saving rule literal characters
		var ruleChars = literalRuleRegex.FindAllString(lineParts[1], -1)
		for i, _ := range ruleChars {
			ruleChars[i] = strings.Replace(ruleChars[i], "\"", "", -1)
		}

		// Saving rule following rules IDs
		var nextRulesList [][]string
		var nextRules = nextRuleRegex.FindAllString(lineParts[1], -1)
		for i, _ := range nextRules {
			nextRulesIds := strings.Split(nextRules[i], " ")
			nextRulesList = append(nextRulesList, nextRulesIds)
		}

		messageRule.validRules = nextRulesList
		messageRule.validChars = ruleChars
		rules[ruleNumber] = messageRule
	}

	return rules
}


func parseMessages(scanner *bufio.Scanner) []string {

	var messages []string

	for scanner.Scan() {
		line := scanner.Text()
		messages = append(messages, line)
	}

	return messages
}


func generateValidMessagesList(rules map[string]Rule, ruleId string) []string {

	var currentRule = rules[ruleId]

	if len(currentRule.validChars) > 0 {
		return currentRule.validChars
	}

	var allMessages []string
	for _, ruleSet := range currentRule.validRules {

		var setMessages []string
		for _, ruleId := range ruleSet {

			var ruleMessages = generateValidMessagesList(rules, ruleId)

			// First rule of the set
			if setMessages == nil {
				setMessages = ruleMessages
				continue
			}

			// Combine previous set rule messages with the next rule ones
			setMessages = combineListsOfMessages(setMessages, ruleMessages)
		}

		// Move all set possible messages to the overall list
		for _, setMsg := range setMessages {
			allMessages = append(allMessages, setMsg)
		}
	}

	return allMessages
}


func combineListsOfMessages(listA []string, listB []string) []string {

	var combinedList []string

	for _, messageA := range listA {
		for _, messageB := range listB {
			combineMsg  := messageA + messageB
			combinedList = append(combinedList, combineMsg)
		}
	}

	return combinedList
}


func buildValidMessagesMap(allValidMessages []string) map[string]bool {

	var validMessagesMap = make(map[string]bool)

	for _, message := range allValidMessages {
		validMessagesMap[message] = true
	}

	return validMessagesMap
}
