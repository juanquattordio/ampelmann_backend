package reports

import (
	"time"
)

type ResponseFacturacionBetweenDates struct {
	DateFrom    time.Time `json:"date_from"`
	DateTo      time.Time `json:"date_to"`
	TotalAmount float64   `json:"total_amount"`
}

func NewResponseFacturacionBetweenDates(totalAmount float64, dateFrom, dateTo time.Time) *ResponseFacturacionBetweenDates {
	return &ResponseFacturacionBetweenDates{
		DateFrom:    dateFrom,
		DateTo:      dateTo,
		TotalAmount: totalAmount,
	}
}
