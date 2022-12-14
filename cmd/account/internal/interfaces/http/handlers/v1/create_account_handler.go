//nolint:godot //unnecessary
package v1

import (
	"context"
	"net/http"

	cmdv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/command/v1"
	qv1 "github.com/ssengalanto/biscuit/cmd/account/internal/application/query/v1"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	apphttp "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/http"
	"github.com/ssengalanto/biscuit/pkg/constants"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

// CreateAccountHandler - http handler struct for account creation.
type CreateAccountHandler struct {
	log      interfaces.Logger
	mediator *midt.Midt
}

// NewCreateAccountHandler creates a new http handler for handling account creation.
func NewCreateAccountHandler(logger interfaces.Logger, mediator *midt.Midt) *CreateAccountHandler {
	return &CreateAccountHandler{log: logger, mediator: mediator}
}

// Handle
// @Tags account
// @Summary Create a new account
// @Description Creates a new account
// @Accept json
// @Produce json
// @Param CreateAccountRequestDto body dto.CreateAccountRequestDto true "Account data"
// @Success 201 {object} dto.GetAccountResponseDto
// @Failure 400 {object} errors.HTTPError
// @Failure 500 {object} errors.HTTPError
// @Router /api/v1/account [post]
func (c *CreateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.RequestTimeout)
	defer cancel()

	var request dto.CreateAccountRequestDto

	err := json.DecodeRequest(w, r, &request)
	if err != nil {
		c.log.Error("invalid request body format", map[string]any{"error": err})
		json.MustEncodeError(w, errors.ErrInvalid)
		return
	}

	if !apphttp.ValidateRequest(w, c.log, request) {
		return
	}

	cmd := cmdv1.NewCreateAccountCommand(dto.CreateAccountRequestDto{
		Email:       request.Email,
		Password:    request.Password,
		Active:      request.Active,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
		DateOfBirth: request.DateOfBirth,
		Locations:   request.Locations,
	})

	rr, err := c.mediator.Send(ctx, cmd)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	resource, ok := rr.(dto.CreateAccountResponseDto)
	if !ok {
		c.log.Error("invalid resource", map[string]any{"resource": resource})
		json.MustEncodeError(w, errors.ErrInvalid)
	}

	q := qv1.NewGetAccountQuery(dto.GetAccountRequestDto{ID: resource.ID}) //nolint:gosimple //explicit
	response, err := c.mediator.Send(ctx, q)
	if err != nil {
		json.MustEncodeError(w, err)
		return
	}

	json.MustEncodeResponse(w, http.StatusCreated, response)
}
