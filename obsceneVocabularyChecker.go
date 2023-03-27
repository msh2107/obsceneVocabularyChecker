package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName, sentence string
	tabooWords := make(map[string]struct{})

	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	readFile(file, tabooWords)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sentence = scanner.Text()
		if sentence == "exit" {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		sentence = checkSentence(sentence, tabooWords)
		fmt.Println(sentence)
	}

	fmt.Println("Bye!")

}

func readFile(file *os.File, tabooWords map[string]struct{}) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		tabooWords[scanner.Text()] = struct{}{}
	}
}

func checkSentence(sentence string, tabooWords map[string]struct{}) string {
	words := strings.Fields(sentence)
	for _, word := range words {
		if _, ok := tabooWords[strings.ToLower(word)]; ok {
			sentence = strings.Replace(sentence, word, strings.Repeat("*", len(word)), -1)
		}
	}
	return sentence
}
