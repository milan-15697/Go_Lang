package main

import "fmt"

func Prime(val int) bool {
    for i := 2; i <= val/2; i++ {
        if val%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
	var no int
	fmt.Scan(&no)
        if Prime(no) {
            fmt.Printf("%v is Prime\n", no)
		} else{
			fmt.Printf("%v is not Prime\n",no)
		}
}
