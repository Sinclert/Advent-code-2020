package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)


func main() {

	adaptersJolts, err := parseAdaptersFile("solutions/10/files/adapters_jolts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	adaptersChain, err := buildAdaptersChain(adaptersJolts, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	lastAdapterIndex := len(adaptersChain)-1
	lastAdapterValue := adaptersChain[lastAdapterIndex]
	userAdapterJolts := lastAdapterValue + 3

	fmt.Printf("The user device joltage is %d\n", userAdapterJolts)

	adaptersChain  = append(adaptersChain, userAdapterJolts)
	adaptersDiffs := computeAdapterDiffs(adaptersChain, 3)

	for diff, counter := range adaptersDiffs {
		fmt.Printf("Jolt differences of %d: %d\n", diff, counter)
	}
}


func parseAdaptersFile(filePath string) ([]int, error) {

	var adaptersJolts []int

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		jolts, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}
		adaptersJolts = append(adaptersJolts, int(jolts))
	}

	return adaptersJolts, nil
}


func buildAdaptersChain(adapterJolts []int, adapterJoltsRange int) ([]int, error) {

	// Initial adapter (implicit)
	var previousAdapter = 0
	var adapterJoltsSet = buildAdaptersSet(adapterJolts)
	var adaptersChain   = []int{previousAdapter}

	for true {
		foundNext := false

		for i := 1; i <= adapterJoltsRange; i++ {
			candidateAdapter := previousAdapter + i
			_, ok := adapterJoltsSet[candidateAdapter]
			if ok {
				previousAdapter = candidateAdapter
				adaptersChain = append(adaptersChain, candidateAdapter)
				foundNext = true
				break
			}
		}

		if foundNext == false {
			break
		}
	}

	// The (implicit) wall adapter is present in one of the lists
	if len(adapterJolts) != len(adaptersChain)-1 {
		return nil, errors.New("he chain of adapters could not include all of them")
	}

	return adaptersChain, nil
}


func buildAdaptersSet(adaptersJolts []int) map[int]bool {

	adaptersSet := make(map[int]bool)
	for _, jolts := range adaptersJolts {
		adaptersSet[jolts] = true
	}

	return adaptersSet
}


func computeAdapterDiffs(chainOfAdapters []int, adapterJoltsRange int) map[int]int {

	adaptersDiffs := make(map[int]int)

	// Initialize difference counters
	for i := 1; i <= adapterJoltsRange; i++ {
		adaptersDiffs[i] = 0
	}

	// Compute adapters differences
	for i := 0; i < len(chainOfAdapters)-1; i++ {
		currentAdapter := chainOfAdapters[i]
		nextAdapter    := chainOfAdapters[i+1]
		diff           := nextAdapter - currentAdapter
		adaptersDiffs[diff] += 1
	}

	return adaptersDiffs
}
