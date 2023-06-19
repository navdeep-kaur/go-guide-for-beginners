package main

import (
	"errors"
	"fmt"
	"time"
)

// DateError is a custom error that will fulfill the Error interface
type DateError struct {
	Message string
	Date    time.Time
}

// Our Method receiver where we Apply the Error() string function that the error interface needs to the DateError type
func (de DateError) Error() string {
	return fmt.Sprintf("%s: %s", de.Date.String(), de.Message)
}

func NewError(message string) DateError {
	return DateError{
		Message: message,
		Date:    time.Now(),
	}
}

// PrintError can take in an DateError since DateError fulfills the error interface
func PrintError(err error) {
	fmt.Println(err.Error())
}

func main() {
	de := NewError("Oh oh, I failed")
	regularErr := errors.New("I am your regular error")
	PrintError(de)
	PrintError(regularErr)
}
