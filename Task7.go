package main
import (
	"fmt"
	"io/ioutil"
	"time"
	"os"
)

type fileOperation struct {
	vowels, Puntutations, newlines, words, spaces float64
}

// func Files(createChannel chan fileOperation, data string) {
// 	operation := fileOperation{}

// 	for _, char := range data {
// 		if string(char) == "a" || string(char) == "e" || string(char) == "i" || string(char) == "o" || string(char) == "u" ||
// 			string(char) == "A" || string(char) == "E" || string(char) == "I" || string(char) == "O" || string(char) == "U" {
// 			operation.vowels++
// 		}
// 		if string(char) == "!" || string(char) == "@" || string(char) == "#" || string(char) == "$" || string(char) == "%" ||
// 			string(char) == "^" || string(char) == "&" || string(char) == "*" || string(char) == "(" || string(char) == ")" ||
// 			string(char) == "-" || string(char) == "_" || string(char) == "+" || string(char) == "=" || string(char) == "{" ||
// 			string(char) == "}" || string(char) == "[" || string(char) == "]" || string(char) == "|" || string(char) == "\\" ||
// 			string(char) == ":" || string(char) == ";" || string(char) == "'" || string(char) == "<" || string(char) == ">" ||
// 			string(char) == "," || string(char) == "." || string(char) == "/" || string(char) == "?" {
// 			operation.Puntutations++
// 		}
// 		if string(char) == "\n" {
// 			operation.newlines++
// 		}
// 		if string(char) == " " {
// 			operation.spaces++
// 		}
// 		operation.words++
// 	}

// 	createChannel <- operation
// }

func main() {
	starttime:=time.Now()
	argsWithProg := os.Args[1]
	file, err := ioutil.ReadFile("task.txt")
	if err != nil {
		fmt.Println("Error in opening file:", err)
		return
	}

	fmt.Println(string(file))

	createChannel := make(chan fileOperation)
	for i:=0;i<argsWithProg;i++{
		go Files(createChannel, string(file[:len(file)/argsWithProg]))
		go Files(createChannel, string(file[len(file)/argsWithProg:]))
	}

	vowel := float64(0)
	punctutation := float64(0)
	newline := float64(0)
	word := float64(0)
	space := float64(0)

	for i := 0; i < argsWithProg; i++ {
		result := <-createChannel
		vowel += result.vowels
		punctutation += result.Puntutations
		newline += result.newlines
		word += result.words
		space += result.spaces
	}
	excetiontime:=time.Since(starttime)
	fmt.Println("vowels are ", vowel)
	fmt.Println("Puntutations are ", punctutation)
	fmt.Println("new lines are ", newline)
	fmt.Println("number of words are ", word)
	fmt.Println("number of free spaces are ", space)
	fmt.Println("Exceutiontime",excetiontime)
}
