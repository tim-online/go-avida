package asix

import (
	"encoding/xml"
)

type Ledger struct {
	XMLName xml.Name `xml:"Ledger,omitempty"`

	ProductionDate RequiredDate `xml:"productionDate,attr"`
	Invoices       Invoices     `xml:"Invoice,omitempty"`
	Credits        Credits      `xml:"Credit,omitempty"`
}

type Invoices []Invoice

type Credits []Credit

func NewLedger() *Ledger {
	return &Ledger{}
}

type Invoice struct {
	// Client Number in Asitis Finance, ask your contact at the Finance company.
	ClientNo NonEmptyCode20 `xml:"ClientNo,omitempty"`

	// Customers Legal No.
	// 1234561234
	CustLegalNo Code20 `xml:"CustLegalNo"`

	// Customers No.
	// C00002
	CustNo NonEmptyCode20 `xml:"CustNo,omitempty"`

	// Customers Name
	// Albert Nobel
	Name Text50 `xml:"Name,omitempty"`

	// Customers postal Adress.
	// Dynamic road 2
	Adress Text30 `xml:"Adress"`

	// Customers postal Adress, second line.
	Adress2 Text30 `xml:"Adress2,omitempty"`

	// Customers Postal code.
	// 54130
	PostCode Text20 `xml:"PostCode"`

	// Customers postal city.
	// Skövde
	City Text30 `xml:"City"`

	// Invoice No.
	// F002110
	InvoiceNo NonEmptyCode20 `xml:"InvoiceNo,omitempty"`

	// Invoice Date
	// 2010-02-14
	InvoiceDate RequiredDate `xml:"InvoiceDate"`

	// Due Date for the Invoice
	// 2010-03-14
	InvoiceDueDate RequiredDate `xml:"InvoiceDueDate"` // description: Due Date for the Invoice

	// Total Sum of Invoice. Always positive.
	// 34521
	Amount PositiveAmount `xml:"Amount,omitempty"` // description:

	// Total VAT Amount. Always positive.
	// 8630.25
	VATAmount PositiveAmount `xml:"VATAmount"` // description:

	// Currency of Invoice
	// SEK
	Currency CurrencyType `xml:"Currency,omitempty"` // description:

	// Payment reference number or code.
	// 0021108
	PaymentRefNo Code30 `xml:"PaymentRefNo,omitempty"` // description:

	// "Our Reference" on the document
	// Salesperson Xy
	OurRef Text30 `xml:"OurRef,omitempty"` // description:

	// "Your Reference" on the document
	// John Doe
	YourRef Text30 `xml:"YourRef,omitempty"` // description:

	// The Customers Email
	// john@some.where
	Email Text80 `xml:"Email,omitempty"` // description:

	// See Lines
	Lines Lines `xml:"Line,omitempty"` // description:
}

type Credit struct {
	// Client Number in Asitis Finance, ask your contact at the Finance company.
	// 0050, C001
	ClientNo NonEmptyCode20 `xml:"ClientNo,omitempty"`

	// Customers Legal No.
	// 1234561234
	CustLegalNo Code20 `xml:"CustLegalNo"`

	// Customers No.
	// C00002
	CustNo NonEmptyCode20 `xml:"CustNo,omitempty"`

	// Customers Name
	// Albert Nobel
	Name Text50 `xml:"Name,omitempty"`

	// Customers postal Adress.
	// Dynamic road 2
	Adress Text30 `xml:"Adress"`

	// Customers postal Adress, second line.
	Adress2 Text30 `xml:"Adress2,omitempty"`

	// Customers Postal code.
	// 54130
	PostCode Text20 `xml:"PostCode"`

	// Customers postal city.
	// Skövde
	City Text30 `xml:"City"`

	// Credit Note No.
	// F002110
	CreditNo NonEmptyCode20 `xml:"CreditNo,omitempty"`

	// Credit Note Date
	// 2010-02-14
	CreditDate RequiredDate `xml:"CreditDate"`

	// Due Date for the Credit
	// 2010-03-14
	CreditDueDate RequiredDate `xml:"CreditDueDate"`

	// Ref type
	// 0
	CreditRefType CreditRefType `xml:"CreditRefType"`

	// Ref Identifier
	// 12354
	CreditRefNo Code30 `xml:"CreditRefNo"`

	// Total Sum of Invoice. Always positive.
	// 34521
	Amount PositiveAmount `xml:"Amount,omitempty"`

	// Total VAT Amount. Always positive.
	// 8630.25
	VATAmount PositiveAmount `xml:"VATAmount"`

	// Currency of Invoice
	// SEK
	Currency CurrencyType `xml:"Currency,omitempty"`

	// "Our Reference" on the document
	// Salesperson Xy
	OurRef Text30 `xml:"OurRef,omitempty"` // description:

	// "Your Reference" on the document
	// John Doe
	YourRef Text30 `xml:"YourRef,omitempty"` // description:

	// The Customers Email
	// john@some.where
	Email Text80 `xml:"Email,omitempty"` // description:

	// See Lines
	Lines Lines `xml:"Line,omitempty"`
}

type InvoiceAccount struct {
	Type              int  `xml:"type,attr"`
	InterestStartDate Date `xml:"InterestStartDate,omitempty"`
	PaymentStartDate  Date `xml:"PaymentStartDate,omitempty"`
}

func (i InvoiceAccount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if i.Type == 0 {
		return nil
	}
	return e.EncodeElement(i, start)
}

type Lines []Line

// If the ‘DiscountPct’ is used, then it should be together with ‘UnitPrice’ in all formulas. (i.e. a positive
// ‘DiscountPct’ reduces the value of the ‘UnitPrice’ that is used in the calculation. It does not reduce the
// actual ‘UnitPrice’-value)
// Any numeric field can be negative, as long as the formulas are still valid.
// The sum of all ‘LineAmountInclVAT’ must always be positive.
type Line struct {
	// Item No, Article No.
	// 001220
	ItemNo Text20 `xml:"ItemNo"`

	// Description of Item.
	// Black Pencil
	Description Text80 `xml:"Description,omitempty"`

	// Description of Item, Second line.
	// With red line.
	Description2 Text80 `xml:"Description2,omitempty"`

	// Measurement Unit
	// M, kg, Hour
	UnitOfMeasure Text10 `xml:"UnitOfMeasure,omitempty"`

	// Quantity
	// 10
	Quantity Decimal `xml:"Quantity,omitempty"`

	// The price for one item. Excluding VAT.
	// 2.9805
	UnitPrice UnitAmount `xml:"UnitPrice,omitempty"`

	// VAT Percent for item
	// 25
	VATPct Percentage `xml:"VATPct,omitempty"`

	// VAT Amount for all Items.  7.45 Quantity * UnitPrice * (VATPct/100)
	// 7.45
	VATAmount Amount `xml:"VATAmount"`

	// Discount in percent.
	// 0
	DiscountPct Percentage `xml:"DiscountPct,omitempty"`

	// Line Amount, Quantity * UnitPrice
	// 29.8
	LineAmountExclVAT Amount `xml:"LineAmountExclVAT,omitempty"`

	// LineAmountExclVAT + VATAmount
	// 37.25
	LineAmountInclVAT Amount `xml:"LineAmountInclVAT,omitempty"`
}
