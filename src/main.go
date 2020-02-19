/*
Word Counting application that prompts user for text file and counts the words in it
Returns currently as wordcount.txt output.
v1.0.0
Jacob Bogner 9/18/2018
*/
package main

import (
"fmt"
"io/ioutil"
"log"
"os"
"regexp"
"strings"
)


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
/*
Word Counter function based off example codes
from: http://www.golangprograms.com/how-to-count-number-of-repeating-words-in-a-given-string.html
*/


func main() {
	var directory string //Declare a string to hold the name of the directory
	fmt.Println("Enter the name of the directory that contains books you want me to read: (.txt files only)") //Prompt user to enter directory
	_, err := fmt.Scan(&directory) //read keyboard input
	rawData, err := ioutil.ReadFile(directory) //attempt to open the file
	if err != nil { //Handle file opening error
		fmt.Println(err)
		return
	}
	unsortedWords := string(rawData) // convert file data to type String

	reg, err := regexp.Compile("[^a-zA-Z0-9 \n]+") //Define the set of "keeper" data
	if err != nil { //Handle reg compilation error
		log.Fatal(err)
	}
	trimmedData := reg.ReplaceAllString(unsortedWords, " ")//Read the unsorted data and throw out any chars that do not match with previously defined keeper data
	datamap := wordCounter(trimmedData) //Create a map file containing all trimmed data

	for index, element := range(datamap){ //Print each word and its associated count to the command window
		fmt.Println(index, " => ", element)
	}

	outputFile, err := os.Create("wordcount.txt") //Create a 'wordcount.txt' file to write our datamap to
	if err != nil { //handle file creation error
		log.Fatal("There was a problem creating the file. ", err)
	}
	defer outputFile.Close()

	for index, element := range(datamap){ //Print each word and its associated count to the text file
		fmt.Fprintln(outputFile, index, " => ", element)
	}
	fmt.Println("\nAttempted to write results to 'wordcount.txt'.")//Declare an attempt was made to write the file

}
