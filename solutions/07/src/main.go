package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


var BagColors = [...]string{
	"light-red",
	"dark-orange",
	"bright-white",
	"muted-yellow",
	"shiny-gold",
	"dark-olive",
	"vibrant-plum",
	"faded-blue",
	"dotted-black",
}


func main() {

	bagRules, err := parseBagRulesFile("solutions/07/files/bag_rules.json")
	if err != nil {
		fmt.Println(err)
	}

	// Serves as a cache for previously checked bags
	wrappingBags := buildWrappingBagMap()

	for bagName, bagRequirements := range bagRules {
		if wrappingBags[bagName] == false {
			wrappingBags[bagName] = recursiveSearch(bagRules, bagRequirements, "shiny-gold")
		}
	}

	wrappingBagsNum := calcWrappingBagNum(wrappingBags)
	fmt.Printf("Number of wrapping bags: %d\n", wrappingBagsNum)
}


func parseBagRulesFile(filePath string) (map[string]map[string]int, error) {

	var bagRules map[string]map[string]int

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &bagRules)
	if err != nil {
		return nil, err
	}

	return bagRules, nil
}


func buildWrappingBagMap() map[string]bool {

	wrappingBags := make(map[string]bool)
	for _, color := range BagColors {
		wrappingBags[color] = false
	}

	return wrappingBags
}


func recursiveSearch(bagRules map[string]map[string]int, requirements map[string]int, desiredBag string) bool {

	if requirements[desiredBag] > 0 {
		return true
	}

	for bagName, bagNum := range requirements {
		if bagNum > 0 {
			return recursiveSearch(bagRules, bagRules[bagName], desiredBag)
		}
	}

	return false
}


func calcWrappingBagNum(wrappingBags map[string]bool) int {

	wrappingBagsNum := 0
	for bagName, isWrapping := range wrappingBags {
		if isWrapping {
			fmt.Printf("The bag '%s' contains the desired bag\n", bagName)
			wrappingBagsNum += 1
		}
	}

	return wrappingBagsNum
}
