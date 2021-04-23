package main

import (
	"log"
	"os"
	"strconv"

	"github.com/WendelHime/transaction-routine/pkg/gateway/http"
	"github.com/WendelHime/transaction-routine/pkg/registry"
	"github.com/WendelHime/transaction-routine/pkg/usecases"
)

func main() {
	dbport, err := strconv.Atoi(os.Getenv("DBPORT"))
	if err != nil {
		log.Panicf("fail to convert port to int: %v\n", err)
	}
	accountCreatorCtn, err := registry.NewAccountCreatorContainer(
		dbport,
		os.Getenv("DBNAME"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPASS"),
		os.Getenv("DBUSER"))
	if err != nil {
		panic(err)
	}
	accountGetterCtn, err := registry.NewAccountGetterContainer(
		dbport,
		os.Getenv("DBNAME"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPASS"),
		os.Getenv("DBUSER"))
	if err != nil {
		panic(err)
	}
	defer accountGetterCtn.Clean()
	transactionCreatorCtn, err := registry.NewTransactionCreatorContainer(
		dbport,
		os.Getenv("DBNAME"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPASS"),
		os.Getenv("DBUSER"))
	if err != nil {
		panic(err)
	}
	app := http.NewApp(
		accountCreatorCtn.Resolve("accountCreator").(usecases.UCAccountCreator),
		accountGetterCtn.Resolve("accountGetter").(usecases.UCAccountGetter),
		transactionCreatorCtn.Resolve("transactionCreator").(usecases.UCTransactionCreator),
	)

	app.Run(":3000")
}
