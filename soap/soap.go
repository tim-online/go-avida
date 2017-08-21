package soap

import (
	"encoding/xml"
)

func NewEnvelope() *Envelope {
	return &Envelope{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
		// XMLNS:  "http://www.w3.org/2003/05/soap-envelope soap:Envelope",
		Header: NewHeader(),
		Body:   NewBody(),
	}
}

// http://stackoverflow.com/questions/16202170/marshalling-xml-go-xmlname-xmlns
type Envelope struct {
	// XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	// XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope soap:Envelope"`
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`

	Header *Header `xml:"soap:Header"`
	Body   *Body   `xml:"soap:Body"`
}

func NewHeader() *Header {
	return &Header{
		Data: nil,
	}
}

type Header struct {
	Data interface{}
}

type Body struct {
	// If the XML element contains a sub-element that hasn't matched any
	// of the above rules and the struct has a field with tag ",any",
	// unmarshal maps the sub-element to that struct field.
	Data interface{} `xml:",any"`
}

func NewBody() *Body {
	return &Body{}
}
