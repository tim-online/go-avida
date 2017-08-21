package soap

import "net/url"

func NewRequest() *Request {
	return &Request{
		envelope: *NewEnvelope(),
	}
}

type Request struct {
	envelope Envelope
	action   url.URL `xml:""`
}

func (r *Request) Envelope() Envelope {
	return r.envelope
}

func (r *Request) WithData(data interface{}) *Request {
	r.envelope.Body.Data = data
	return r
}

func (r *Request) WithAction(action url.URL) *Request {
	r.action = action
	return r
}

func (r *Request) Action() *url.URL {
	return &r.action
}
