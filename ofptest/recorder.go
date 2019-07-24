package ofptest

import (
	"io"

	of "github.com/akostrikov/openflow"
)

// ResponseRecorder is an implementation of http.ResponseWriter that
// records its mutations for later inspection in tests.
type ResponseRecorder struct {
	reqs []*of.Request
}

// NewRecorder returns an initialized ResponseRecorder.
func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{}
}

// Write saves the given message into the list of requests.
func (r *ResponseRecorder) Write(h *of.Header, w io.WriterTo) error {
	req := of.NewRequest(h.Type, w)
	req.Header = *h

	r.reqs = append(r.reqs, req)
	return nil
}

// notempty thows a panic if the response list is empty.
func (r *ResponseRecorder) notempty() {
	if r == nil {
		panic("ofptest: response list is empty")
	}
}

// First returns the first response message generated by
// the handler.
func (r *ResponseRecorder) First() *of.Request {
	r.notempty()
	return r.reqs[0]
}

// Last returns the last response message generated by
// the handler.
func (r *ResponseRecorder) Last() *of.Request {
	r.notempty()
	return r.reqs[len(r.reqs)-1]
}

// All returns complete list of the messages generated
// by the handler.
func (r *ResponseRecorder) All() []*of.Request {
	return r.reqs[:]
}
