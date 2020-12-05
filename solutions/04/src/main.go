package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


var MandatoryFields = [...]string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
var OptionalFields  = [...]string{"cid"}


func main() {

	specifications, err := parsePassportsFile("solutions/04/files/passports.txt")
	if err != nil {
		fmt.Print(err)
	}

	validPassports := 0

	for _, spec := range *specifications {
		specMap, err := buildSpecMap(spec)
		if err != nil {
			fmt.Print(err)
			return
		}
		if isValidPassport(specMap) {
			validPassports += 1
		}
	}

	fmt.Printf("Number of valid passports: %d\n", validPassports)
}


func parsePassportsFile(filePath string) (*[]string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var allSpecs []string
	var currentSpec strings.Builder

	var fileScanner = bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			currentSpec.WriteString(line)
			currentSpec.WriteString(" ")
		} else {
			strSpec := builderToString(&currentSpec)
			allSpecs = append(allSpecs, strSpec)
		}
	}

	strSpec := builderToString(&currentSpec)
	allSpecs = append(allSpecs, strSpec)

	return &allSpecs, nil
}


func builderToString(builder *strings.Builder) string {
	defer builder.Reset()
	return strings.TrimSpace(builder.String())
}


func buildSpecMap(specification string) (*map[string]string, error) {

	fieldsStr := strings.Split(specification, " ")
	fieldsMap := make(map[string]string)

	for _, field := range fieldsStr {
		keyVal := strings.Split(field, ":")
		if len(keyVal) != 2 {
			return nil, errors.New("invalid passport format")
		}
		fieldsMap[keyVal[0]] = keyVal[1]
	}

	return &fieldsMap, nil
}


func isValidPassport(specification *map[string]string) bool {

	for _, field := range MandatoryFields {
		_, ok := (*specification)[field]
		if ok == false {
			return false
		}
	}

	return true
}
