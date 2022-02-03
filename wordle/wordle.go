package wordle

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var isAlphaNum = regexp.MustCompile("^[a-z]*$")

//var trys int = 1

func GetWord() string {
	f, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var ListOfWords []string

	for scanner.Scan() {

		word := scanner.Text()

		if isAlphaNum.MatchString(word) {
			wordLength := len([]rune(word))
			if wordLength == 5 {
				ListOfWords = append(ListOfWords, word)
			}

		}

	}

	dateNow := time.Now().Format("01-02-2006")
	parseTime, e := time.Parse("01-02-2006", dateNow)
	if e != nil {
		panic("Can't parse")
	}
	intTime := parseTime.Unix()

	seed1 := rand.NewSource(intTime)
	rand1 := rand.New(seed1)

	randomIndex := rand1.Intn(len(ListOfWords))
	pick := ListOfWords[randomIndex]
	//fmt.Println(pick)
	return pick

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ""

}

func UserWord(trys int) []string {

	attempts := trys
	CorrectGuess := []string{"o", "o", "o", "o", "o"}
	wordOfTheDay := GetWord()
	fmt.Println("Try #: ", attempts)
	fmt.Println("Enter your guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error while reading input")
	}

	//var arrayChar []string
	trimmedInput := strings.TrimSpace(input)
	inputChar := strings.Split(trimmedInput, "")
	wordleChar := strings.Split(wordOfTheDay, "")
	//fmt.Println(inputChar)
	//fmt.Println(wordleChar)

	if len(inputChar) != 5 {
		fmt.Println("Please enter a five letter word.")
		UserWord(attempts)
	} else {
		attempts += 1
		fmt.Println(LetterExists(inputChar, wordleChar))

		if attempts < 7 {
			if reflect.DeepEqual(LetterExists(inputChar, wordleChar), CorrectGuess) {
				fmt.Printf("Correct in %d tries", attempts)
			} else {
				fmt.Println("Try again")
				UserWord(attempts)
			}
		} else {
			fmt.Println("Maximum tries reached")
		}
	}
	return inputChar
}

const greenBox = "🟩"
const yellowBox = "🟨"
const greyBox = "⬜"

func LetterExists(guess []string, wordle []string) []string {
	var clueArray []string
	for k, v := range guess {
		//fmt.Println(k, v)
		//fmt.Println(wordle[k])
		if v == wordle[k] {
			clueArray = append(clueArray, "o")
		} else if YellowBox(v, wordle) {
			clueArray = append(clueArray, "x")
		} else {
			clueArray = append(clueArray, "_")
		}
	}

	return clueArray

}

func YellowBox(input string, answer []string) bool {
	for _, value := range answer {
		if value == input {
			return true
		}
	}
	return false
}
