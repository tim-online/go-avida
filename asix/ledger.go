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
	ClientNo NonEmptyCode20 `xml:"ClientNo,omitempty"`
	// Action
	CustLegalNo        Code20         `xml:"CustLegalNo,omitempty"`
	CustNo             NonEmptyCode20 `xml:"CustNo,omitempty"`
	Name               Text50         `xml:"Name,omitempty"`
	Name2              Text50         `xml:"Name2,omitempty"`
	Adress             Text30         `xml:"Adress,omitempty"`
	Adress2            Text30         `xml:"Adress2,omitempty"`
	PostCode           Text20         `xml:"PostCode,omitempty"`
	City               Text30         `xml:"City,omitempty"`
	CountryCode        CountryCode    `xml:"CountryCode,omitempty"`
	InvoiceNo          NonEmptyCode20 `xml:"InvoiceNo,omitempty"`
	InvoiceDate        RequiredDate   `xml:"InvoiceDate,omitempty"`
	InvoiceDueDate     RequiredDate   `xml:"InvoiceDueDate,omitempty"`
	Amount             PositiveAmount `xml:"Amount,omitempty"`
	VATAmount          PositiveAmount `xml:"VATAmount,omitempty"`
	Currency           CurrencyType   `xml:"Currency,omitempty"`
	OrderDate          RequiredDate   `xml:"OrderDate,omitempty"`
	DeliveryDate       RequiredDate   `xml:"DeliveryDate,omitempty"`
	PaymentRefNo       Code30         `xml:"PaymentRefNo,omitempty"`
	OrderNo            Code20         `xml:"OrderNo,omitempty"`
	PackageNo          Text30         `xml:"PackageNo,omitempty"`
	PartialPaymentCode Code10         `xml:"PartialPaymentCode,omitempty"`
	PostProcessingCode Code10         `xml:"PostProcessingCode,omitempty"`
	InvoiceAccount     InvoiceAccount `xml:"InvoiceAccount,omitempty"`
	Lines              Lines          `xml:"Line,omitempty"`

	// Optional Fields
	OurRef             Text30      `xml:"OurRef,omitempty"`
	YourRef            Text30      `xml:"YourRef,omitempty"`
	CustVATNo          Text30      `xml:"CustVATNo,omitempty"`
	Email              Text80      `xml:"Email,omitempty"`
	PhoneNo            Text30      `xml:"PhoneNo,omitempty"`
	DeliveryOption     int         `xml:"DeliveryOption,omitempty"`
	GLN                Code20      `xml:"GLN,omitempty"`
	ShipName           Text50      `xml:"ShipName,omitempty"`
	ShipName2          Text50      `xml:"ShipName2,omitempty"`
	ShipAdress         Text30      `xml:"ShipAdress,omitempty"`
	ShipAdress2        Text30      `xml:"ShipAdress2,omitempty"`
	ShipPostCode       Text20      `xml:"ShipPostCode,omitempty"`
	ShipCity           Text30      `xml:"ShipCity,omitempty"`
	ShipCountryCode    CountryCode `xml:"ShipCountryCode,omitempty"`
	Language           Code10      `xml:"Language,omitempty"`
	PaymentRequestType int         `xml:"PaymentRequestType,omitempty"`
	PaymentRequestId   Text30      `xml:"PaymentRequestId,omitempty"`
	SourceSystem       Code20      `xml:"SourceSystem,omitempty"`

	// Only for invoice
	Attachments Attachments `xml:"Attachment,omitempty"`
}

type Credit struct {
	ClientNo      NonEmptyCode20 `xml:"ClientNo,omitempty"`
	CustLegalNo   Code20         `xml:"CustLegalNo,omitempty"`
	CustNo        NonEmptyCode20 `xml:"CustNo,omitempty"`
	Name          Text50         `xml:"Name,omitempty"`
	Name2         Text50         `xml:"Name2,omitempty"`
	Adress        Text30         `xml:"Adress,omitempty"`
	Adress2       Text30         `xml:"Adress2,omitempty"`
	PostCode      Text20         `xml:"PostCode,omitempty"`
	City          Text30         `xml:"City,omitempty"`
	CountryCode   CountryCode    `xml:"CountryCode,omitempty"`
	CreditNo      NonEmptyCode20 `xml:"CreditNo,omitempty"`
	CreditDate    RequiredDate   `xml:"CreditDate,omitempty"`
	CreditDueDate RequiredDate   `xml:"CreditDueDate,omitempty"`
	CreditRefType CreditRefType  `xml:"CreditRefType,omitempty"`
	Amount        PositiveAmount `xml:"Amount,omitempty"`
	VATAmount     PositiveAmount `xml:"VATAmount,omitempty"`
	Currency      CurrencyType   `xml:"Currency,omitempty"`
	OrderDate     RequiredDate   `xml:"OrderDate,omitempty"`
	PaymentRefNo  Code30         `xml:"PaymentRefNo,omitempty"`
	OrderNo       Code20         `xml:"OrderNo,omitempty"`
	PackageNo     Text30         `xml:"PackageNo,omitempty"`
	Lines         Lines          `xml:"Line,omitempty"`
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
	ItemNo            Text20     `xml:"ItemNo"` // description: Item No, Article No. // example: 001220
	Description       Text80     `xml:"Description,omitempty"`
	Description2      Text80     `xml:"Description2,omitempty"`
	UnitOfMeasure     Text10     `xml:"UnitOfMeasure,omitempty"`
	Quantity          Decimal    `xml:"Quantity,omitempty"`
	UnitPrice         UnitAmount `xml:"UnitPrice,omitempty"`
	VATPct            Percentage `xml:"VATPct,omitempty"`
	VATAmount         Amount     `xml:"VATAmount,omitempty"`
	DiscountPct       Percentage `xml:"DiscountPct,omitempty"`
	LineAmountExclVAT Amount     `xml:"LineAmountExclVAT,omitempty"`
	LineAmountInclVAT Amount     `xml:"LineAmountInclVAT,omitempty"`
}

type Attachments []Attachment

type Attachment struct {
	Name  Text50  `xml:"Name,omitempty"`
	Pages Text150 `xml:"Pages,omitempty"`
	Realm int     `xml:"Realm,omitempty"`
}
