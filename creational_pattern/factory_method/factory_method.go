package main

import (
	"fmt"
	"time"
)

type PaymentType int

const (
	Cash PaymentType = iota
	DebitCard
)

type IPaymentMethod interface {
	Init() IPaymentMethod
	Pay(amount float32) bool
}

type IPaymentFactory interface {
	GetPayment(string) (IPaymentMethod, error)
}

type CashPayment struct{}

func (c *CashPayment) Pay(amount float32) bool {
	return true
}

func (c *CashPayment) Init() IPaymentMethod {
	return c
}

type DebitCardPayment struct {
	balance float32
	expire  time.Time
}

func (d *DebitCardPayment) Pay(amount float32) bool {
	now := time.Now().In(time.UTC)
	if d.balance < amount || d.expire.Before(now) {
		return false
	}
	return true
}

func (d *DebitCardPayment) SetBalance(balance float32) IPaymentMethod {
	d.balance = balance
	return d
}

func (d *DebitCardPayment) SetExpire(date string) IPaymentMethod {
	const shortForm = "2006-01-02 15:04:05"
	d.expire, _ = time.Parse(shortForm, date)
	return d
}

func (d *DebitCardPayment) Init() IPaymentMethod {
	d.SetBalance(100)
	d.SetExpire("2030-12-30 23:59:59")
	return d
}

type PaymentFactory struct{}

func (pF *PaymentFactory) GetPayment(t PaymentType) (IPaymentMethod, error) {
	switch t {
	case Cash:
		return new(CashPayment).Init(), nil
	case DebitCard:
		return new(DebitCardPayment).Init(), nil
	}
	return nil, fmt.Errorf("Wrong payment type passed")
}
