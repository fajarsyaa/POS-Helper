package transaction

import "time"

type TransactionFormatter struct {
	OrderID     string    `json:"order_id"`
	ExpiredTime time.Time `json:"expired_time"`
}

func FormatterTRXResponse(id string, exp time.Time) TransactionFormatter {
	return TransactionFormatter{OrderID: id, ExpiredTime: exp}
}
