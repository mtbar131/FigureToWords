package main

import "fmt"


func converter(position int, s string) string{

	var str = "";
	var one = []string{""," one"," two"," three"," four"," five"," six"," seven"," eight"," Nine"," ten"," eleven"," twelve"," thirteen"," fourteen","fifteen"," sixteen"," seventeen"," eighteen"," nineteen"}
	var ten = []string{"",""," twenty"," thirty"," forty"," fifty"," sixty", " seventy"," eighty"," ninety"}
	
	if position > 19{
		str += ten[position / 10]+" "+one[position % 10]
	} else{
		str += one[position]
	}
	if position > 0 {
		return str + s
	}
	return str
}

func printToWord(inputNumber int) string{

	var ans = ""
	var tensPowers = []int{1000000000, 10000000, }
	ans += converter((inputNumber  / 1000000000), " Hundred")
	ans += converter((inputNumber  / 10000000) % 100, " Crore")
	ans += converter(((inputNumber / 100000)%100)," lakh")
	ans += converter(((inputNumber / 1000)%100)," thousand")		
	ans += converter(((inputNumber / 100)%10)," hundred")
	if inputNumber % 100 != 0 && inputNumber > 100 {
		ans += " and"
	}
	ans += converter((inputNumber % 100)," ")
	fmt.Println()	 	
	return ans
}	



func main() {

	var inputNumber int = 0
	
	for true{
		_, err  := fmt.Scanf("%d", &inputNumber)
		if inputNumber == 0 {
			fmt.Println("Zero")
		} else {
			if err == nil {
				fmt.Println(printToWord(inputNumber))
			}
		}
	}
	
}
