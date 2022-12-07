package pgsql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
)

type AccountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository creates a new account repository.
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Exists checks if an account record with specific ID exists in the database.
func (a *AccountRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int

	query := MustBeValidAccountQuery(QueryExists)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

// Create begins a new transaction to process and insert a new Account record together with its associated
// Person and Address records. If transaction fails it will roll back all the changes it made,
// otherwise it will commit the changes to the database.
func (a *AccountRepository) Create(ctx context.Context, entity account.Entity) (account.Entity, error) {
	e := account.Entity{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	acc, err := createAccount(ctx, tx, entity)
	if err != nil {
		return e, err
	}

	p, err := createPerson(ctx, tx, *entity.Person)
	if err != nil {
		return e, err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	e = acc
	e.Person = &p
	return e, nil
}

// FindByID gets an account record with specific ID in the database.
func (a *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// FindByEmail gets an account record with specific email in the database.
func (a *AccountRepository) FindByEmail(ctx context.Context, email string) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindByEmail)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, email)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// UpdateByID updates an account record with specific id in the database.
func (a *AccountRepository) UpdateByID(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryUpdateByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// DeleteByID deletes an account record with specific ID in the database.
func (a *AccountRepository) DeleteByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryDeleteByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// createAccount inserts a new Account record in the database.
func createAccount(ctx context.Context, tx *sqlx.Tx, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryCreateAccount)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return acc.ToEntity(), err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// createPerson inserts a new Person record associated with account in the database.
func createPerson(ctx context.Context, tx *sqlx.Tx, entity person.Entity) (person.Entity, error) {
	p := Person{}

	query := MustBeValidAccountQuery(QueryCreatePerson)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return p.ToEntity(), err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.AccountID,
		entity.Details.FirstName,
		entity.Details.LastName,
		entity.Details.Email,
		entity.Details.Phone,
		entity.Details.DateOfBirth,
		entity.Avatar,
	)

	err = row.StructScan(&p)
	if err != nil {
		return p.ToEntity(), err
	}

	return p.ToEntity(), nil
}
