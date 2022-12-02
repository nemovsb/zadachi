package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type logEvent struct {
	day    int
	hour   int
	minute int
	id     int
	event  rune
}

type TravelTime struct {
	id   int
	time int
}

// A = 65
// B = 66
// C = 67
// S = 83

var patterns = [][]rune{
	{65, 66, 67},
	{65, 66, 83},
	{65, 67},
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Scan()
	strCount := fileScanner.Text()

	num, err := strconv.Atoi(strCount)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	FullLog := make([]logEvent, 0)
	ids := make([]int, 0)

	for i := 1; i <= num; i++ {
		fileScanner.Scan()
		strArr := strings.Split(fileScanner.Text(), " ")

		d, err := strconv.Atoi(strArr[0])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		h, err := strconv.Atoi(strArr[1])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		m, err := strconv.Atoi(strArr[2])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		id, err := strconv.Atoi(strArr[3])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		{
			ok := true
			for _, g := range ids {

				ok = !(id == g)

				if !ok {
					break
				}

			}
			if ok {
				ids = append(ids, id)
			}
		}

		FullLog = append(FullLog, logEvent{
			day:    d,
			hour:   h,
			minute: m,
			id:     id,
			event:  rune(strArr[4][0]),
		})

	}

	var result string

	timeResult := make(chan (TravelTime))

	for _, id := range ids {

		go func(id int) {
			var timeEvent int

			log := getLogById(id, FullLog)
			sortLogByTime(log)

			//for i := 0; i < len(log); i++ {
			for i := range log {

				for _, pattern := range patterns {
					for j, event := range pattern {
						if event != log[i+j].event {
							break
						}
						if j == len(pattern)-1 {
							timeEvent += getMinute(log[i+j]) - getMinute(log[i])
							i += j
						}
					}
				}

			}
			timeResult <- TravelTime{
				id:   id,
				time: timeEvent,
			}
		}(id)
	}

	resultTravelTime := make([]TravelTime, 0)

	for range ids {
		resultTravelTime = append(resultTravelTime, <-timeResult)
	}

	sort.Slice(resultTravelTime, func(i, j int) bool {
		return resultTravelTime[i].id < resultTravelTime[j].id
	})

	for _, res := range resultTravelTime {

		result = fmt.Sprint(result, " ", strconv.Itoa(res.time))
	}

	output, err := os.OpenFile(`output.txt`, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer output.Close()

	err = output.Truncate(0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	output.WriteString(strings.Trim(result, " "))

}

func getLogById(id int, l []logEvent) []logEvent {
	res := make([]logEvent, 0)
	for i := 0; i < len(l); i++ {
		if id == l[i].id {
			res = append(res, l[i])
		}
	}
	return res
}

func getMinute(l logEvent) int {
	return l.minute + l.hour*60 + l.day*1440
}

func sortLogByTime(l []logEvent) {
	sort.Slice(l, func(i, j int) bool {
		return getMinute(l[i]) < getMinute(l[j])
	})
}
