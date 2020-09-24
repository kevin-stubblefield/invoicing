package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	Id        uuid.UUID `json:"id"`
	Lines     []Line    `json:"lines"`
	Purchaser string    `json:"purchaser"`
	CreatedAt time.Time `json:"invoice_timestamp"`
}

type Line struct {
	Id         int32 `json:"id"`
	LineNumber int32 `json:"line_number"`
	Quantity   int32 `json:"quantity"`
	Item       Item  `json:"item"`
	InvoiceId  uuid.UUID
}

type Item struct {
	Id           int32
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Type         string  `json:"type"`
	UnitPrice    int32   `json:"unit_price"`
	Discount     float32 `json:"discount"`
	DiscountType string  `json:"discount_type"`
}

func (db DB) FetchInvoices() ([]*Invoice, error) {
	rows, err := db.Query("select row_to_json(row) from ( select * from invoice_lines_json ) row;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := make([]*Invoice, 0)
	for rows.Next() {
		invoice := new(Invoice)
		var result []byte
		err := rows.Scan(&result)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(result, &invoice)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
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

	invoice := Invoice{uuid.New(), newLines, "Kevin", time.Now()}

	return invoice
}

func (invoice Invoice) String() string {
	invoiceString := fmt.Sprintf("Invoice %v Purchaser: %v Date: %v\n-------------------------------------------------------------", invoice.Id, invoice.Purchaser, invoice.CreatedAt)
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
