package asix

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math"

	"github.com/aodin/date"
	"golang.org/x/text/currency"
)

// Amount with (maximum) 2 digits
// Ex. 100.48 0.1 48
type Amount float64

func (a Amount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var formatted string
	if math.Mod(float64(a), 1.0) == 0 {
		formatted = fmt.Sprintf("%.1f", a)
	} else {
		formatted = fmt.Sprintf("%g", a)
	}
	return e.EncodeElement(formatted, start)
}

// Amount with (maximum) 5 digits
// Ex. 100.4805 0.1 48
type UnitAmount float64

func (a UnitAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var formatted string
	if math.Mod(float64(a), 1.0) == 0 {
		formatted = fmt.Sprintf("%.1f", a)
	} else {
		formatted = fmt.Sprintf("%g", a)
	}
	return e.EncodeElement(formatted, start)
}

// String with upper characters.
// Valid characters: A-Z, 0-9, ÅÄÖ, -_
type Code string

// String with upper characters.
// Valid characters: A-Z, 0-9, ÅÄÖ, -_
// Max Length 10 characters
type Code10 string

// String with upper characters.
// Valid characters: A-Z, 0-9, ÅÄÖ, -_
// Max Length 20 characters
type Code20 string

// String with upper characters.
// Valid characters: A-Z, 0-9, ÅÄÖ, -_
// Max Length 30 characters
type Code30 string

// ISO 3166 Formatted Country Code
// (http://www.iso.org/iso/country_codes/iso_3166_code_lists.htm)
type CountryCode string

// type CountryCode string
// 	language.Region
// }

// Reference type of CreditNote.
// Valid is 0-3.
// 0 – Non connected credit
// 1 - References a previous InvoiceNo
// 2 - References a previous OrderNo
// 3 - References a previous PackageNo
type CreditRefType int

// Currency Used in Invoice/CreditNote.
// Valid format is 3 uppercase characters.
// Should follow ISO 4217 (http://wikipedia.org/wiki/ISO_4217)
type CurrencyType struct {
	currency.Unit
}

func (c CurrencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *CurrencyType) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	c.Unit, err = currency.ParseISO(value)
	return err
}

func (c CurrencyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(c.String(), start)
}

func (c *CurrencyType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	c.Unit, err = currency.ParseISO(value)
	return err
}

// Not in use today.
type InvoiceActionType int

// Same as Code20. Cannot be empty.
type NonEmptyCode20 string

// Decimal entry
type Percentage float64

// Positive Decimal
type PositiveAmount float64

func (a PositiveAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var formatted string
	if math.Mod(float64(a), 1.0) == 0 {
		formatted = fmt.Sprintf("%.1f", a)
	} else {
		formatted = fmt.Sprintf("%g", a)
	}
	return e.EncodeElement(formatted, start)
}

// Date in format YYYY-MM-DD. After 1800-01-01.
// Ex. 2010-02-24
type RequiredDate struct {
	date.Date
}

func (d RequiredDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsZero() {
		return nil
	}
	return e.EncodeElement(d.String(), start)
}

func (d RequiredDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	// b, err := xml.Marshal(d)
	return xml.Attr{name, d.String()}, nil
}

// Same as RequiredDate, but blank values are allowed
type Date struct {
	date.Date
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsZero() {
		return nil
	}
	return e.EncodeElement(d.String(), start)
}

// Textstring with max length 10 characters.
type Text10 string

// Textstring with max length 20 characters.
type Text20 string

// Textstring with max length 30 characters.
type Text30 string

// Textstring with max length 50 characters.
type Text50 string

// Textstring with max length 80 characters.
type Text80 string

// Textstring with max length 150 characters.
type Text150 string

// How a document should be sent
// 0 – no preference / use system setting
// 1 – Email preferred (if available)
// 2 – Mail Prioritaire (“A-Post” in Swedish)
// 3 – Mail (normal) (“B-Post” in Swedish)
// 4 – E-Invoice
// 5 – Autogiro
type DeliveryOptionType int

// Save the ‘DeliveryOptionType’ to the system or not
// 0 – Don’t save (default)
// 1 – Save as preferred value for future documents
type SaveAttribType int

// Type of invoice account
// 1 – Unique Account
// 2 – Shared Account
type AccountTypeEnum int

type Decimal float64

func (d Decimal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var formatted string
	if math.Mod(float64(d), 1.0) == 0 {
		formatted = fmt.Sprintf("%.1f", d)
	} else {
		formatted = fmt.Sprintf("%g", d)
	}
	return e.EncodeElement(formatted, start)
}
