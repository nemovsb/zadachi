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

// A = 65
// B = 66
// C = 67
// S = 83

var event1 = [3]rune{65, 66, 67}
var event2 = [3]rune{65, 66, 83}
var event3 = [2]rune{65, 67}

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

	eLog := make([]logEvent, 0)

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

		eLog = append(eLog, logEvent{
			day:    d,
			hour:   h,
			minute: m,
			id:     id,
			event:  rune(strArr[4][0]),
		})

	}

	sort.Slice(eLog, func(i, j int) bool {
		return eLog[i].id < eLog[j].id
	})

	ids := make([]int, 0)
	for _, v := range eLog {

		ok := true
		for _, g := range ids {

			ok = !(v.id == g)

			if !ok {
				break
			}

		}

		if ok {
			ids = append(ids, v.id)
		}

	}

	var result string

	for _, id := range ids {
		var timeRes int
		log := getLogById(id, eLog)

		sortLogByTime(log)

		eventNum := 0
		var time int

		for i, event := range log {
			res, done := checkEvent(event, eventNum)
			if res {
				if eventNum != 0 {
					time += getMinute(event) - getMinute(log[i-1])
				}
				if done {
					timeRes += time
					eventNum = 0
				}
				eventNum++
			} else {
				eventNum = 0
				res, _ := checkEvent(event, eventNum)
				if res {
					eventNum++
					time = 0
				}
			}
		}
		result = fmt.Sprint(result, " ", strconv.Itoa(timeRes))
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

	output.WriteString(result)

}

func checkEvent(e logEvent, i int) (res, done bool) {

	if i < len(event3) {
		if (i >= len(event1)) && (i >= len(event2)) {
			done = true
		}

		res = e.event == event1[i] || e.event == event2[i] || e.event == event3[i]

		if (e.event == event3[i]) && (len(event3)-1 == i) {
			done = true
		}

		return res, done

	} else {
		if i >= len(event3) {
			done = true
		}
		res = e.event == event1[i] || e.event == event2[i]
		return res, done
	}

}

func getLogById(id int, l []logEvent) []logEvent {
	res := make([]logEvent, 0)
	for _, v := range l {
		if id == v.id {
			res = append(res, v)
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
