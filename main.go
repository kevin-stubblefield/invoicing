package main

import (
	"fmt"
	"invoicing/models"
	"log"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://postgres:glitter@localhost:5432/invoicing?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}
	invoices, err := env.db.FetchInvoices()
	if err != nil {
		fmt.Printf("Error: %b\n", err)
	}

	fmt.Println("Invoices")
	for _, invoice := range invoices {
		fmt.Println(invoice)
	}
}
