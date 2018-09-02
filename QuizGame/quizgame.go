// Quizz game without timeout
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// OpenCSV loads the csv file and returns your rows
func OpenCSV(filename string) [][]string {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err)
	}
	r := csv.NewReader(strings.NewReader(string(file)))

	r.Comma = ','
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	return records
}

func quizGame(problems [][]string) bool {
	var userInput string
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	tam := len(problems)
	row := r1.Intn(tam)

	problem, answer := problems[row][0], problems[row][1]
	fmt.Printf("%v = ", problem)
	fmt.Scanln(&userInput)
	var result = true

	if answer != userInput {
		result = false
	}

	return result
}

func main() {
	filename := flag.String("filename", "problems.csv", "Digite o banco de dados das perguntas")
	flag.Parse()
	csvFile := OpenCSV(*filename)
	total := 0
	for i := 0; i < len(csvFile); i++ {
		if quizGame(csvFile) {
			println("Correct")
			total++
		}
	}
}
