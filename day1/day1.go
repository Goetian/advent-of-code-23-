package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	ret := []string{}
	configValue := []string{}
	tempV := []int{}

	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		values := string(line)
		tempV = tempV[:0]
		for _, value := range values {

			if unicode.IsNumber(rune(value)) {
				v, _ := strconv.Atoi(string(value))
				tempV = append(tempV, v)

			}
		}

		configValue = append(configValue, fmt.Sprint(tempV[0]))
		configValue = append(configValue, fmt.Sprint(tempV[len(tempV)-1]))
		ret = append(ret, strings.Join(configValue, ""))
		configValue = configValue[:0]

	}
	returnValue := 0
	for _, v := range ret {
		temp, _ := strconv.Atoi(v)
		returnValue += temp
	}
	fmt.Println(returnValue)
}

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
