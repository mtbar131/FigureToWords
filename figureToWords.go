package main
import "os"
import "fmt"

func converter(s string) string {
    runes := []rune(s)
    
    return string(runes)
}
func main() {

inputNumber := os.Args[1:]

 
    fmt.Println(convert(inputNumber))
    
}
