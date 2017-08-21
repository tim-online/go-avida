package avida

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/tim-online/go-avida/soap"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews/" + libraryVersion
	// mediaType      = "application/soap+xml"
	mediaType = "text/xml"
	charset   = "utf-8"
	// xmlns          = "https://test-srv.avidafinance.com/InvoiceInformation1/InvoiceService.svc?wsdl=wsdl0"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "srv.avidafinance.com",
		Path:   "/InvoiceInformation1/InvoiceService.svc",
	}
	BaseURLTest = url.URL{
		Scheme: "https",
		Host:   "test-srv.avidafinance.com",
		Path:   "/InvoiceInformation1/InvoiceService.svc/basic",
	}
)

// NewClient returns a new MEWS API client
func NewClient(httpClient *http.Client, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		http: httpClient,
	}
	c.SetUsername(username)
	c.SetPassword(password)
	c.SetBaseURL(BaseURL)
	c.SetUserAgent(userAgent)
	c.SetMediaType(mediaType)
	c.SetCharset(charset)

	// Services
	c.Invoice = NewInvoiceService(c)

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	// Credentials
	username string
	password string
	debug    bool
	baseURL  url.URL

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Services used for communicating with the API
	Invoice *InvoiceService
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) Auth() Auth {
	return Auth{
		Username: c.username,
		Password: c.password,
	}
}

func (c *Client) NewRequest(ctx context.Context, soapRequest *soap.Request) (*http.Request, error) {
	// convert body struct to xml
	buf := new(bytes.Buffer)
	if soapRequest != nil {
		// Add xml declaration
		buf.Write([]byte(xml.Header))
		err := xml.NewEncoder(buf).Encode(soapRequest.Envelope())
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest("POST", c.baseURL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())
	req.Header.Add("SOAPAction", soapRequest.Action().String())

	return req, nil
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, soapResponse *soap.Response) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	err = xml.NewDecoder(httpResp.Body).Decode(soapResponse.Envelope())
	if err != nil {
		errorResponse := &ErrorResponse{Response: httpResp}
		errorResponse.Message = err.Error()
		return httpResp, errorResponse
	}

	return httpResp, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	err := checkContentType(r)
	if err != nil {
		errorResponse.Message = err.Error()
	}

	if r.Header.Get("Content-Length") == "0" {
		errorResponse.Message = r.Status
		return errorResponse
	}

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = nopCloser{bytes.NewReader(data)}
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert xml to struct
	err = xml.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Message = fmt.Sprintf("Malformed xml response")
		return errorResponse
	}

	if errorResponse.Message != "" {
		return errorResponse
	}

	return nil
}

// An ErrorResponse reports the error caused by an API request
// <?xml version="1.0" encoding="UTF-8"?>
// <SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">
//   <SOAP-ENV:Body>
//     <SOAP-ENV:Fault>
//       <faultcode>Sender</faultcode>
//       <faultstring>Invalid XML</faultstring>
//     </SOAP-ENV:Fault>
//   </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>type ErrorResponse struct {
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Fault code
	Code string `xml:"Body>Fault>faultcode"`

	// Fault message
	Message string `xml:"Body>Fault>faultstring"`

	// Reason
	Reason string `xml:"Body>Fault>Reason>Text"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d (%v %v)",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Reason)
}

func checkContentType(response *http.Response) error {
	// check content-type (application/soap+xml; charset=utf-8)
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != "text/xml" {
		return fmt.Errorf("Expected Content-Type \"text/xml\", got \"%s\"", contentType)
	}

	return nil
}

type Auth struct {
	Username string `xml:"Username"`
	Password string `xml:"Password"`
}

func (a Auth) Empty() bool {
	if a.Username == "" && a.Password == "" {
		return true
	}

	return false
}
