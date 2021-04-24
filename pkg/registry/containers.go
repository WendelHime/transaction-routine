// Package registry implements the dependency injection for
// all use cases. If you want to use a usecase you'll have
// to create a container and a enum for the specified usecase
package registry

import (
	"github.com/WendelHime/transaction-routine/pkg/domain/services"
	"github.com/WendelHime/transaction-routine/pkg/gateway/persistence/postgres"
	"github.com/WendelHime/transaction-routine/pkg/usecases"
	"github.com/sarulabs/di"
)

// Container for dependency injection
type Container struct {
	ctn di.Container
}

// Resolve get a container based on name
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean destroy all containers
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func newContainer(defs []di.Def) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	err = builder.Add(defs...)

	if err != nil {
		return nil, err
	}

	return &Container{ctn: builder.Build()}, nil
}

// NewAccountCreatorContainer builds a container for account creator dependencies
func NewAccountCreatorContainer(
	dbport int,
	dbname,
	dbhost,
	dbpass,
	dbuser string) (*Container, error) {

	defs := []di.Def{
		{
			Name: "pgConnection",
			Build: func(ctn di.Container) (interface{}, error) {
				return postgres.NewConnection(dbhost, dbport, dbuser, dbpass, dbname), nil
			},
		},
		{
			Name: "accountCreator",
			Build: func(ctn di.Container) (interface{}, error) {
				conn := ctn.Get("pgConnection").(*postgres.Connection)
				return usecases.NewUCAccountCreator(services.NewAccountService(postgres.NewAccountCreator(conn), nil, nil)), nil
			},
		},
	}
	return newContainer(defs)
}

// NewAccountGetterContainer builds a container for account getter dependencies
func NewAccountGetterContainer(
	dbport int,
	dbname,
	dbhost,
	dbpass,
	dbuser string) (*Container, error) {

	defs := []di.Def{
		{
			Name: "pgConnection",
			Build: func(ctn di.Container) (interface{}, error) {
				return postgres.NewConnection(dbhost, dbport, dbuser, dbpass, dbname), nil
			},
		},
		{
			Name: "accountGetter",
			Build: func(ctn di.Container) (interface{}, error) {
				conn := ctn.Get("pgConnection").(*postgres.Connection)
				return usecases.NewUCAccountGetter(services.NewAccountService(nil, postgres.NewAccountGetter(conn), nil)), nil
			},
		},
	}
	return newContainer(defs)
}

// NewTransactionCreatorContainer builds a container for transaction creator dependencies
func NewTransactionCreatorContainer(
	dbport int,
	dbname,
	dbhost,
	dbpass,
	dbuser string) (*Container, error) {

	defs := []di.Def{
		{
			Name: "pgConnection",
			Build: func(ctn di.Container) (interface{}, error) {
				return postgres.NewConnection(dbhost, dbport, dbuser, dbpass, dbname), nil
			},
		},
		{
			Name: "transactionCreator",
			Build: func(ctn di.Container) (interface{}, error) {
				conn := ctn.Get("pgConnection").(*postgres.Connection)
				return usecases.NewUCTransactionCreator(services.NewTransactionService(postgres.NewTransactionCreator(conn)), services.NewAccountService(nil, postgres.NewAccountGetter(conn), postgres.NewAccountCreditLimitUpdater(conn))), nil
			},
		},
	}
	return newContainer(defs)
}
