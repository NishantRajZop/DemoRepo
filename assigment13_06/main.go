package main

/*
   bufio is a Package which is used to read buffered Input and Outputs
   os is a Package there which is used for doing FileManipulations , reading Enviroment Variables
   strings is a Package which Gives You builtin methods for manipulationg Strings , like - strings.hasPrefix(s , word)
   time package provides multiple built in functions to work with current time and date
   fmt is for Printing and taking Inputs
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func percent(count, totalLines int) float64 {
	if totalLines == 0 {
		return 0.0
	}
	return float64(count) / float64(totalLines) * 100
}
func main() {

	if len(os.Args) < 2 {
		fmt.Println("please Provide a File Name to read Logs From")
		return
	}

	fmt.Println("reading this file - > ", os.Args[0])
	// os.Args[0] always refers to the name of the program itself, i.e., the path used to run the executable.

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	/* it is necessary to close the File otherWise it can cause a Memory leak there */

	var infoCount, warningCount, errorCount, totalLines int

	scanner := bufio.NewScanner(file) // creates a Scanner to read file
	for scanner.Scan() {              // scanner.Scan() next Line , Here For scanner.Scan() means untill there is next line run loop
		line := scanner.Text() // scanner.Text() gets the Next Line as a String
		totalLines++

		switch {
		case strings.HasPrefix(line, "[INFO]"):
			infoCount++
		case strings.HasPrefix(line, "[WARNING]"):
			warningCount++
		case strings.HasPrefix(line, "[ERROR]"):
			errorCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	t := time.Now()

	fmt.Println("Log Analysis of file:", filename)
	fmt.Printf("INFO    : %d entries (%.2f%%)\n", infoCount, percent(infoCount, totalLines))
	fmt.Printf("WARNING : %d entries (%.2f%%)\n", warningCount, percent(warningCount, totalLines))
	fmt.Printf("ERROR   : %d entries (%.2f%%)\n", errorCount, percent(errorCount, totalLines))
	fmt.Println()
	fmt.Printf("Total   : %d lines\n", totalLines)
	fmt.Print("Analyzed at -> ")
	fmt.Println(t.Date())
}

// fmt.Println() vs fmt.Printf()

//fmt.Println()	Prints values with automatic spacing and a newline at the end. Best for simple output.
//fmt.Printf()	Prints formatted output using format specifiers (like %d, %s, %.2f, etc.). More control. No newline unless you add \n
