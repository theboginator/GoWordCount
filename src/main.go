/*
Word Counting application that prompts user for text file and counts the words in it
Returns currently as wordcount.txt output.
v1.0.0
Jacob Bogner 9/18/2018
*/
package main

import (
	"bufio"
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
		directory, err := os.Open(location) //Open the directory
		if err != nil {
			fmt.Println("Opening directory went horribly wrong: ", err)
			return
		}
		directoryFiles, err := directory.Readdir(0) //Read the files in the directory
		if err != nil {
			fmt.Println("AAAAAANNNNNND file read error: ", err)
			return
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
		for index := range fileArray{
			rawData, err := ioutil.ReadFile(fileArray[index]) //attempt to open the file
			if err != nil { //Handle file opening error
				fmt.Println("AAAANNNND error: ", err)
				return
			}
			unsortedWords := string(rawData) // convert file data to type String

			reg, err := regexp.Compile("[^a-zA-Z0-9 \n]+") //Define the set of "keeper" data
			if err != nil { //Handle reg compilation error
				log.Fatal(err)
			}
			trimmedData := reg.ReplaceAllString(unsortedWords, " ")//Read the unsorted data and throw out any chars that do not match with previously defined keeper data
			datamap = wordCounter(trimmedData) //Create a map file containing all trimmed data
			fmt.Println("Attempting to write book info for ", fileArray[index])
			fmt.Fprintln(outputFile, "WORDCOUNT RESULTS: ", fileArray[index])
			for index, element := range(datamap){ //Print each word and its associated count to the command window and output file
				fmt.Println(index, " => ", element)
				fmt.Fprintln(outputFile, index, " => ", element)
			}
			fmt.Println("\nAttempted to write results to 'wordcount.txt'.")//Declare an attempt was made to write the file
		}
		fmt.Println("\nGob's Program: Y/N?\n?: ") //Prompt user to run program
		keyboard.Scan()
		run = keyboard.Text() //read keyboard input
	}
}
