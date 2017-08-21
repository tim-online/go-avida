package soap

func NewResponse() *Response {
	return &Response{
		envelope: *NewEnvelope(),
	}
}

type Response struct {
	envelope Envelope
}

func (r *Response) Envelope() Envelope {
	return r.envelope
}

func (r *Response) WithData(data interface{}) *Response {
	r.envelope.Body.Data = data
	return r
}
