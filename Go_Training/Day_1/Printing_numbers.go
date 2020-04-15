package main

import "fmt"

func main() {
	k:=1
	for i:=1; i<=5; i++{
    		for j:=0; j<i; j++{
			fmt.Printf("%d ",k);
			k++
        }
        fmt.Println()
}

}
