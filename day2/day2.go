package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	line map[string]int
}

func newGame() *Game {
	return &Game{
		line: make(map[string]int),
	}

}

func (g *Game) printGame() {
	fmt.Printf("Red: %v\n", g.line["red"])
	fmt.Printf("Blue: %v\n", g.line["blue"])
	fmt.Printf("green: %v\n", g.line["green"])
}

type GameSettings struct {
	red   int
	green int
	blue  int
}

var (
	settings = newGameSetting(12, 13, 14)
)

func newGameSetting(r int, g int, b int) *GameSettings {
	return &GameSettings{
		red:   r,
		green: g,
		blue:  b,
	}
}

func Day2() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	result := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		stringLine := string(line)
		intValue, valid := getLineValues(stringLine)
		if valid {
			result += intValue
			fmt.Printf("GAME NO %v is %v", intValue, valid)
		} else {

			fmt.Printf("GAME NO %v is %v", intValue, valid)
		}
	}

	fmt.Printf("FINAL RESULT : %v \n", result)
}

func getLineValues(line string) (int, bool) {
	splLine := strings.Split(line, ":")
	game_id := strings.Split(splLine[0], " ")
	fmt.Println(game_id[1])
	id, _ := strconv.Atoi(game_id[1])
	values := strings.Split(splLine[1], ";")
	// game := newGame()
	checkValue := 0
	for _, v := range values {
		s := strings.Replace(v, ",", " ", -1)
		temp := strings.Split(s, " ")
		temp = temp[1:]
		res := getMap(temp)
		valdid, _ := isValid(res)
		if valdid {
			checkValue += id
		} else {
			return id, false
		}
	}
	return id, true
}

func getMap(slice []string) map[string]int {
	resultMap := make(map[string]int)
	slice = removeEmptyStrings(slice)
	for i := 0; i < len(slice)-1; i += 2 {
		if len(slice[i]) != 0 {
			key := slice[i+1]
			resultMap[key], _ = strconv.Atoi(slice[i])
		}
	}
	return resultMap
}

func removeEmptyStrings(input []string) []string {
	var result []string

	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

func isValid(m map[string]int) (bool, int) {
	sumOfLine := 0
	for k, v := range m {
		switch k {
		case "green":
			isValid, value := checkGreen(v)
			if isValid {
				sumOfLine += value
			} else {
				return false, 0
			}
		case "red":
			isValid, value := checkRed(v)
			if isValid {
				sumOfLine += value
			} else {
				return false, 0
			}
		case "blue":
			isValid, value := checkBlue(v)
			if isValid {
				sumOfLine += value
			} else {
				return false, 0
			}
		default:
			return false, 0
		}
	}
	return true, sumOfLine
}

func checkBlue(s int) (bool, int) {

	number := s
	if number > settings.blue {
		return false, 0
	}
	return true, number
}
func checkGreen(s int) (bool, int) {
	number := s

	if number > settings.green {
		return false, 0
	}
	return true, number
}

func checkRed(s int) (bool, int) {
	number := s
	if number > settings.red {
		return false, 0
	}
	return true, number
}
