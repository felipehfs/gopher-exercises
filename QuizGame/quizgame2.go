// Quizzgame with timeout
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

func quizGame(problems [][]string, progress chan bool) {
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

	progress <- result
}

func main() {
	filename := flag.String("filename", "problems.csv", "Digite o banco de dados das perguntas")
	t := flag.Int("time", 30, "Digite o tempo de espera válido")
	flag.Parse()
	csvFile := OpenCSV(*filename)
	timeout := time.After(time.Second * time.Duration(*t))
	total := 0
	prog := make(chan bool)
	finish := make(chan bool)

	go func() {
		for {
			select {
			case <-finish:
				break
			default:
				quizGame(csvFile, prog)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-timeout:
				finish <- true
				break
			case p := <-prog:
				if p {
					total += 1
				}
			}
		}
	}()
	<-finish
	fmt.Printf("\nScore: %v\n", total)
}
