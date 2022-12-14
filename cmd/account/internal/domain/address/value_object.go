package address

import (
	"github.com/ssengalanto/biscuit/pkg/validator"
)

// Components contains the core fields for address entity.
type Components struct {
	Street     string `json:"street" validate:"required"`
	Unit       string `json:"unit,omitempty"`
	City       string `json:"city" validate:"required"`
	District   string `json:"district" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// IsValid checks the validity of the person details.
func (c Components) IsValid() error {
	err := validator.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

// Update checks the validity of the address components and updates its value.
func (c Components) Update(input Components) (Components, error) {
	err := input.IsValid()
	if err != nil {
		return Components{}, err
	}

	return input, nil
}
