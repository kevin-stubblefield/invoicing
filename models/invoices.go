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
	Item       Item
	InvoiceId  uuid.UUID
}

type Item struct {
	Id           int32
	Name         string
	Description  string
	Type         string
	UnitPrice    int32
	Discount     float32
	DiscountType string
}

func (db DB) FetchInvoices() ([]Invoice, error) {
	return db.invoices, nil
}

func NewInvoice() Invoice {
	newId := uuid.New()

	newItems := []Item{
		{1, "Beans", "It's just beans!", "Food", 100, 50, "CASH"},
		{2, "Chicken", "Whole chicken", "Food", 649, 15, "PERCENT"},
		{3, "Hourly Labor", "Work completed", "Labor", 3500, 0, "NONE"},
	}

	newLines := []Line{
		{1, 1, 4, newItems[0], newId},
		{2, 2, 1, newItems[1], newId},
		{3, 3, 8, newItems[2], newId},
	}

	invoice := Invoice{uuid.New(), newLines}

	return invoice
}

func (invoice Invoice) String() string {
	invoiceString := fmt.Sprintf("Invoice %v\n-------------------------------------------------------------", invoice.Id)
	invoiceString += fmt.Sprintf("\n%-5s %-5s %-8s %-12s %-10s %-10s %-10s %-8s", "Id", "Line", "Quantity", "Item", "Unit Price", "Discount", "Disc Type", "Total")

	for _, line := range invoice.Lines {
		invoiceString += fmt.Sprintf("\n%5d %5d %8d %-12s $%10.2f %10.2f %-10s $%8.2f", line.Id, line.LineNumber, line.Quantity, line.Item.Name, float32(line.Item.UnitPrice)/100, printDiscount(line.Item.Discount, line.Item.DiscountType), line.Item.DiscountType, calculateTotal(line))
	}

	return invoiceString
}

func calculateTotal(line Line) float32 {
	lineTotal := float32(line.Quantity * line.Item.UnitPrice)

	fmt.Printf("Line Total: $%v, Discount: $%v\n", lineTotal, lineTotal*line.Item.Discount/100)

	if line.Item.DiscountType == "PERCENT" && line.Item.Discount > 0 {
		lineTotal -= lineTotal * line.Item.Discount / 100
	} else if line.Item.DiscountType == "CASH" && line.Item.Discount > 0 {
		lineTotal -= float32(line.Quantity) * line.Item.Discount
	}

	return lineTotal / 100
}

func printDiscount(discount float32, discountType string) float32 {
	if discountType == "PERCENT" {
		return discount
	} else {
		return discount / 100
	}
}
