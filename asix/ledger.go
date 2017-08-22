package asix

import (
	"encoding/xml"
)

type Ledger struct {
	XMLName xml.Name `xml:"Ledger,omitempty"`

	ProductionDate RequiredDate `xml:"productionDate,attr,omitempty"`
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

	// (not used at this time)
	// Action             interface{}    `xml:"Action,omitempty"`

	// Customers Legal No.
	// 1234561234
	CustLegalNo Code20 `xml:"CustLegalNo,omitempty"`

	// Customers No.
	// C00002
	CustNo NonEmptyCode20 `xml:"CustNo,omitempty"`

	// Customers Name
	// Albert Nobel
	Name Text50 `xml:"Name,omitempty"`

	// Customers Name, second line.
	// C/O Name
	Name2 Text50 `xml:"Name2,omitempty"`

	// Customers postal Adress.
	// Dynamic road 2
	Adress Text30 `xml:"Adress,omitempty"`

	// Customers postal Adress, second line.
	Adress2 Text30 `xml:"Adress2,omitempty"`

	// Customers Postal code.
	// 54130
	PostCode Text20 `xml:"PostCode,omitempty"`

	// Customers postal city.
	// Skövde
	City Text30 `xml:"City,omitempty"`

	// Customers Country code.
	// SE
	CountryCode CountryCode `xml:"CountryCode,omitempty"` // description: Customers Country code.

	// Invoice No.
	// F002110
	InvoiceNo NonEmptyCode20 `xml:"InvoiceNo,omitempty"`

	// Invoice Date
	// 2010-02-14
	InvoiceDate RequiredDate `xml:"InvoiceDate,omitempty"`

	// Due Date for the Invoice
	// 2010-03-14
	InvoiceDueDate RequiredDate `xml:"InvoiceDueDate,omitempty"` // description: Due Date for the Invoice

	// Total Sum of Invoice. Always positive.
	// 34521
	Amount PositiveAmount `xml:"Amount,omitempty"` // description:

	// Total VAT Amount. Always positive.
	// 8630.25
	VATAmount PositiveAmount `xml:"VATAmount,omitempty"` // description:

	// Currency of Invoice
	// SEK
	Currency CurrencyType `xml:"Currency,omitempty"` // description:

	// Date when Order was made.
	// 2010-02-10
	OrderDate RequiredDate `xml:"OrderDate,omitempty"` // description:

	// Date then Order was delivered.
	// 2010-02-14
	DeliveryDate RequiredDate `xml:"DeliveryDate,omitempty"` // description:

	// Payment reference number or code.
	// 0021108
	PaymentRefNo Code30 `xml:"PaymentRefNo,omitempty"` // description:

	// Order Number.
	// 002110
	OrderNo Code20 `xml:"OrderNo,omitempty"` // description:

	// Package / tracking number for delivery.
	PackageNo Text30 `xml:"PackageNo,omitempty"` // description:

	// Partial Payment Code, when payment is partial.
	PartialPaymentCode Code10 `xml:"PartialPaymentCode,omitempty"` // description:

	// Identifier for debt. collection processing.
	// INV02
	PostProcessingCode Code10 `xml:"PostProcessingCode,omitempty"` // description:

	// Indicates if the invoice shall be posted against an account or not
	InvoiceAccount InvoiceAccount `xml:"InvoiceAccount,omitempty"` // description:

	// See Lines
	Lines Lines `xml:"Line,omitempty"` // description:

	// Optional Fields

	// "Our Reference" on the document
	// Salesperson Xy
	OurRef Text30 `xml:"OurRef,omitempty"` // description:

	// "Your Reference" on the document
	// John Doe
	YourRef Text30 `xml:"YourRef,omitempty"` // description:

	// The Customers VAT Reg. No
	// SE554433221101
	CustVATNo Text30 `xml:"CustVATNo,omitempty"` // description:

	// The Customers Email
	// john@some.where
	Email Text80 `xml:"Email,omitempty"` // description:

	// The Customers Phone Number
	// 016 13 75 20
	PhoneNo Text30 `xml:"PhoneNo,omitempty"` // description:

	// How to deliver the document
	// 1
	DeliveryOption int `xml:"DeliveryOption,omitempty"` // description:

	// Global Location Number
	// 3322115544621
	GLN Code20 `xml:"GLN,omitempty"` // description:

	// Shipment Name on document
	ShipName Text50 `xml:"ShipName,omitempty"` // description:

	// Shipment Name, second line.
	// C/O Name
	ShipName2 Text50 `xml:"ShipName2,omitempty"` // description:

	// Shipment postal Adress.
	// Dynamic road 2
	ShipAdress Text30 `xml:"ShipAdress,omitempty"` // description:

	// Shipment postal Adress, second line.
	ShipAdress2 Text30 `xml:"ShipAdress2,omitempty"` // description:

	// Shipment Postal code.
	// 54130
	ShipPostCode Text20 `xml:"ShipPostCode,omitempty"` // description:

	// Shipment postal city.
	// Skövde
	ShipCity Text30 `xml:"ShipCity,omitempty"` // description:

	// Shipment Country code.
	// SE
	ShipCountryCode CountryCode `xml:"ShipCountryCode,omitempty"` // description:

	// A languagecode in the receiving system, or a decimal value from this list (1033 for English, 1053 for Swedish)
	// SV
	// 1053
	Language Code10 `xml:"Language,omitempty"` // description:

	// 0 = No payment request
	// 1 = Autogiro
	// 0
	PaymentRequestType int `xml:"PaymentRequestType,omitempty"` // description:

	// Id to be used with the selected payment request type.  For Autogiro this is 'betalarnr'
	// 438294
	PaymentRequestId Text30 `xml:"PaymentRequestId,omitempty"` // description:

	// Source System ID. Indicates that the client no. specified will be converted to the internal client no. used in Asitis Finance.  The Source System ID and mapping has to be entered in the application.
	// EXTERNAL
	SourceSystem Code20 `xml:"SourceSystem,omitempty"` // description:

	// Only for invoice
	Attachments Attachments `xml:"Attachment,omitempty"` // description:
}

type Credit struct {
	// Client Number in Asitis Finance, ask your contact at the Finance company.
	// 0050, C001
	ClientNo NonEmptyCode20 `xml:"ClientNo,omitempty"`

	// Customers Legal No.
	// 1234561234
	CustLegalNo Code20 `xml:"CustLegalNo,omitempty"`

	// Customers No.
	// C00002
	CustNo NonEmptyCode20 `xml:"CustNo,omitempty"`

	// Customers Name
	// Albert Nobel
	Name Text50 `xml:"Name,omitempty"`

	// Customers Name, second line.
	// C/O Name
	Name2 Text50 `xml:"Name2,omitempty"`

	// Customers postal Adress.
	// Dynamic road 2
	Adress Text30 `xml:"Adress,omitempty"`

	// Customers postal Adress, second line.
	Adress2 Text30 `xml:"Adress2,omitempty"`

	// Customers Postal code.
	// 54130
	PostCode Text20 `xml:"PostCode,omitempty"`

	// Customers postal city.
	// Skövde
	City Text30 `xml:"City,omitempty"`

	// Customers Country code.
	// SE
	CountryCode CountryCode `xml:"CountryCode,omitempty"`

	// Credit Note No.
	// F002110
	CreditNo NonEmptyCode20 `xml:"CreditNo,omitempty"`

	// Credit Note Date
	// 2010-02-14
	CreditDate RequiredDate `xml:"CreditDate,omitempty"`

	// Due Date for the Credit
	// 2010-03-14
	CreditDueDate RequiredDate `xml:"CreditDueDate,omitempty"`

	// Ref type
	// 0
	CreditRefType CreditRefType `xml:"CreditRefType,omitempty"`

	// Ref Identifier
	// 12354
	CreditRefNo Code30 `xml:"CreditRefNo,omitempty"`

	// Total Sum of Invoice. Always positive.
	// 34521
	Amount PositiveAmount `xml:"Amount,omitempty"`

	// Total VAT Amount. Always positive.
	// 8630.25
	VATAmount PositiveAmount `xml:"VATAmount,omitempty"`

	// Currency of Invoice
	// SEK
	Currency CurrencyType `xml:"Currency,omitempty"`

	// Date when Order was made.
	// 2010-02-10
	OrderDate RequiredDate `xml:"OrderDate,omitempty"`

	// Date then Order was delivered.
	// 2010-02-14
	DeliveryDate RequiredDate `xml:"RequiredDate,omitempty"`

	// Payment reference number or code.
	PaymentRefNo Code30 `xml:"PaymentRefNo,omitempty"`

	// Order Number.
	// 002110
	OrderNo Code20 `xml:"OrderNo,omitempty"`

	// Package / tracking number for delivery.
	PackageNo Text30 `xml:"PackageNo,omitempty"`

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
	VATAmount Amount `xml:"VATAmount,omitempty"`

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

type Attachments []Attachment

type Attachment struct {
	// Namespace for attachment
	// Company A or Client A
	Realm Text50 `xml:"Realm,omitempty"`

	// File Name
	// InvoieNo_1_Attachment1.pdf
	Name Text150 `xml:"Name,omitempty"`

	// Number of pages of attachment
	// 5
	Pages int `xml:"Pages,omitempty"`
}
