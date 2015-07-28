package main
import "os"
import "fmt"
func main() {

argsWithoutProg := os.Args[1:]

 
    fmt.Println(argsWithoutProg)
    
}
