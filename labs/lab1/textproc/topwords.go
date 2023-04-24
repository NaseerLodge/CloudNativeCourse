// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Topwords(path string, K int) []WordCount {

	//Open file
	itemsSlice := make([]string, 0)
	dat, err := os.Open(path)
	checkError(err)

	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	check := make(map[string]int)
	res := make([]string, 0)
	FinalAnswer := make([]WordCount, 0)

	//Scan the file and store the words
	//into a slice
	for scanner.Scan() {

		line := scanner.Text()
		items := strings.Split(line, " ")
		itemsSlice = append(itemsSlice, items...)

	}

	//Error check for scanning file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//Removes repeating strings from the slice
	//and stores the values in a new slice
	for _, val := range itemsSlice {
		check[val] = 1
	}

	for letter := range check {
		res = append(res, letter)
	}

	//returns a wordcount slice with the strings and their respective
	//count
	Done := ConverttoMaps(itemsSlice, res, len(itemsSlice), len(res))

	//Sort by number of string apperances
	sortWordCounts(Done)

	for l := 0; l < K; l++ {

		FinalAnswer = append(FinalAnswer, Done[l])
	}

	return FinalAnswer
}

var FinalSlice []string

func search(element []string, s string, n int) int {

	counter := 0

	for j := 0; j < n; j++ {

		if s == element[j] {
			counter++
		}
	}

	return counter
}

// returns a wordcount slice with the strings and their respective
// count
func ConverttoMaps(Full []string, Short []string, FullLength int, ShortLength int) []WordCount {

	FinalSlice := make([]WordCount, 0)

	for i := 0; i < ShortLength; i++ {

		NumberofRepeat := search(Full, Short[i], FullLength)

		s := WordCount{Short[i], NumberofRepeat}

		FinalSlice = append(FinalSlice, s)
	}

	return FinalSlice
}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
