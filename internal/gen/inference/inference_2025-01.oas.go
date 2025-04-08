// Package inference provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package inference

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for ErrorResponseErrorCode.
const (
	ABORTED            ErrorResponseErrorCode = "ABORTED"
	ALREADYEXISTS      ErrorResponseErrorCode = "ALREADY_EXISTS"
	DATALOSS           ErrorResponseErrorCode = "DATA_LOSS"
	DEADLINEEXCEEDED   ErrorResponseErrorCode = "DEADLINE_EXCEEDED"
	FAILEDPRECONDITION ErrorResponseErrorCode = "FAILED_PRECONDITION"
	FORBIDDEN          ErrorResponseErrorCode = "FORBIDDEN"
	INTERNAL           ErrorResponseErrorCode = "INTERNAL"
	INVALIDARGUMENT    ErrorResponseErrorCode = "INVALID_ARGUMENT"
	NOTFOUND           ErrorResponseErrorCode = "NOT_FOUND"
	OK                 ErrorResponseErrorCode = "OK"
	OUTOFRANGE         ErrorResponseErrorCode = "OUT_OF_RANGE"
	PERMISSIONDENIED   ErrorResponseErrorCode = "PERMISSION_DENIED"
	QUOTAEXCEEDED      ErrorResponseErrorCode = "QUOTA_EXCEEDED"
	RESOURCEEXHAUSTED  ErrorResponseErrorCode = "RESOURCE_EXHAUSTED"
	UNAUTHENTICATED    ErrorResponseErrorCode = "UNAUTHENTICATED"
	UNAVAILABLE        ErrorResponseErrorCode = "UNAVAILABLE"
	UNIMPLEMENTED      ErrorResponseErrorCode = "UNIMPLEMENTED"
	UNKNOWN            ErrorResponseErrorCode = "UNKNOWN"
)

// Defines values for VectorType.
const (
	Dense  VectorType = "dense"
	Sparse VectorType = "sparse"
)

// DenseEmbedding A dense embedding of a single input
type DenseEmbedding struct {
	// Values The dense embedding values.
	Values []float64 `json:"values"`

	// VectorType Indicates whether this is a 'dense' or 'sparse' embedding.
	VectorType VectorType `json:"vector_type"`
}

// Document Document for reranking
type Document map[string]interface{}

// EmbedRequest defines model for EmbedRequest.
type EmbedRequest struct {
	// Inputs List of inputs to generate embeddings for.
	Inputs []struct {
		Text *string `json:"text,omitempty"`
	} `json:"inputs"`

	// Model The [model](https://docs.pinecone.io/guides/inference/understanding-inference#embedding-models) to use for embedding generation.
	Model string `json:"model"`

	// Parameters Additional model-specific parameters. Refer to the [model guide](https://docs.pinecone.io/guides/inference/understanding-inference#embedding-models) for available model parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty"`
}

// Embedding Embedding of a single input
type Embedding struct {
	union json.RawMessage
}

// EmbeddingsList Embeddings generated for the input.
type EmbeddingsList struct {
	// Data The embeddings generated for the inputs.
	Data []Embedding `json:"data"`

	// Model The model used to generate the embeddings
	Model string `json:"model"`

	// Usage Usage statistics for the model inference.
	Usage struct {
		// TotalTokens Total number of tokens consumed across all inputs.
		TotalTokens *int32 `json:"total_tokens,omitempty"`
	} `json:"usage"`

	// VectorType Indicates whether the response data contains 'dense' or 'sparse' embeddings.
	VectorType string `json:"vector_type"`
}

// ErrorResponse The response shape used for all error responses.
type ErrorResponse struct {
	// Error Detailed information about the error that occurred.
	Error struct {
		Code ErrorResponseErrorCode `json:"code"`

		// Details Additional information about the error. This field is not guaranteed to be present.
		Details *map[string]interface{} `json:"details,omitempty"`
		Message string                  `json:"message"`
	} `json:"error"`

	// Status The HTTP status code of the error.
	Status int `json:"status"`
}

// ErrorResponseErrorCode defines model for ErrorResponse.Error.Code.
type ErrorResponseErrorCode string

// RankedDocument A ranked document with a relevance score and an index position.
type RankedDocument struct {
	// Document Document for reranking
	Document *Document `json:"document,omitempty"`

	// Index The index position of the document from the original request.
	Index int `json:"index"`

	// Score The relevance of the document to the query, normalized between 0 and 1, with scores closer to 1 indicating higher relevance.
	Score float64 `json:"score"`
}

// RerankRequest defines model for RerankRequest.
type RerankRequest struct {
	// Documents The documents to rerank.
	Documents []Document `json:"documents"`

	// Model The [model](https://docs.pinecone.io/guides/inference/understanding-inference#reranking-models) to use for reranking.
	Model string `json:"model"`

	// Parameters Additional model-specific parameters. Refer to the [model guide](https://docs.pinecone.io/guides/inference/understanding-inference#reranking-models) for available model parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty"`

	// Query The query to rerank documents against.
	Query string `json:"query"`

	// RankFields The field(s) to consider for reranking. If not provided, the default is `["text"]`.
	//
	// The number of fields supported is [model-specific](https://docs.pinecone.io/guides/inference/understanding-inference#reranking-models).
	RankFields *[]string `json:"rank_fields,omitempty"`

	// ReturnDocuments Whether to return the documents in the response.
	ReturnDocuments *bool `json:"return_documents,omitempty"`

	// TopN The number of results to return sorted by relevance. Defaults to the number of inputs.
	TopN *int `json:"top_n,omitempty"`
}

// RerankResult The result of a reranking request.
type RerankResult struct {
	// Data The reranked documents.
	Data []RankedDocument `json:"data"`

	// Model The model used to rerank documents.
	Model string `json:"model"`

	// Usage Usage statistics for the model inference.
	Usage struct {
		// RerankUnits The number of rerank units consumed by this operation.
		RerankUnits *int32 `json:"rerank_units,omitempty"`
	} `json:"usage"`
}

// SparseEmbedding A sparse embedding of a single input
type SparseEmbedding struct {
	// SparseIndices The sparse embedding indices.
	SparseIndices []int32 `json:"sparse_indices"`

	// SparseTokens The normalized tokens used to create the sparse embedding.
	SparseTokens *[]string `json:"sparse_tokens,omitempty"`

	// SparseValues The sparse embedding values.
	SparseValues []float64 `json:"sparse_values"`

	// VectorType Indicates whether this is a 'dense' or 'sparse' embedding.
	VectorType VectorType `json:"vector_type"`
}

// VectorType Indicates whether this is a 'dense' or 'sparse' embedding.
type VectorType string

// EmbedJSONRequestBody defines body for Embed for application/json ContentType.
type EmbedJSONRequestBody = EmbedRequest

// RerankJSONRequestBody defines body for Rerank for application/json ContentType.
type RerankJSONRequestBody = RerankRequest

// AsDenseEmbedding returns the union data inside the Embedding as a DenseEmbedding
func (t Embedding) AsDenseEmbedding() (DenseEmbedding, error) {
	var body DenseEmbedding
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDenseEmbedding overwrites any union data inside the Embedding as the provided DenseEmbedding
func (t *Embedding) FromDenseEmbedding(v DenseEmbedding) error {
	v.VectorType = "dense"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeDenseEmbedding performs a merge with any union data inside the Embedding, using the provided DenseEmbedding
func (t *Embedding) MergeDenseEmbedding(v DenseEmbedding) error {
	v.VectorType = "dense"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsSparseEmbedding returns the union data inside the Embedding as a SparseEmbedding
func (t Embedding) AsSparseEmbedding() (SparseEmbedding, error) {
	var body SparseEmbedding
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSparseEmbedding overwrites any union data inside the Embedding as the provided SparseEmbedding
func (t *Embedding) FromSparseEmbedding(v SparseEmbedding) error {
	v.VectorType = "sparse"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSparseEmbedding performs a merge with any union data inside the Embedding, using the provided SparseEmbedding
func (t *Embedding) MergeSparseEmbedding(v SparseEmbedding) error {
	v.VectorType = "sparse"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Embedding) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"vector_type"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t Embedding) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "dense":
		return t.AsDenseEmbedding()
	case "sparse":
		return t.AsSparseEmbedding()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t Embedding) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Embedding) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// EmbedWithBody request with any body
	EmbedWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Embed(ctx context.Context, body EmbedJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RerankWithBody request with any body
	RerankWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Rerank(ctx context.Context, body RerankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) EmbedWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewEmbedRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Embed(ctx context.Context, body EmbedJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewEmbedRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RerankWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRerankRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Rerank(ctx context.Context, body RerankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRerankRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewEmbedRequest calls the generic Embed builder with application/json body
func NewEmbedRequest(server string, body EmbedJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewEmbedRequestWithBody(server, "application/json", bodyReader)
}

// NewEmbedRequestWithBody generates requests for Embed with any type of body
func NewEmbedRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/embed")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRerankRequest calls the generic Rerank builder with application/json body
func NewRerankRequest(server string, body RerankJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRerankRequestWithBody(server, "application/json", bodyReader)
}

// NewRerankRequestWithBody generates requests for Rerank with any type of body
func NewRerankRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/rerank")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// EmbedWithBodyWithResponse request with any body
	EmbedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*EmbedResponse, error)

	EmbedWithResponse(ctx context.Context, body EmbedJSONRequestBody, reqEditors ...RequestEditorFn) (*EmbedResponse, error)

	// RerankWithBodyWithResponse request with any body
	RerankWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RerankResponse, error)

	RerankWithResponse(ctx context.Context, body RerankJSONRequestBody, reqEditors ...RequestEditorFn) (*RerankResponse, error)
}

type EmbedResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *EmbeddingsList
	JSON400      *ErrorResponse
	JSON401      *ErrorResponse
	JSON500      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r EmbedResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r EmbedResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RerankResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RerankResult
	JSON400      *ErrorResponse
	JSON401      *ErrorResponse
	JSON500      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r RerankResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RerankResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// EmbedWithBodyWithResponse request with arbitrary body returning *EmbedResponse
func (c *ClientWithResponses) EmbedWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*EmbedResponse, error) {
	rsp, err := c.EmbedWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseEmbedResponse(rsp)
}

func (c *ClientWithResponses) EmbedWithResponse(ctx context.Context, body EmbedJSONRequestBody, reqEditors ...RequestEditorFn) (*EmbedResponse, error) {
	rsp, err := c.Embed(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseEmbedResponse(rsp)
}

// RerankWithBodyWithResponse request with arbitrary body returning *RerankResponse
func (c *ClientWithResponses) RerankWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RerankResponse, error) {
	rsp, err := c.RerankWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRerankResponse(rsp)
}

func (c *ClientWithResponses) RerankWithResponse(ctx context.Context, body RerankJSONRequestBody, reqEditors ...RequestEditorFn) (*RerankResponse, error) {
	rsp, err := c.Rerank(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRerankResponse(rsp)
}

// ParseEmbedResponse parses an HTTP response from a EmbedWithResponse call
func ParseEmbedResponse(rsp *http.Response) (*EmbedResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &EmbedResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest EmbeddingsList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseRerankResponse parses an HTTP response from a RerankWithResponse call
func ParseRerankResponse(rsp *http.Response) (*RerankResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RerankResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RerankResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
