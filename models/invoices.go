package models

import (
	"fmt"

	"github.com/google/uuid"
)

type Invoice struct {
	Id    uuid.UUID
	Lines []Line
}

type Line struct {
	Id         int32
	LineNumber int32
	Quantity   int32
	LineTotal  int32
	Item       Item
}

type Item struct {
	Id          int32
	Name        string
	Description string
	Type        string
}

func NewInvoice() Invoice {
	invoice := Invoice{}
	invoice.Id = uuid.New()
	return invoice
}

func (invoice Invoice) String() string {
	return fmt.Sprintf("Invoice %v", invoice.Id)
}

func (db DB) FetchInvoices() ([]Invoice, error) {
	return db.invoices, nil
}
