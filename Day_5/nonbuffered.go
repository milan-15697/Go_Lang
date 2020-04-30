package main
import (
	"fmt"
	"net/http"
)
func check_status_link(check chan string, link string){
	_,err := http.Get(link)
	
	if err!=nil{
		check <- link+" down"
	}else{
	   check <- link+"   up and running"
	}
}

func main(){
	my_links := []string{
		"https://www.facebook.com",
		"https://www.linkedin.com",
		"https://www.play.golang.org",
		"https://www.calsoft.org",
	}

	check := make(chan string)
	
	for _,link :=  range my_links{
	   go check_status_link(check,link)
	}
	
	fmt.Println(<-check)
	fmt.Println(<-check)
	fmt.Println(<-check)
	fmt.Println(<-check)
}
