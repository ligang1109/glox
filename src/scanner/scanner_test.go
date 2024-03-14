package scanner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// source from https://www.craftinginterpreters.com/the-lox-language.html

func TestScanHello(t *testing.T) {
	sourceList := []string{
		`
		// Your first Lox program!
		print "Hello, world!";	
		`,
	}

	scanSourceList(t, sourceList)
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

	scanSourceList(t, sourceList)
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

	scanSourceList(t, sourceList)
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

	scanSourceList(t, sourceList)
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

	scanSourceList(t, sourceList)
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

	scanSourceList(t, sourceList)
}

func TestScanFunctions(t *testing.T) {
	sourceList := []string{
		`
		makeBreakfast(bacon, eggs, toast);
		`,
		`
		makeBreakfast();
		`,
		`
		fun printSum(a, b) {
			print a + b;
		}
		`,
		`
		fun returnSum(a, b) {
			return a + b;
		}
		`,
		`
		fun addPair(a, b) {
			return a + b;
		}
		  
		fun identity(a) {
			return a;
		}
		
		print identity(addPair)(1, 2); // Prints "3".	
		`,
		`
		fun outerFunction() {
			fun localFunction() {
			  print "I'm local!";
			}
		  
			localFunction();
		}
		`,
		`
		fun returnFunction() {
			var outside = "outside";
		  
			fun inner() {
			  print outside;
			}
		  
			return inner;
		}
		  
		var fn = returnFunction();
		fn();
		`,
	}

	scanSourceList(t, sourceList)
}

func TestScanClasses(t *testing.T) {
	sourceList := []string{
		`
		class Breakfast {
			cook() {
			  print "Eggs a-fryin'!";
			}
		  
			serve(who) {
			  print "Enjoy your breakfast, " + who + ".";
			}
		}
		`,
		`
		// Store it in variables.
		var someVariable = Breakfast;

		// Pass it to functions.
		someFunction(Breakfast);
		`,
		`
		var breakfast = Breakfast();
		print breakfast; // "Breakfast instance".
		`,
		`
		breakfast.meat = "sausage";
		breakfast.bread = "sourdough";
		`,
		`
		class Breakfast {
			serve(who) {
			  print "Enjoy your " + this.meat + " and " +
				  this.bread + ", " + who + ".";
			}
		  
			// ...
		}
		`,
		`
		class Breakfast {
			init(meat, bread) {
			  this.meat = meat;
			  this.bread = bread;
			}
		  
			// ...
		}
		  
		var baconAndToast = Breakfast("bacon", "toast");
		baconAndToast.serve("Dear Reader");
		// "Enjoy your bacon and toast, Dear Reader."
		`,
		`
		class Brunch < Breakfast {
			drink() {
			  print "How about a Bloody Mary?";
			}
		}
		`,
		`
		var benedict = Brunch("ham", "English muffin");
		benedict.serve("Noble Reader");
		`,
		`
		class Brunch < Breakfast {
			init(meat, bread, drink) {
			  super.init(meat, bread);
			  this.drink = drink;
			}
		}
		`,
	}

	scanSourceList(t, sourceList)
}

func scanSourceList(t *testing.T, sourceList []string) {
	for i, source := range sourceList {
		fmt.Println("source", i, source)
		scanSource(t, source)
	}
}

func scanSource(t *testing.T, source string) {
	scanner := &Scanner{}
	tokens := scanner.ScanTokens(source)
	for i, token := range tokens {
		fmt.Println(i, token)
	}

	assert.False(t, scanner.HasError())
}
