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

type RocketsLog map[int][]logEvent

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

	rocketsLog := make(RocketsLog, 0)

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

		rocketsLog[id] = append(rocketsLog[id], logEvent{
			day:    d,
			hour:   h,
			minute: m,
			id:     id,
			event:  rune(strArr[4][0]),
		})

	}

	var result string

	timeResult := make(chan (TravelTime))

	// Цикл по ракетам
	for id, rocketLog := range rocketsLog {

		//Для лога каждой ракеты свой обработчик
		go func(id int, log []logEvent) {
			var timeEvent int

			sortLogByTime(log)

			//Перебор событий ракеты
			for i := range log {

				//Перебор возможных сценариев
				for _, pattern := range patterns {

					hit := false

					//Перебор событий сценария
					for j, event := range pattern {

						//Если событие сценария не совпадает с текущим, то переход к следующему сценарию
						if event != log[i+j].event {
							break
						}

						//Если последнее событие сценария, считаем время и наращиваем счетчик записей лога ракеты (i+j)
						if j == len(pattern)-1 {
							hit = true
							timeEvent += getMinute(log[i+j]) - getMinute(log[i])
							i += j
						}
					}

					//Если паттерн совпал, переходим к следующим записям лога
					if hit {
						break
					}
				}

			}
			//Отправляем результат
			timeResult <- TravelTime{
				id:   id,
				time: timeEvent,
			}
		}(id, rocketLog)
	}

	resultTravelTime := make([]TravelTime, 0)

	for range rocketsLog {
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
