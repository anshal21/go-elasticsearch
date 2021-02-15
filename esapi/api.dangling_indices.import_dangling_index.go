// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 7.11.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newDanglingIndicesImportDanglingIndexFunc(t Transport) DanglingIndicesImportDanglingIndex {
	return func(index_uuid string, o ...func(*DanglingIndicesImportDanglingIndexRequest)) (*Response, error) {
		var r = DanglingIndicesImportDanglingIndexRequest{IndexUUID: index_uuid}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// DanglingIndicesImportDanglingIndex imports the specified dangling index
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-gateway-dangling-indices.html.
//
type DanglingIndicesImportDanglingIndex func(index_uuid string, o ...func(*DanglingIndicesImportDanglingIndexRequest)) (*Response, error)

// DanglingIndicesImportDanglingIndexRequest configures the Dangling Indices Import Dangling Index API request.
//
type DanglingIndicesImportDanglingIndexRequest struct {
	IndexUUID string

	AcceptDataLoss *bool
	MasterTimeout  time.Duration
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r DanglingIndicesImportDanglingIndexRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_dangling") + 1 + len(r.IndexUUID))
	path.WriteString("/")
	path.WriteString("_dangling")
	path.WriteString("/")
	path.WriteString(r.IndexUUID)

	params = make(map[string]string)

	if r.AcceptDataLoss != nil {
		params["accept_data_loss"] = strconv.FormatBool(*r.AcceptDataLoss)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f DanglingIndicesImportDanglingIndex) WithContext(v context.Context) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.ctx = v
	}
}

// WithAcceptDataLoss - must be set to true in order to import the dangling index.
//
func (f DanglingIndicesImportDanglingIndex) WithAcceptDataLoss(v bool) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.AcceptDataLoss = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
//
func (f DanglingIndicesImportDanglingIndex) WithMasterTimeout(v time.Duration) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f DanglingIndicesImportDanglingIndex) WithTimeout(v time.Duration) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f DanglingIndicesImportDanglingIndex) WithPretty() func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f DanglingIndicesImportDanglingIndex) WithHuman() func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f DanglingIndicesImportDanglingIndex) WithErrorTrace() func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f DanglingIndicesImportDanglingIndex) WithFilterPath(v ...string) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f DanglingIndicesImportDanglingIndex) WithHeader(h map[string]string) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f DanglingIndicesImportDanglingIndex) WithOpaqueID(s string) func(*DanglingIndicesImportDanglingIndexRequest) {
	return func(r *DanglingIndicesImportDanglingIndexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
