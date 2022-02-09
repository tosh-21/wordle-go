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

func GetWord() (string, []string) {
	f, err := os.Open("words2.txt")

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pick, ListOfWords
}

func UserWord(trys int, resultArray [][]string) []string {

	attempts := trys
	CorrectGuess := []string{"o", "o", "o", "o", "o"}
	//var finalArray [][]string
	wordOfTheDay, wordList := GetWord()
	fmt.Println("Try #: ", attempts)
	fmt.Println("Enter your guess: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error while reading input")
	}

	trimmedInput := strings.TrimSpace(input)
	inputChar := strings.Split(trimmedInput, "")
	wordleChar := strings.Split(wordOfTheDay, "")
	//fmt.Println(inputChar)
	//fmt.Println(wordOfTheDay)

	if len(inputChar) != 5 {
		fmt.Println("Please enter a five letter word.")
		UserWord(attempts, resultArray)
	} else if WordCheck(trimmedInput, wordList) == false {
		fmt.Println("Word does not exist. Try again.")
		UserWord(attempts, resultArray)
	} else {
		attempts += 1
		fmt.Println(LetterExists(inputChar, wordleChar))
		resultArray = AppendSlice(resultArray, LetterExists(inputChar, wordleChar))
		if attempts < 7 {
			if reflect.DeepEqual(LetterExists(inputChar, wordleChar), CorrectGuess) {
				fmt.Printf("Correct in %d tries \n", attempts-1)
				PrintArray(resultArray)
			} else {
				fmt.Println("Try again")
				UserWord(attempts, resultArray)
			}
		} else {
			fmt.Println("Maximum tries reached")
			PrintArray(resultArray)
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

func WordCheck(userInput string, wordList []string) bool {
	wordMap := make(map[string]bool)
	for _, value := range wordList {
		wordMap[value] = true
	}
	return wordMap[userInput]
}

func AppendSlice(source [][]string, value []string) [][]string {
	return append(source, value)
}

func PrintArray(array [][]string) {
	for _, j := range array {
		for _, i := range j {
			fmt.Printf("%s ", i)
		}
		fmt.Printf("\n")
	}
}
