package main

import (
	"fmt"
	"invoicing/models"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB()
	if err != nil {
		fmt.Printf("Error: %b\n", err)
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
