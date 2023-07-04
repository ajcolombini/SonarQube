package main

import (
	"fmt"
	"testing"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func TestSum(t *testing.T) {
    result := sum(2,3)
    if result != 5 {
        t.Error("Result faild")
    }

}

func sum(a int, b int) int {
	return a + b
}

