package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func main() {

	leaveTimestamp, busSchedules, err := parseBusSchedulesFile("solutions/13/files/bus_schedules.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bestSchedule := 0
	smallestDiff := math.Inf(1)

	for _, sched := range busSchedules {
		timeDiff := float64(leaveTimestamp % sched)
		timeDiff  = float64(sched) - timeDiff
		if timeDiff < smallestDiff {
			bestSchedule = sched
			smallestDiff = timeDiff
		}
	}

	waitingMins := int(smallestDiff)
	fmt.Printf("The best bus to take has the ID: %d\n", bestSchedule)
	fmt.Printf("The user needs to wait for %d minutes\n", waitingMins)
	fmt.Printf("Multiplying the bus ID by the waiting minutes: %d\n", bestSchedule * waitingMins)
}


func parseBusSchedulesFile(filePath string) (int, []int, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return 0, nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	leaveTimestamp, leaveErr := parseLeaveTimestamp(scanner)
	busSchedules, busesErr   := parseBusSchedules(scanner)

	if (leaveErr != nil) || (busesErr != nil) {
		return 0, nil, err
	}

	return leaveTimestamp, busSchedules, nil
}


func parseLeaveTimestamp(scanner *bufio.Scanner) (int, error) {

	scanner.Scan()

	timestampStr := scanner.Text()
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(timestampInt), nil
}


func parseBusSchedules(scanner *bufio.Scanner) ([]int, error) {

	var busSchedules []int

	scanner.Scan()

	schedulesStr := scanner.Text()
	schedulesArr := strings.Split(schedulesStr, ",")

	for _, sched := range schedulesArr {

		// Skip out of service buses
		if sched == "x" {
			continue
		}

		s, err := strconv.ParseInt(sched, 10, 32)
		if err != nil {
			return nil, err
		}
		busSchedules = append(busSchedules, int(s))
	}

	return busSchedules, nil
}
