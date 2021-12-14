/*
Word Counting application that prompts user for text file and counts the words in it
Returns currently as wordcount.txt output. Makes use of goroutines (multi-threading) to divide load!!!
v1.6.0
Jacob Bogner 2/18/2020
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func Task2() {
	var location = "assets/" //Declare a string to hold the name of the directory
	keyboard := bufio.NewScanner(os.Stdin)
	fmt.Println("Run Program: Y/N?\n?: ") //Prompt user to run program
	keyboard.Scan()
	run := keyboard.Text()                        //read keyboard input
	outputFile, err := os.Create("wordcount.txt") //Create a 'wordcount.txt' file to write our datamap to
	if err != nil {                               //handle file creation error
		log.Fatal("There was a problem creating the file. ", err)
	}
	defer outputFile.Close()
	for run == "Y" || run == "y" {
		fmt.Println("Brace yourself")
		wordCount := make(chan map[string]int)
		fileArray := getTextFiles(location) //Get the list of .txt files
		total := 4                          //Current total number of goroutines we will run
		var wg sync.WaitGroup
		wg.Add(total)

		for ctr := 0; ctr < total; ctr++ {
			go wordCounterRoutine(wordCount, fileArray, &wg)
		}

		datamap := <-wordCount
		// Wait for all goroutines to be finished
		wg.Wait()

		writeWordCount(outputFile, datamap, "INSERTED VIA GOROUTINE") //Print the word count to the output file
		fmt.Println("\nRun Program: Y/N?\n?: ")                       //Prompt user to run program
		keyboard.Scan()
		run = keyboard.Text() //read keyboard input
	}
}

func wordCounterRoutine(wordCount chan<- map[string]int, fileArray []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := range fileArray { //For each file in the list
		datamap := countWords(fileArray[index]) //Count the words in the file
		wordCount <- datamap
	}
}
