package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)


type Food struct {
	ingredients         map[string]bool
	allergens           map[string]bool
}


var ingredientsRegex = regexp.MustCompile(`^\w+( \w+)*`)
var allergensRegex   = regexp.MustCompile(`\(contains (\w+(, \w+)*)\)`)


/*
INITIAL CLARIFICATION

This problem is not well explained in the Advent of Code website (https://adventofcode.com/2020/day/21)
Therefore, I am stating here a couple of insights to simplify what is required to solve the problem.

Insight 1:
Only one ingredient from the overall list of ingredients can contain each allergen.

Insight 2:
If an ingredient contains any of the three known allergens, it will show in the food allergens list.

Real example:
Considering these two insights, we could look at the food specification in the reverse direction (right to left),
determining that for a range of food to contain the allergen "Dairy", there must be a common ingredient in both lists.
Clearly, the candidate ingredient for containing "Dairy", is "mxmxvkd"
*/


func main() {

	foods, allIngredientsSet, allAllergensSet, err := parseFoodListsFile("solutions/21/files/food_list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var allergenCandidates = computeAllergenCandidates(foods, allAllergensSet)

	// Generate set of confirmed candidates
	var inIngredientsSet = make(map[string]bool)
	for _, candidates := range allergenCandidates {
		for candidate := range candidates {
			inIngredientsSet[candidate] = true
		}
	}

	var allIngredientsArr = buildArrayFromSet(allIngredientsSet)
	var inIngredientsArr  = buildArrayFromSet(inIngredientsSet)

	var outIngredientsSet = filterOutSet(allIngredientsSet, inIngredientsArr)
	var outIngredientsArr = buildArrayFromSet(outIngredientsSet)
	fmt.Printf("The ingredients that cannot contain any allergen are: %s\n", outIngredientsArr)

	var allCounters    = countIngredientsAppearances(foods, allIngredientsArr)
	var outCounters    = filterOutMap(allCounters, inIngredientsArr)
	var sumAppearances = sumCounters(outCounters)
	fmt.Printf("Total number of appearances: %d\n", sumAppearances)
}


func parseFoodListsFile(filePath string) ([]Food, map[string]bool, map[string]bool, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, err
	}

	defer f.Close()

	var ingredientsSet = make(map[string]bool)
	var allergensSet   = make(map[string]bool)
	var foods          []Food

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		ingredientsStr := ingredientsRegex.FindString(line)
		ingredientsArr := strings.Split(ingredientsStr, " ")
		ingredientsSet  = extendSet(ingredientsSet, ingredientsArr)

		allergensStr   := allergensRegex.FindStringSubmatch(line)
		allergensArr   := strings.Split(allergensStr[1], ", ")
		allergensSet  = extendSet(allergensSet, allergensArr)

		foods = append(foods, Food{
			ingredients: buildSetFromArray(ingredientsArr),
			allergens:   buildSetFromArray(allergensArr),
		})

	}

	return foods, ingredientsSet, allergensSet, nil
}


func computeAllergenCandidates(foods []Food, allAllergens map[string]bool) map[string]map[string]bool {

	var allergenCandidates = make(map[string]map[string]bool)

	for allergen, _  := range allAllergens {
		for _, food  := range foods {
			_, found := food.allergens[allergen]

			if found == false {
				continue
			}

			if allergenCandidates[allergen] == nil {
				allergenCandidates[allergen] = food.ingredients
				continue
			}

			allergenCandidates[allergen] = buildSetIntersection(allergenCandidates[allergen], food.ingredients)
		}
	}

	return allergenCandidates
}


func filterOutSet(set map[string]bool, arr []string) map[string]bool {

	for _, key := range arr {
		delete(set, key)
	}

	return set
}


func filterOutMap(set map[string]int, arr []string) map[string]int {

	for _, key := range arr {
		delete(set, key)
	}

	return set
}


func extendSet(set map[string]bool, arr []string) map[string]bool {

	for _, element := range arr {
		set[element] = true
	}

	return set
}


func buildArrayFromSet(set map[string]bool) []string {

	var newArray []string
	for key, _  := range set {
		newArray = append(newArray, key)
	}

	return newArray
}


func buildSetFromArray(arr []string) map[string]bool {

	var filledSet = make(map[string]bool)
	for _, element := range arr {
		filledSet[element] = true
	}

	return filledSet
}


func buildSetIntersection(mapA map[string]bool, mapB map[string]bool) map[string]bool {

	var jointSet  = make(map[string]bool)
	for key, val := range mapA {
		_, ok := mapB[key]
		if ok {
			jointSet[key] = val
		}
	}

	return jointSet
}


func countIngredientsAppearances(foods []Food, keys []string) map[string]int {

	var counters = make(map[string]int)
	for _, key  := range keys {
		counters[key] = 0
	}

	for _, food  := range foods {
		for ingr := range food.ingredients {
			counters[ingr] += 1
		}
	}

	return counters
}


func sumCounters(counters map[string]int) int {

	var totalSum = 0
	for _, num   := range counters {
		totalSum += num
	}

	return totalSum
}
