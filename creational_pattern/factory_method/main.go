package main

import "fmt"

func main() {
	paymentFactory := &PaymentFactory{}
	cashPayment, _ := paymentFactory.GetPayment(Cash)
	if cashPayment.Pay(99) {
		fmt.Println("Cash paid successful")
	}
	debitPayment, _ := paymentFactory.GetPayment(DebitCard)
	if debitPayment.Pay(99) {
		fmt.Println("DebitCard paid successful")
	}
}
