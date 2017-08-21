package avida_test

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"golang.org/x/text/currency"

	"github.com/aodin/date"
	avida "github.com/tim-online/go-avida"
	"github.com/tim-online/go-avida/asix"
)

func TestDing(t *testing.T) {
	var err error

	// get username & password
	username := os.Getenv("AVIDA_USERNAME")
	password := os.Getenv("AVIDA_PASSWORD")

	// build client
	client := avida.NewClient(nil, username, password)
	client.SetDebug(true)
	client.SetBaseURL(avida.BaseURLTest)

	// ledger := asix.NewLedger()
	ledger := &asix.Ledger{
		ProductionDate: asix.RequiredDate{date.FromTime(time.Now())},
		Invoices: asix.Invoices{
			asix.Invoice{
				ClientNo:           "10223",
				CustLegalNo:        "",
				CustNo:             "893b3784-2651-4d37-9",
				Name:               "Group One",
				Name2:              "",
				Adress:             "",
				Adress2:            "",
				PostCode:           "",
				City:               "",
				CountryCode:        "",
				InvoiceNo:          "2231000005",
				InvoiceDate:        asix.RequiredDate{date.New(2017, 8, 9)},
				InvoiceDueDate:     asix.RequiredDate{date.New(2017, 8, 9)},
				Amount:             2400.0,
				VATAmount:          218,
				Currency:           asix.CurrencyType{currency.NOK},
				OrderDate:          asix.RequiredDate{date.FromTime(time.Now())},
				DeliveryDate:       asix.RequiredDate{date.FromTime(time.Time{})},
				PaymentRefNo:       "",
				OrderNo:            "",
				PackageNo:          "",
				PartialPaymentCode: "",
				PostProcessingCode: "",
				InvoiceAccount:     asix.InvoiceAccount{},
				Lines: asix.Lines{
					asix.Line{
						ItemNo:            "", // GUID xor'en
						Description:       "Night 10.08.2017",
						Description2:      "",
						UnitOfMeasure:     "",
						Quantity:          1.0, // leeg laten
						UnitPrice:         1091.0,
						VATPct:            10.0,
						VATAmount:         109.0,
						DiscountPct:       0.0,
						LineAmountExclVAT: 1091.0,
						LineAmountInclVAT: 1200.0,
					},
					asix.Line{
						ItemNo:            "", // GUID xor'en
						Description:       "Night 11.08.2017",
						Description2:      "",
						UnitOfMeasure:     "",
						Quantity:          1.0, // leeg laten
						UnitPrice:         1091.0,
						VATPct:            10.0,
						VATAmount:         109.0,
						DiscountPct:       0.0,
						LineAmountExclVAT: 1091.0,
						LineAmountInclVAT: 1200.0,
					},
				},
				Email:   "",
				PhoneNo: "",
			},
		},
	}
	request := client.Invoice.NewUploadInvoicesAixRequest()
	request.IsProdMode = true

	b, err := xml.MarshalIndent(ledger, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	// request.FileContent, err = xml.Marshal(ledger)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// response, err := client.Invoice.UploadInvoicesAix(nil, request)
	// log.Println(response)
	// log.Println(err)
}
