package scanner

import (
	"fmt"
	"testing"
)

func TestScanTokens(t *testing.T) {
	// source from https://www.craftinginterpreters.com/the-lox-language.html
	sourceList := []string{
		`
		// Your first Lox program!
		print "Hello, world!";	
		`,
		`
		true;  // Not false.
		false; // Not *not* false.
		`,
		`
		1234;  // An integer.
		12.34; // A decimal number.
		`,
		`
		"I am a string";
		"";    // The empty string.
		"123"; // This is a string, not a number.
		`,
		`
		add + me;
		subtract - me;
		multiply * me;
		divide / me;
		`,
		`
		-negateMe;
		`,
	}

	for i, source := range sourceList {
		t.Log("source", i, source)
		scanTokens(source)
	}
}

func scanTokens(source string) {
	scanner := &Scanner{}
	tokens := scanner.ScanTokens(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}
}
