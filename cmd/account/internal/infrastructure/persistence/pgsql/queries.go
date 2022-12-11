package pgsql

import "fmt"

// List of valid keys for accountQueries map.
const (
	QueryExists                = "exists"
	QueryCreateAccount         = "createAccount"
	QueryCreatePerson          = "createPerson"
	QueryCreateAddress         = "createAddress"
	QueryFindAccountByID       = "findAccountById"
	QueryFindPersonByAccountID = "findPersonByAccountID"
	QueryFindAddressByPersonID = "findAddressByPersonID"
	QueryFindAccountByEmail    = "findAccountByEmail"
	QueryUpdateAccountByID     = "updateAccountByID"
	QueryUpdatePersonByID      = "updatePersonByID"
	QueryUpdateAddressByID     = "updateAddressByID"
	QueryDeleteAccountByID     = "deleteAccountByID"
)

// AccountQueries is a map holds all queries for account table.
var accountQueries = map[string]string{ //nolint:gochecknoglobals //intended
	QueryExists:                accountExistsQuery,
	QueryCreateAccount:         createAccountQuery,
	QueryCreatePerson:          createPersonQuery,
	QueryCreateAddress:         createAddressQuery,
	QueryFindAccountByID:       findAccountByIDQuery,
	QueryFindPersonByAccountID: findPersonByAccountIDQuery,
	QueryFindAddressByPersonID: findAddressByPersonIDQuery,
	QueryFindAccountByEmail:    findAccountByEmailQuery,
	QueryUpdateAccountByID:     updateAccountByIDQuery,
	QueryUpdatePersonByID:      updatePersonByIDQuery,
	QueryUpdateAddressByID:     updateAddressByIDQuery,
	QueryDeleteAccountByID:     deleteAccountByIDQuery,
}

const accountExistsQuery = `
	SELECT COUNT(1)
	FROM account
	WHERE id = $1;
	`

const createAccountQuery = `
	INSERT INTO account (id, email, password, active, last_login_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, email, password, active, last_login_at;
	`

const createPersonQuery = `
	INSERT INTO person (id, account_id, first_name, last_name, email, phone, date_of_birth, avatar)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, account_id, first_name, last_name, email, phone, date_of_birth, avatar;
	`

const createAddressQuery = `
	INSERT INTO address (
		id, 
		person_id, 
		place_id, 
		address_line1, 
		address_line2, 
		city, 
		state, 
		country, 
		postal_code, 
		formatted_address, 
		lat, 
		lng
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING 
		id, 
		person_id, 
		place_id, 
		address_line1, 
		address_line2, 
		city, 
		state, 
		country, 
		postal_code, 
		formatted_address, 
		lat, 
		lng;
	`

const findAccountByIDQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE id = $1;
	`

const findPersonByAccountIDQuery = `
	SELECT id, account_id, first_name, last_name, email, phone, date_of_birth, avatar
	FROM person
	WHERE account_id = $1;
	`

const findAddressByPersonIDQuery = `
	SELECT 
		id, 
		person_id, 
		place_id, 
		address_line1, 
		address_line2, 
		city, 
		state, 
		country, 
		postal_code, 
		formatted_address, 
		lat, 
		lng
	FROM address
	WHERE person_id = $1;
	`

const findAccountByEmailQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE email = $1;
	`

const updateAccountByIDQuery = `
	UPDATE account
	SET email = $2, password = $3, active = $4, last_login_at = $5, updated_at = NOW()
	WHERE id = $1
	RETURNING id, email, password, active, last_login_at;
	`

const updatePersonByIDQuery = `
	UPDATE person
	SET first_name = $2, last_name = $3, email = $4, phone = $5, date_of_birth = $6, avatar = $7, updated_at = NOW()
	WHERE id = $1
	RETURNING id, account_id, first_name, last_name, email, phone, date_of_birth, avatar;
	`

const updateAddressByIDQuery = `
	UPDATE address
	SET 
		place_id = $2,
		address_line1 = $3,
		address_line2 = $4,
		city = $5,
		state = $6,
		country = $7,
		postal_code = $8,
		formatted_address = $9,
		lat = $10,
		lng = $11,
		updated_at = NOW()
	WHERE id = $1
	RETURNING 
		id, 
		person_id, 
		place_id, 
		address_line1, 
		address_line2, 
		city, 
		state, 
		country, 
		postal_code, 
		formatted_address, 
		lat, 
		lng;
	`

const deleteAccountByIDQuery = `
	DELETE FROM account WHERE id = $1
	`

// MustBeValidAccountQuery accepts a string parameter that will be used
// as a key to accountQueries map. If the key doesn't exist it will
// throw a panic, otherwise it will return the query.
func MustBeValidAccountQuery(s string) string {
	query, ok := accountQueries[s]
	if !ok {
		panic(fmt.Errorf("%w: `%s` doesn't exists in account queries", ErrInvalidQuery, s))
	}

	return query
}
