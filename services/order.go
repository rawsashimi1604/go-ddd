// the services package holds all services that relate to business logic
// Service Configuration Generator Pattern
package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/go-ddd/domain/customer"
	"github.com/rawsashimi1604/go-ddd/domain/customer/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	// Need to create ProductRepository and product aggregate.
	log.Println(c)
	return nil
}
