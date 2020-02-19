/*
Word Counting application that prompts user for text file and counts the words in it
Returns currently as wordcount.txt output.
v1.5.0
Jacob Bogner 2/18/2020
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func getTextFiles(location string) []string{
	directory, err := os.Open(location) //Open the directory
	if err != nil {
		fmt.Println("Opening directory went horribly wrong: ", err)
	}
	directoryFiles, err := directory.Readdir(0) //Read the files in the directory
	if err != nil {
		fmt.Println("AAAAAANNNNNND file read error: ", err)
	}
	var fileArray []string //Holds the names of all the files
	for index := range directoryFiles { //Generate paths from files found in the directory
		newFile := directoryFiles[index]
		name := newFile.Name()
		fmt.Println("Found: ", name)
		file := location + name
		fileArray = append(fileArray, file)
		fmt.Println("The location is: ", fileArray[index])
	}
	return fileArray
}

func countWords(fileArray string) map[string]int {
	rawData, err := ioutil.ReadFile(fileArray) //attempt to open the file
	if err != nil { //Handle file opening error
		fmt.Println("AAAANNNND error: ", err)
	}
	unsortedWords := string(rawData) // convert file data to type String

	reg, err := regexp.Compile("[^a-zA-Z0-9 \n]+") //Define the set of "keeper" data
	if err != nil { //Handle reg compilation error
		log.Fatal(err)
	}
	trimmedData := reg.ReplaceAllString(unsortedWords, " ")//Read the unsorted data and throw out any chars that do not match with previously defined keeper data
	datamap := wordCounter(trimmedData) //Create a map file containing all trimmed data
	return datamap
}


func wordCounter(data string) map[string]int {
	/*
		Generates a table containing every word that appears in the input string, and the number of times it
		appears in the string. Outputs this data as a map
	*/
	list := strings.Fields(data) //break the string into words and store each as a slice
	datamap := make(map[string]int) //make the datamap
	for _, word := range list {
		_, repeat := datamap[word]
		if repeat { //If the word has already appeared...
			datamap[word] += 1 //add 1 to that word's corresponding int entry
		} else { //If the word is new
			datamap[word] = 1 //Initialize it into the map and assign its' int entry to 1
		}
	}
	return datamap
}

func writeWordCount(outputFile io.Writer, datamap map[string]int, filename string){
	fmt.Println("Attempting to write book info for ", filename)
	fmt.Fprintln(outputFile, "WORDCOUNT RESULTS: ", filename)
	for index, element := range(datamap){ //Print each word and its associated count to the command window and output file
		fmt.Println(index, " => ", element)
		fmt.Fprintln(outputFile, index, " => ", element)
	}
	fmt.Println("\nAttempted to write results to 'wordcount.txt'.")//Declare an attempt was made to write the file
}
/*
Word Counter function based off example codes
from: http://www.golangprograms.com/how-to-count-number-of-repeating-words-in-a-given-string.html
*/
func Task1() {
	var location = "assets/" //Declare a string to hold the name of the directory
	var datamap map[string]int
	keyboard := bufio.NewScanner(os.Stdin)
	fmt.Println("Gob's Program: Y/N?\n?: ") //Prompt user to run program
	keyboard.Scan()
	run := keyboard.Text() //read keyboard input
	outputFile, err := os.Create("wordcount.txt") //Create a 'wordcount.txt' file to write our datamap to
	if err != nil { //handle file creation error
		log.Fatal("There was a problem creating the file. ", err)
	}
	defer outputFile.Close()
	for run == "Y" || run == "y" {
		fmt.Println("Brace yourself")
		fileArray := getTextFiles(location) //Get the list of .txt files
		for index := range fileArray{ //For each file in the list
			datamap = countWords(fileArray[index]) //Count the words in the file
			writeWordCount(outputFile, datamap, fileArray[index]) //Print the word count to the output file
		}
		fmt.Println("\nGob's Program: Y/N?\n?: ") //Prompt user to run program
		keyboard.Scan()
		run = keyboard.Text() //read keyboard input
	}
}