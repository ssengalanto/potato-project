package command

import (
	"time"

	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
)

type CreateAccountCommand struct {
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func (c *CreateAccountCommand) Topic() string {
	return CommandCreateAccount
}

func NewCreateAccountCommand(input dto.CreateAccountRequestDto) *CreateAccountCommand {
	return &CreateAccountCommand{
		Email:       input.Email,
		Password:    input.Password,
		Active:      input.Active,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		DateOfBirth: input.DateOfBirth,
	}
}
