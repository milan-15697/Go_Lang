package main

import "fmt"

type student struct {
	first_name string
	last_name  string
	phone contact
	gradepoint      []int
}
type contact struct {
	ph_no int
}

func main() {
	info := student{first_name: "Milan",
		last_name:  "Bhardwaj",
		phone: contact{ph_no: 12345},
		gradepoint:      []int{200, 200}}
	print(info.contact_no.phone_number)
}
func (i student) print() {

	fmt.Println(i)
}
