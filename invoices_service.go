package avida

import (
	"context"
	"encoding/xml"
	"net/url"

	"github.com/tim-online/go-avida/soap"
)

type InvoiceService struct {
	client *Client
}

func NewInvoiceService(client *Client) *InvoiceService {
	return &InvoiceService{client: client}
}

func (s *InvoiceService) UploadInvoicesAix(ctx context.Context, requestBody *UploadInvoicesAixRequest) (*UploadInvoicesAixResponse, error) {
	action, err := url.Parse("http://avida.se/IInvoiceService/UploadInvoicesAix")
	if err != nil {
		return nil, err
	}

	// set Auth
	if requestBody.Auth.Empty() {
		requestBody.Auth = s.client.Auth()
	}
	// set body and header SOAPAction
	request := soap.NewRequest().
		WithData(requestBody).
		WithAction(*action)
	// create Response
	responseBody := s.NewUploadInvoicesAixResponse()
	response := soap.NewResponse().WithData(responseBody)

	// make Request
	httpReq, err := s.client.NewRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	// do request
	_, err = s.client.Do(httpReq, response)
	return responseBody, err
}

func (s *InvoiceService) NewUploadInvoicesAixRequest() *UploadInvoicesAixRequest {
	return &UploadInvoicesAixRequest{}
}

type UploadInvoicesAixRequest struct {
	XMLName xml.Name `xml:"http://avida.se UploadInvoicesAix"`

	Auth        Auth   `xml:"auth"`
	Email       string `xml:"email"`
	FileContent []byte `xml:"fileContent"`
	IsProdMode  bool   `xml:"isProdMode"`
}

func (s *InvoiceService) NewUploadInvoicesAixResponse() *UploadInvoicesAixResponse {
	return &UploadInvoicesAixResponse{}
}

type UploadInvoicesAixResponse struct{}
