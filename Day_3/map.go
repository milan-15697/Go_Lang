package main
import "fmt"
func map1() {
dict := map[int]string{1:"Ronaldo"}
dict[2]="Messi"
fmt.Println(dict)
dict[1]="Maradona"
fmt.Println(dict)
delete(dict,1)
fmt.Println(dict)
}
func main() {
	map1()
}
