package scanner

import (
	"fmt"
	"testing"
)

// source from https://www.craftinginterpreters.com/the-lox-language.html

func TestScanHello(t *testing.T) {
	sourceList := []string{
		`
		// Your first Lox program!
		print "Hello, world!";	
		`,
	}

	scanSourceList(sourceList)
}

func TestScanDataTypes(t *testing.T) {
	sourceList := []string{
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
	}

	scanSourceList(sourceList)
}

func TestScanExpressions(t *testing.T) {
	sourceList := []string{
		`
		add + me;
		subtract - me;
		multiply * me;
		divide / me;
		`,
		`
		-negateMe;
		`,
		`
		less < than;
		lessThan <= orEqual;
		greater > than;
		greaterThan >= orEqual;
		`,
		`
		1 == 2;         // false.
		"cat" != "dog"; // true.
		`,
		`
		314 == "pi";
		`,
		`
		123 == "123";
		`,
		`
		!true;  // false.
		!false; // true.
		`,
		`
		true and false; // false.
		true and true;  // true.
		`,
		`
		false or false; // false.
		true or false;  // true.
		`,
		`
		var average = (min + max) / 2;
		`,
	}

	scanSourceList(sourceList)
}

func TestScanStatements(t *testing.T) {
	sourceList := []string{
		`
		print "Hello, world!";
		`,
		`
		"some expression";
		`,
		`
		{
			print "One statement.";
			print "Two statements.";
		}
		`,
	}

	scanSourceList(sourceList)
}

func TestScanVariables(t *testing.T) {
	sourceList := []string{
		`
		var imAVariable = "here is my value";
		var iAmNil;
		`,
		`
		var breakfast = "bagels";
		print breakfast; // "bagels".
		breakfast = "beignets";
		print breakfast; // "beignets".
		`,
	}

	scanSourceList(sourceList)
}

func TestScanControlFlow(t *testing.T) {
	sourceList := []string{
		`
		if (condition) {
			print "yes";
		} else {
			print "no";
		}
		`,
		`
		var a = 1;
		while (a < 10) {
			print a;
			a = a + 1;
		}
		`,
		`
		for (var a = 1; a < 10; a = a + 1) {
			print a;
		}
		`,
	}

	scanSourceList(sourceList)
}

func scanSourceList(sourceList []string) {
	for i, source := range sourceList {
		fmt.Println("source", i, source)
		scanSource(source)
	}
}

func scanSource(source string) {
	scanner := &Scanner{}
	tokens := scanner.ScanTokens(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}
}
