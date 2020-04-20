package main
import "fmt"
func function1() {
	fmt.Println("function1")
}
func function2() string{
    defer function1()
	return "function2"
}
func main() {
	fmt.Println(function2())
}
