package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	var result string

	for i := 1; i <= num; i++ {
		fileScanner.Scan()
		strArr := strings.Split(fileScanner.Text(), ",")

		surname := strArr[0]
		name := strArr[1]
		patronymic := strArr[2]
		day, err := strconv.Atoi(strArr[3])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		month, err := strconv.Atoi(strArr[4])
		if err != nil {
			log.Fatal("Error: ", err)
		}
		year, err := strconv.Atoi(strArr[5])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		result = fmt.Sprintf("%s %s", result, encrypt(surname, name, patronymic, day, month, year))

	}

	//fmt.Printf("result = %s\n", result)

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

func encrypt(surname, name, patronymic string, day, month, year int) string {

	gems := make([]rune, 0)
	for _, v := range []rune(fmt.Sprint(surname, name, patronymic)) {

		ok := true
		for _, g := range gems {

			ok = !(v == g)

			if !ok {
				break
			}

		}

		if ok {
			gems = append(gems, v)
		}

	}

	s1 := len(gems)
	//fmt.Printf("s1: %d\n", s1)

	s2 := (dgsum(day) + dgsum(month)) * 64
	//fmt.Printf("num character : %d\n", (int([]rune(strings.ToLower(surname))[0]) - 96))

	//fmt.Printf("s2: %d\n", s2)
	s3 := (int([]rune(strings.ToLower(surname))[0]) - 96) * 256
	//fmt.Printf("s3: %d\n", s3)

	res := s1 + s2 + s3
	//fmt.Printf("res: %d\n", res)
	v := []rune(strings.ToUpper(strconv.FormatInt(int64(res), 16)))

	cipher := string(v[len(v)-3:])
	//fmt.Printf("cipher: %s\n", cipher)

	return cipher

}
func dgsum(n int) int {
	var r = 0
	for ; n > 0; n /= 10 {
		r += n % 10
	}
	return r
}
