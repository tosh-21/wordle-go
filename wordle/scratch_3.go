package main

import "fmt"

func main() {
	array1 := []string{"a","b","c","d","e"}
	array2 := []string{"f","g","h","i","j"}
	array3 := []string{"k","l","m","n","o"}


	var finalArray [][]string
	finalArray = AppendSlice(finalArray, array1)
	finalArray = AppendSlice(finalArray, array2)
	finalArray = AppendSlice(finalArray, array3)


	PrintArray(finalArray)
	PrintArray2(finalArray)

}

func AppendSlice(source [][]string, value []string) [][]string {
	return append(source, value)
}

func PrintArray(array [][]string){
	for i := 0; i < len(array); i++{
		for j := 0; j <= len(array); j++{
			fmt.Printf("%v ", array[i][j])
		}
		fmt.Printf("\n")
		//fmt.Println(len(array))
	}

}

func PrintArray2(array [][]string){
	for _, j := range array{
		for _, i := range j{
			fmt.Printf("%s ", i)
		}
		fmt.Printf("\n")
	}
}

//func WordCheck(userInput string, wordList []string) string {
//	sort.Strings(wordList)
//	find := sort.SearchStrings(wordList, userInput)
//	wordMap := make(map[int]string)
//	for key, value := range wordList {
//		wordMap[key] = value
//	}
//	return wordMap[find]
//}
//
//func WordCheck2(userInput string, wordList []string) bool {
//	wordMap := make(map[string]bool)
//	for _, value := range wordList {
//		wordMap[value] = true
//	}
//	return wordMap[userInput]
//}
