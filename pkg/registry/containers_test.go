package registry

import (
	"testing"

	"github.com/WendelHime/transaction-routine/pkg/usecases"
)

func TestNewAccountCreatorContainer(t *testing.T) {

	container, err := NewAccountCreatorContainer(
		5432,
		"routines",
		"127.0.0.1",
		"docker",
		"docker")
	if err != nil {
		t.Fatal(err)
	}

	defer container.Clean()
	// getting the demand creator
	uc := container.Resolve("accountCreator")

	if _, ok := uc.(usecases.UCAccountCreator); !ok {
		t.Errorf("wrong instance returned, received %T", uc)
	}

}

func TestNewAccountGetterContainer(t *testing.T) {

	container, err := NewAccountGetterContainer(
		5432,
		"routines",
		"127.0.0.1",
		"docker",
		"docker")
	if err != nil {
		t.Fatal(err)
	}

	defer container.Clean()
	// getting the demand creator
	uc := container.Resolve("accountGetter")

	if _, ok := uc.(usecases.UCAccountGetter); !ok {
		t.Errorf("wrong instance returned, received %T", uc)
	}

}

func TestNewTransactionCreatorContainer(t *testing.T) {

	container, err := NewTransactionCreatorContainer(
		5432,
		"routines",
		"127.0.0.1",
		"docker",
		"docker")
	if err != nil {
		t.Fatal(err)
	}

	defer container.Clean()
	// getting the demand creator
	uc := container.Resolve("transactionCreator")

	if _, ok := uc.(usecases.UCTransactionCreator); !ok {
		t.Errorf("wrong instance returned, received %T", uc)
	}

}
