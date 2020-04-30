package main
import (
	"fmt"
	"net/http"
	"time"
)
func check_status_Link(link string){
	_,error := http.Get(link)
	if error!=nil{
		fmt.Println(link,":-  The link is down")
		
	}else{
	   fmt.Println(link,":-   the link is up and running")
	}
}
func main(){
	my_links := []string{
		"https://www.facebook.com",
		"https://www.linkedin.com",
		"https://www.play.golang.org",
		"https://www.calsoft.org",
	}
	for _,link :=  range my_links{
	   go check_status_Link(link)
	}
	
    time.Sleep(12*time.Second)
}
