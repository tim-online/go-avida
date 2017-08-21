package asix

import (
	"encoding/xml"

	"github.com/aodin/date"
	"golang.org/x/text/currency"
)

type Amount float64

type UnitAmount float64

type Code string

type Code10 string

type Code20 string

type Code30 string

type CountryCode string

// type CountryCode string
// 	language.Region
// }

type CreditRefType int

type CurrencyType struct {
	currency.Unit
}

func (c CurrencyType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(c.String(), start)
}

type InvoiceActionType int

type NonEmptyCode20 string

type Percentage float64

type PositiveAmount float64

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

type Date struct {
	date.Date
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsZero() {
		return nil
	}
	return e.EncodeElement(d.String(), start)
}

type Text10 string

type Text20 string

type Text30 string

type Text50 string

type Text80 string

type Text150 string

type DeliveryOptionType int

type SaveAttribType int

type AccountTypeEnum int

type Decimal float64
