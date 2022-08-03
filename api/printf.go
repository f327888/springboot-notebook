package main

import "fmt"

func main() {
	// 源码查看 %T %v %
	var firstName string
	var lastName string
	var email string
	var ticketNum uint

	fmt.Println("Please input your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please input your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please input your email")
	fmt.Scan(&email)

	fmt.Println("Please input your ticketNum")
	fmt.Scan(&ticketNum)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v ",
		firstName, lastName, ticketNum, email)

}
