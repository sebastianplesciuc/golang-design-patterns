package main

import (
	"fmt"
	"os"
)

/*
		The Single Responsibility Principle (SRP)

 		A module/class (in Golang, a structure) should only have responsibility over a single part
	 of te functionality provided by the software (i.e. it should only do one thing and not more than it should).
	 This is also referred to as High Cohesion.

	Ref: https://en.wikipedia.org/wiki/Single_responsibility_principle
 	Ref: https://stackoverflow.com/questions/10620022/example-of-single-responsibility-principle
*/

type Logger struct {
	logEntries []string
}

func (l *Logger) Log(entry string) {
	l.logEntries = append(l.logEntries, entry)
}

// !!! This is wrong. This violates SRP
func (l *Logger) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, entry := range l.logEntries {
		f.WriteString(entry + "\n")
	}
	return nil
}

// !!! This is better
type LogFileWriter struct {
}

func (lfw *LogFileWriter) Save(logger *Logger, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, entry := range logger.logEntries {
		f.WriteString(entry + "\n")
	}
	return nil
}

func main() {
	fmt.Println("Single Responsibility Principle")

	logger := &Logger{}
	logger.Log("entry1")
	logger.Log("entry2")

	// !!! This is wrong. This violates SRP
	logger.Save("./wrong.log")

	// !!! This is better
	logWriter := &LogFileWriter{}
	logWriter.Save(logger, "./better.log")
}
