package main
import (
	"fmt"
	"net/http"
)
func check_status_link(check chan string, link string){
	_,error := http.Get(link)
	if error!=nil{
		check <- link+"  is down"		
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

	check := make(chan string,2)
	for _,link :=  range my_links{
	   go check_status_link(check,link)
	   fmt.Println(<-check)
	}
}
