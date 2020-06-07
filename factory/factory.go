package factory

import (
	"errors"
	"fmt"
)

// 工厂模式
const (
	Credit = 1
	Cash   = 2
)

type PaymentMethod interface {
	Pay(money float32) string
}

type CreditPM struct{}
type CashPM struct{}

func (c *CreditPM) Pay(money float32) string {
	return fmt.Sprintf("%0.2f Paid by CreditCard", money)
}

func (c *CashPM) Pay(money float32) string {
	return fmt.Sprintf("%#0.2f Paid by Cash", money)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Credit:
		return new(CreditPM), nil
	case Cash:
		return new(CashPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalide m: %d", m))
	}
}
