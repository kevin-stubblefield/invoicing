package models

type Datastore interface {
	FetchInvoices() ([]Invoice, error)
}

type DB struct {
	invoices []Invoice
}

func NewDB() (DB, error) {
	db := DB{}

	db.invoices = append(db.invoices, NewInvoice())

	return db, nil
}
