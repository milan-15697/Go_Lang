package main
import (
	"fmt"
)
type userPrivilage interface{
	 deposit(amount float64) 
	 withdraw(amount float64)
	 checkBalance() float64
}
type Current_Account struct{
	balance float64
}
type Savings_Account struct{
	balance float64
	Rate_of_interest float64
}
func (CA *Current_Account) deposit(amount float64) {
	CA.balance = CA.balance + amount
	
}
func (CA *Current_Account) withdraw(amount float64) {
  CA.balance-=amount
  
}
func (CA *Current_Account) checkBalance() float64 {
  return CA.balance
}
func (SA *Savings_Account) deposit(amount float64) {
	SA.balance+=(amount+SA.Rate_of_interest*amount)
}
func (SA *Savings_Account) withdraw(amount float64) {
  SA.balance-=amount
}
func (SA *Savings_Account) checkBalance() float64 {
  return SA.balance
}
func account_details(u userPrivilage){
	fmt.Println("Total Available Balance:- ",u.checkBalance())
	fmt.Println("Deposit Money:-  150")
	u.deposit(150)
	fmt.Println("Withdraw Money:-  50",)
	u.withdraw(50) 
	fmt.Println("Current Balance",u.checkBalance())
	
}
func main(){
CA:=Current_Account{balance:0}
SA:=Savings_Account{balance:0,Rate_of_interest:0.10}

fmt.Println("Available balance in CA")
account_details(&CA)
fmt.Println("Available balance in SA")
account_details(&SA)

}
