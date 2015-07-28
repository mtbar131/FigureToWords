package main
import "os"
import "fmt"
import "strconv"


func converter(s string) string {

	var answer = ""	
	var single_digits = []string{ "zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}
	

	var two_digits = []string{"", "ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	

	var tens_multiple = []string{"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}
	
	var tens_power = []string{"hundred", "thousand", "lakhs", "Crores"}

	var num, zero int 
	var err1, err2 error
	if len(s) == 1 {
		
		num, err1 = strconv.Atoi(s)
		zero, err2 = strconv.Atoi("0")


		if err1 != nil {
			fmt.Println(err1)
		}

		if err2 != nil {
			fmt.Println(err2)
		}
		
		return single_digits[num - zero]
	}

	fmt.Println(single_digits);
	fmt.Println(two_digits);
	fmt.Println(tens_power);
	fmt.Println(tens_multiple);
    
    return string(answer)

}

func main() {

	inputNumber := os.Args[1:]

	fmt.Println(converter("8"))
	fmt.Println(inputNumber)
	
}
