package main
import "fmt"
func slice() {
list:=[]int{10,20,30}
list = append(list,40)
list = append(list,50)
fmt.Println(list)
list1:=list
list1=append(list1,60)
fmt.Println(list1)
}
func main() {
	slice()
}
