package account_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/stretchr/testify/require"
)

func TestEmailIsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "valid email",
			payload: "johndoe@example.com",
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("email should be valid: %s", err)
				require.True(t, result, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "invalid email",
			payload: "invalid-email",
			assert: func(t *testing.T, result bool, err error) {
				errMsg := fmt.Sprintf("email should be valid: %s", err)
				require.False(t, result, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := account.Email(tc.payload)
			ok, err := email.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestEmailUpdate(t *testing.T) {
	testCases := []struct {
		name    string
		current string
		update  string
		assert  func(t *testing.T, expected account.Email, actual account.Email, err error)
	}{
		{
			name:    "change email success",
			current: "johndoe@example.com",
			update:  "janedoe@example.com",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("email change should succeed: %s", err)
				require.Equal(t, expected, actual, errMsg)
				require.Nil(t, err, errMsg)
			},
		},
		{
			name:    "change email failed",
			current: "johndoe@example.com",
			update:  "invalid-email",
			assert: func(t *testing.T, expected account.Email, actual account.Email, err error) {
				errMsg := fmt.Sprintf("email change should fail: %s", err)
				require.NotEqual(t, expected, actual, errMsg)
				require.NotNil(t, err, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email := account.Email(tc.current)
			newEmail, err := email.Update(tc.update)
			tc.assert(t, account.Email(tc.update), newEmail, err)
		})
	}
}

func TestEmailString(t *testing.T) {
	email := account.Email("johndoe@example.com").String()
	kind := reflect.TypeOf(email).String()
	require.Equal(t, "string", kind, "type should be `string`")
}
