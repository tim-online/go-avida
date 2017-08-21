package aix

import (
	"encoding/json"
	"encoding/xml"

	"github.com/aodin/date"
	"golang.org/x/text/currency"
)

// https://www.avidafinance.com/contentassets/1005/fakturafilspec.pdf

const (
	version = 3
)

type InvoiceList struct {
	XMLName xml.Name `xml:"Ledger"`

	Version int     `xml:"Version,omitempty"`
	Seller  Seller  `xml:"Seller,omitempty"`
	Invoice Invoice `xml:"Invoice,omitempty"`
}

func NewInvoiceList() *InvoiceList {
	return &InvoiceList{
		Version: version,
	}
}

type Seller struct {
	PartnerID string   `xml:"PartnerID,omitempty"`
	Company   string   `xml:"Company,omitempty"`
	VatRegNo  VatRegNo `xml:"VatRegNo,omitempty"`
}

type VatRegNo struct {
}

type Invoice struct {
	Buyer              Buyer         `xml:"Buyer,omitempty"`
	InvoiceNo          string        `xml:"InvoiceNo,omitempty"`
	InvoiceDate        date.Date     `xml:"InvoiceDate,omitempty"`
	DueDate            date.Date     `xml:"DueDate,omitempty"`
	Currency           currency.Unit `xml:"Currency,omitempty"`
	Lines              Lines         `xml:"Lines,omitempty"`
	TotalVATAmount     json.Number   `xml:"TotalVATAmount,omitempty"`
	TotalAmountInclVAT json.Number   `xml:"TotalAmountInclVAT,omitempty"`
	TotalAmountExclVAT json.Number   `xml:"TotalAmountExclVAT,omitempty"`
	TotalAmount        json.Number   `xml:"TotalAmount,omitempty"`
	Rounding           json.Number   `xml:"Rounding,omitempty"`
	Comments           Comments      `xml:"Comments,omitempty"`
}

func (i Invoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Invoice
	type export struct {
		alias
		Currency string `xml:"Currency,omitempty"`
	}

	aux := export{
		alias:    alias(i),
		Currency: i.Currency.String(),
	}
	return e.EncodeElement(aux, start)
}

type Buyer struct {
	CustomerNo  string    `xml:"CustomerNo,omitempty"`
	Company     string    `xml:"Company,omitempty"`
	BuyerID     string    `xml:"BuyerID,omitempty"`
	ContactInfo string    `xml:"ContactInfo,omitempty"`
	Address     []Address `xml:"Address,omitempty"`
}

type ContactInfo struct {
	Contact       string `xml:"Contact,omitempty"`
	YourReference string `xml:"YourReference,omitempty"`
	OurReference  string `xml:"OurReference,omitempty"`
}

type Address struct {
	Type     string `xml:"type,attr" `
	Name1    string `xml:"Name1,omitempty"`
	Address1 string `xml:"Address1,omitempty"`
	ZipCode  string `xml:"ZipCode,omitempty"`
	City     string `xml:"City,omitempty"`
	Phone1   string `xml:"Phone1,omitempty"`
	Email    string `xml:"Email,omitempty"`
}

type Lines struct {
	Line []Line `xml:"Line,omitempty"`
}

type Line struct {
	LineNo            int         `xml:"LineNo,omitempty"`
	Quantity          int         `xml:"Quantity,omitempty"`
	UnitMeasure       string      `xml:"UnitMeasure,omitempty"`
	Description       string      `xml:"Description,omitempty"`
	UnitAmount        json.Number `xml:"UnitAmount,omitempty"`
	VATPct            json.Number `xml:"VATPct,omitempty"`
	VATAmount         json.Number `xml:"VATAmount,omitempty"`
	DiscountPct       json.Number `xml:"DiscountPct,omitempty"`
	LineAmountInclVAT json.Number `xml:"LineAmountInclVAT,omitempty"`
	LineAmountExclVAT json.Number `xml:"LineAmountExclVAT,omitempty"`
}

type Comments []Comment

type Comment struct {
	Code string `xml:"Code,omitempty"`
	No   string `xml:"No,omitempty"`
	Text string `xml:"Text,omitempty"`
}
