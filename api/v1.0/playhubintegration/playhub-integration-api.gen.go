// Package playhubintegrationapiv1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package playhubintegrationapiv1

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
	externalRef0 "github.com/playhubstudio/public-api/api/v1.0"
)

// Defines values for BalanceRequestAction.
const (
	Balance BalanceRequestAction = "balance"
)

// Defines values for BetRequestAction.
const (
	BetRequestActionBet BetRequestAction = "bet"
)

// Defines values for BetRequestType.
const (
	BetRequestTypeBet  BetRequestType = "bet"
	BetRequestTypeFree BetRequestType = "free"
)

// Defines values for RefundRequestAction.
const (
	Refund RefundRequestAction = "refund"
)

// Defines values for WinRequestAction.
const (
	WinRequestActionWin WinRequestAction = "win"
)

// Defines values for WinRequestType.
const (
	WinRequestTypeFree WinRequestType = "free"
	WinRequestTypeWin  WinRequestType = "win"
)

// BalanceRequest defines model for BalanceRequest.
type BalanceRequest struct {
	// Action The action name. Should be 'balance'.
	Action BalanceRequestAction `json:"action"`

	// SessionId The Game Session's ID (external). Provided by client via Create New Game request.
	SessionId string `json:"session_id"`
}

// BalanceRequestAction The action name. Should be 'balance'.
type BalanceRequestAction string

// BalanceResponse defines model for BalanceResponse.
type BalanceResponse struct {
	// Balance The user's balance in selected currency units. Currency selected by the client during the
	// Create New Game call.
	Balance string `json:"balance"`
}

// BetRequest defines model for BetRequest.
type BetRequest struct {
	// Action The action name. Should be 'bet'.
	Action BetRequestAction `json:"action"`

	// Amount The amount of the bet in currency units.
	Amount string `json:"amount"`

	// SessionId The Game Session's ID (external). Provided by client via Create New Game request.
	SessionId string `json:"session_id"`

	// TxId Unique ID for the bet (internal) on Playhub side.
	// Could be used as idempotency key (within a single round).
	TxId openapi_types.UUID `json:"tx_id"`

	// Type The type of the bet. 'Bet' type means regular bet, 'Free' type means free bet (see Free Rounds section).
	Type BetRequestType `json:"type"`
}

// BetRequestAction The action name. Should be 'bet'.
type BetRequestAction string

// BetRequestType The type of the bet. 'Bet' type means regular bet, 'Free' type means free bet (see Free Rounds section).
type BetRequestType string

// BetResponse defines model for BetResponse.
type BetResponse struct {
	// Balance The user's balance in currency units after bet is applied.
	Balance string `json:"balance"`

	// TxId Unique ID for the bet on your side. This is required for further tracking/debugging pusposes.
	TxId string `json:"tx_id"`
}

// RefundRequest defines model for RefundRequest.
type RefundRequest struct {
	// Action The action name. Should be 'refund'.
	Action RefundRequestAction `json:"action"`

	// SessionId The Game Session's ID (external). Provided by client via Create New Game request.
	SessionId string `json:"session_id"`

	// TxId Unique ID for the bet (internal) on Playhub side for tracking purposes.
	// Could be used as idempotency key (within a single transaction).
	TxId openapi_types.UUID `json:"tx_id"`
}

// RefundRequestAction The action name. Should be 'refund'.
type RefundRequestAction string

// RefundResponse defines model for RefundResponse.
type RefundResponse struct {
	// Balance The user's balance in currency units after win is applied.
	Balance string `json:"balance"`

	// TxId Unique ID for the refund  transaction on your side. This is required for further tracking/debugging pusposes.
	TxId string `json:"tx_id"`
}

// WinRequest defines model for WinRequest.
type WinRequest struct {
	// Action The action name. Should be 'win'.
	Action WinRequestAction `json:"action"`

	// Amount The amount of the bet in currency units.
	Amount string `json:"amount"`

	// SessionId The Game Session's ID (external). Provided by client via Create New Game request.
	SessionId string `json:"session_id"`

	// TxId Unique ID for the win (internal) on Playhub side.
	// Could be used as idempotency key (within a single round).
	TxId openapi_types.UUID `json:"tx_id"`

	// Type The type of the win. 'Win' type means regular bet, 'Free' type means free bet (see Free Rounds section).
	Type WinRequestType `json:"type"`
}

// WinRequestAction The action name. Should be 'win'.
type WinRequestAction string

// WinRequestType The type of the win. 'Win' type means regular bet, 'Free' type means free bet (see Free Rounds section).
type WinRequestType string

// WinResponse defines model for WinResponse.
type WinResponse struct {
	// Balance The user's balance in currency units after win is applied.
	Balance string `json:"balance"`

	// TxId Unique ID for the win on your side. This is required for further tracking/debugging pusposes.
	TxId string `json:"tx_id"`
}

// PostBalanceParams defines parameters for PostBalance.
type PostBalanceParams struct {
	// XREQUESTSIGN Request signature (read more in the Authentication section)
	XREQUESTSIGN string `json:"X-REQUEST-SIGN"`
}

// PostBetParams defines parameters for PostBet.
type PostBetParams struct {
	// XREQUESTSIGN Request signature (read more in the Authentication section)
	XREQUESTSIGN string `json:"X-REQUEST-SIGN"`
}

// PostRefundParams defines parameters for PostRefund.
type PostRefundParams struct {
	// XREQUESTSIGN Request signature (read more in the Authentication section)
	XREQUESTSIGN string `json:"X-REQUEST-SIGN"`
}

// PostWinParams defines parameters for PostWin.
type PostWinParams struct {
	// XREQUESTSIGN Request signature (read more in the Authentication section)
	XREQUESTSIGN string `json:"X-REQUEST-SIGN"`
}

// PostBalanceJSONRequestBody defines body for PostBalance for application/json ContentType.
type PostBalanceJSONRequestBody = BalanceRequest

// PostBetJSONRequestBody defines body for PostBet for application/json ContentType.
type PostBetJSONRequestBody = BetRequest

// PostRefundJSONRequestBody defines body for PostRefund for application/json ContentType.
type PostRefundJSONRequestBody = RefundRequest

// PostWinJSONRequestBody defines body for PostWin for application/json ContentType.
type PostWinJSONRequestBody = WinRequest

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
	// PostBalanceWithBody request with any body
	PostBalanceWithBody(ctx context.Context, params *PostBalanceParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostBalance(ctx context.Context, params *PostBalanceParams, body PostBalanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostBetWithBody request with any body
	PostBetWithBody(ctx context.Context, params *PostBetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostBet(ctx context.Context, params *PostBetParams, body PostBetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostRefundWithBody request with any body
	PostRefundWithBody(ctx context.Context, params *PostRefundParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostRefund(ctx context.Context, params *PostRefundParams, body PostRefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostWinWithBody request with any body
	PostWinWithBody(ctx context.Context, params *PostWinParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostWin(ctx context.Context, params *PostWinParams, body PostWinJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostBalanceWithBody(ctx context.Context, params *PostBalanceParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBalanceRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostBalance(ctx context.Context, params *PostBalanceParams, body PostBalanceJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBalanceRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostBetWithBody(ctx context.Context, params *PostBetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBetRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostBet(ctx context.Context, params *PostBetParams, body PostBetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostBetRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostRefundWithBody(ctx context.Context, params *PostRefundParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostRefundRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostRefund(ctx context.Context, params *PostRefundParams, body PostRefundJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostRefundRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostWinWithBody(ctx context.Context, params *PostWinParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostWinRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostWin(ctx context.Context, params *PostWinParams, body PostWinJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostWinRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostBalanceRequest calls the generic PostBalance builder with application/json body
func NewPostBalanceRequest(server string, params *PostBalanceParams, body PostBalanceJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostBalanceRequestWithBody(server, params, "application/json", bodyReader)
}

// NewPostBalanceRequestWithBody generates requests for PostBalance with any type of body
func NewPostBalanceRequestWithBody(server string, params *PostBalanceParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/balance")
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

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "X-REQUEST-SIGN", runtime.ParamLocationHeader, params.XREQUESTSIGN)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-REQUEST-SIGN", headerParam0)

	}

	return req, nil
}

// NewPostBetRequest calls the generic PostBet builder with application/json body
func NewPostBetRequest(server string, params *PostBetParams, body PostBetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostBetRequestWithBody(server, params, "application/json", bodyReader)
}

// NewPostBetRequestWithBody generates requests for PostBet with any type of body
func NewPostBetRequestWithBody(server string, params *PostBetParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/bet")
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

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "X-REQUEST-SIGN", runtime.ParamLocationHeader, params.XREQUESTSIGN)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-REQUEST-SIGN", headerParam0)

	}

	return req, nil
}

// NewPostRefundRequest calls the generic PostRefund builder with application/json body
func NewPostRefundRequest(server string, params *PostRefundParams, body PostRefundJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostRefundRequestWithBody(server, params, "application/json", bodyReader)
}

// NewPostRefundRequestWithBody generates requests for PostRefund with any type of body
func NewPostRefundRequestWithBody(server string, params *PostRefundParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/refund")
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

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "X-REQUEST-SIGN", runtime.ParamLocationHeader, params.XREQUESTSIGN)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-REQUEST-SIGN", headerParam0)

	}

	return req, nil
}

// NewPostWinRequest calls the generic PostWin builder with application/json body
func NewPostWinRequest(server string, params *PostWinParams, body PostWinJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostWinRequestWithBody(server, params, "application/json", bodyReader)
}

// NewPostWinRequestWithBody generates requests for PostWin with any type of body
func NewPostWinRequestWithBody(server string, params *PostWinParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/win")
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

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "X-REQUEST-SIGN", runtime.ParamLocationHeader, params.XREQUESTSIGN)
		if err != nil {
			return nil, err
		}

		req.Header.Set("X-REQUEST-SIGN", headerParam0)

	}

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
	// PostBalanceWithBodyWithResponse request with any body
	PostBalanceWithBodyWithResponse(ctx context.Context, params *PostBalanceParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBalanceResponse, error)

	PostBalanceWithResponse(ctx context.Context, params *PostBalanceParams, body PostBalanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBalanceResponse, error)

	// PostBetWithBodyWithResponse request with any body
	PostBetWithBodyWithResponse(ctx context.Context, params *PostBetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBetResponse, error)

	PostBetWithResponse(ctx context.Context, params *PostBetParams, body PostBetJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBetResponse, error)

	// PostRefundWithBodyWithResponse request with any body
	PostRefundWithBodyWithResponse(ctx context.Context, params *PostRefundParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRefundResponse, error)

	PostRefundWithResponse(ctx context.Context, params *PostRefundParams, body PostRefundJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRefundResponse, error)

	// PostWinWithBodyWithResponse request with any body
	PostWinWithBodyWithResponse(ctx context.Context, params *PostWinParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostWinResponse, error)

	PostWinWithResponse(ctx context.Context, params *PostWinParams, body PostWinJSONRequestBody, reqEditors ...RequestEditorFn) (*PostWinResponse, error)
}

type PostBalanceResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BalanceResponse
	JSON400      *externalRef0.ErrorResponse
	JSON401      *externalRef0.ErrorResponse
	JSON404      *externalRef0.ErrorResponse
	JSON500      *externalRef0.ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostBalanceResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostBalanceResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostBetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BetResponse
	JSON400      *externalRef0.ErrorResponse
	JSON401      *externalRef0.ErrorResponse
	JSON404      *externalRef0.ErrorResponse
	JSON500      *externalRef0.ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostBetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostBetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostRefundResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RefundResponse
	JSON400      *externalRef0.ErrorResponse
	JSON401      *externalRef0.ErrorResponse
	JSON404      *externalRef0.ErrorResponse
	JSON500      *externalRef0.ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostRefundResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostRefundResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostWinResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *WinResponse
	JSON400      *externalRef0.ErrorResponse
	JSON401      *externalRef0.ErrorResponse
	JSON404      *externalRef0.ErrorResponse
	JSON500      *externalRef0.ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostWinResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostWinResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostBalanceWithBodyWithResponse request with arbitrary body returning *PostBalanceResponse
func (c *ClientWithResponses) PostBalanceWithBodyWithResponse(ctx context.Context, params *PostBalanceParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBalanceResponse, error) {
	rsp, err := c.PostBalanceWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBalanceResponse(rsp)
}

func (c *ClientWithResponses) PostBalanceWithResponse(ctx context.Context, params *PostBalanceParams, body PostBalanceJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBalanceResponse, error) {
	rsp, err := c.PostBalance(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBalanceResponse(rsp)
}

// PostBetWithBodyWithResponse request with arbitrary body returning *PostBetResponse
func (c *ClientWithResponses) PostBetWithBodyWithResponse(ctx context.Context, params *PostBetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostBetResponse, error) {
	rsp, err := c.PostBetWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBetResponse(rsp)
}

func (c *ClientWithResponses) PostBetWithResponse(ctx context.Context, params *PostBetParams, body PostBetJSONRequestBody, reqEditors ...RequestEditorFn) (*PostBetResponse, error) {
	rsp, err := c.PostBet(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostBetResponse(rsp)
}

// PostRefundWithBodyWithResponse request with arbitrary body returning *PostRefundResponse
func (c *ClientWithResponses) PostRefundWithBodyWithResponse(ctx context.Context, params *PostRefundParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRefundResponse, error) {
	rsp, err := c.PostRefundWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostRefundResponse(rsp)
}

func (c *ClientWithResponses) PostRefundWithResponse(ctx context.Context, params *PostRefundParams, body PostRefundJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRefundResponse, error) {
	rsp, err := c.PostRefund(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostRefundResponse(rsp)
}

// PostWinWithBodyWithResponse request with arbitrary body returning *PostWinResponse
func (c *ClientWithResponses) PostWinWithBodyWithResponse(ctx context.Context, params *PostWinParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostWinResponse, error) {
	rsp, err := c.PostWinWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostWinResponse(rsp)
}

func (c *ClientWithResponses) PostWinWithResponse(ctx context.Context, params *PostWinParams, body PostWinJSONRequestBody, reqEditors ...RequestEditorFn) (*PostWinResponse, error) {
	rsp, err := c.PostWin(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostWinResponse(rsp)
}

// ParsePostBalanceResponse parses an HTTP response from a PostBalanceWithResponse call
func ParsePostBalanceResponse(rsp *http.Response) (*PostBalanceResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostBalanceResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BalanceResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostBetResponse parses an HTTP response from a PostBetWithResponse call
func ParsePostBetResponse(rsp *http.Response) (*PostBetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostBetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BetResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostRefundResponse parses an HTTP response from a PostRefundWithResponse call
func ParsePostRefundResponse(rsp *http.Response) (*PostRefundResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostRefundResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RefundResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostWinResponse parses an HTTP response from a PostWinWithResponse call
func ParsePostWinResponse(rsp *http.Response) (*PostWinResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostWinResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest WinResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AgAA//93WnsTiigAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "../common-components.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
