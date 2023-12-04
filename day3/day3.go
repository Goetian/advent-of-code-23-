package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Settings struct {
	max int
}

func newSettings(maxRow int) *Settings {
	return &Settings{
		max: maxRow,
	}
}

var (
	S = newSettings(200)
)

type Tracker struct {
	tracker map[int][]int
}

func newTracker() *Tracker {
	return &Tracker{
		tracker: make(map[int][]int),
	}
}

func Day3() {

	// 1 2 3 4 5 6 7 8 9 10 11
	// . . . . . # . . . . .
	// . . : . . . . . . . .
	// ...222..!..

	file, _ := os.Open("./day3/input.txt")

	reader := bufio.NewReader(file)
	i := 1
	t := newTracker()
	t.tracker[i] = []int{}
	summe := 0
	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		slice := getSlice(string(line))
		symbole := findSymbole(slice)
		addToTracker(t, i, symbole)
		i++
	}
	i = 1
	file2, _ := os.Open("./day3/input.txt")
	reader = bufio.NewReader(file2)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		slice := getSlice(string(line))
		adjNum, postionNumber := getNumberMatrix(slice)
		all_pos := getAllPostions(adjNum, postionNumber)
		summe += getValidValues(t, all_pos, adjNum, i)
		i++
	}
	fmt.Println(summe)
}

func getValidValues(t *Tracker, pos [][]int, numbers []int, row int) int {
	ret := 0
	var wasAdded []int
	for i, x := range numbers {
		postions := pos[i]
		good_pos := t.tracker[row]
		for _, p := range postions {
			if slices.Contains(good_pos, p) {
				if !slices.Contains(wasAdded, x) {
					wasAdded = append(wasAdded, x)
				}
			}

		}

	}

	for _, val := range wasAdded {
		ret += val
	}
	return ret
}

func getAllPostions(numb []int, pos []int) [][]int {
	allPostions := [][]int{}
	toAdd := []int{}
	for i := 0; i < len(numb); i++ {
		digitsOfNumb := digits(numb[i])

		for n := digitsOfNumb - 1; n > -1; n-- {
			toAdd = append(toAdd, pos[i]-n)
		}
		allPostions = append(allPostions, toAdd)
		toAdd = []int{}
	}
	return allPostions
}

func digits(i int) int {
	ret := 0
	strI := strconv.Itoa(i)
	ret = len(strI)
	return ret
}

func findSymbole(s []string) []int {

	ret := []int{}
	for i, v := range s {
		if v == "." {
			continue
		}
		_, ok := strconv.Atoi(v)
		if ok != nil {
			ret = append(ret, i+1)
		}

	}
	return ret
}

func getSlice(s string) []string {

	sliceRunes := strings.Split(s, "")
	return sliceRunes
}

func addToTracker(m *Tracker, row int, v []int) {

	for _, number := range v {
		pre := number - 1
		post := number + 1
		if slices.Contains(m.tracker[row], pre) {

		} else {
			m.tracker[row] = append(m.tracker[row], pre)
		}
		if slices.Contains(m.tracker[row], post) {
		} else {
			m.tracker[row] = append(m.tracker[row], post)
		}
	}
	addToTrackeNextLine(m, row+1, v)
	updatedTrackerPreLine(m, row-1, v)
}

func addToTrackeNextLine(m *Tracker, row int, v []int) {
	if row == S.max {
		return
	}
	for _, number := range v {
		pre := number - 1
		post := number + 1
		m.tracker[row] = append(m.tracker[row], pre)
		m.tracker[row] = append(m.tracker[row], post)
		m.tracker[row] = append(m.tracker[row], number)
	}
}

func updatedTrackerPreLine(m *Tracker, row int, v []int) {
	if row == 0 {
		return
	}
	for _, number := range v {
		pre := number - 1
		post := number + 1
		if slices.Contains(m.tracker[row], pre) {
		} else {
			m.tracker[row] = append(m.tracker[row], pre)
		}
		if slices.Contains(m.tracker[row], post) {
		} else {
			m.tracker[row] = append(m.tracker[row], post)
		}
		if slices.Contains(m.tracker[row], number) {
		} else {
			m.tracker[row] = append(m.tracker[row], number)
		}
	}

}

func getNumberMatrix(v []string) ([]int, []int) {
	numbersAdj := []int{}
	temp := []int{}
	lastNum := []int{}
	for column, value := range v {
		isNum, intV := isNumber(value)
		if !isNum {
			if len(temp) > 0 {
				numbersAdj = append(numbersAdj, getIntegerFromSlice(temp))
				if column > 1 {
					lastNum = append(lastNum, column)
				} else {
					lastNum = append(lastNum, column)
				}
			}
			temp = []int{}
		} else {
			temp = append(temp, intV)
		}
	}
	for _, x := range numbersAdj {
		if x < 0 {
			fmt.Println(x)
		}
	}
	return numbersAdj, lastNum
}

func isNumber(s string) (bool, int) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return false, 0
	} else {
		return true, v
	}
}

func getIntegerFromSlice(in []int) int {
	temp := ""
	ret := 0
	for _, x := range in {
		temp += strconv.Itoa(x)
	}

	ret, _ = strconv.Atoi(temp)
	return ret
}
