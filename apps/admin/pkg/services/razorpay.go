package services

import (
	"github.com/razorpay/razorpay-go"
)

// RazorpayService handles Razorpay-related operations
type RazorpayService struct {
	client *razorpay.Client
}

// NewRazorpayService creates a new RazorpayService instance
func NewRazorpayService(key, secret string) *RazorpayService {
	client := razorpay.NewClient(key, secret)
	return &RazorpayService{client: client}
}

// CreateOrder creates a new order in Razorpay
func (s *RazorpayService) CreateOrder(amount int, currency string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"amount":   amount,
		"currency": currency,
	}
	order, err := s.client.Order.Create(data, nil)
	return order, err
}

// FetchPayment retrieves payment details
func (s *RazorpayService) FetchPayment(paymentID string) (map[string]interface{}, error) {
	return s.client.Payment.Fetch(paymentID, nil, nil)
}

func (s *RazorpayService) RefundPayment(paymentID string, amount int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"amount": amount,
	}
	return s.client.Payment.Refund(paymentID, amount, data, nil)
}
