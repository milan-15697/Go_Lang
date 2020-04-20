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

	fmt.Println("Before Update")
	print(info.contact_no.phone_number)
	fmt.Println("\nAfter Update")
	ptr_info := &info
	ptr_info.contact_no.phone_number = 98765
	print(info.contact_no.phone_number)
}
func (stu student) print() {

	fmt.Println(stu)
}
