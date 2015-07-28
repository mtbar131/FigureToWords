package main
import "os"
import "fmt"


func converter(s string) string {

	var single_digits = []string{ "zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}
	

	var two_digits = []string{"", "ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	

	var tens_multiple = []string{"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}
	
	var tens_power = []string{"hundred", "thousand", "lakhs", "Crores"}
	// fmt.Println(single_digits);
	// fmt.Println(two_digits);
	// fmt.Println(tens_power);
	// fmt.Println(tens_multiple);

}

func main() {

	inputNumber := os.Args[1:]

	converter("test")
	fmt.Println(inputNumber)
	
}
