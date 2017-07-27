// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Network functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Timing information for the request.
type NetworkResourceTiming struct {
	RequestTime       float64 `json:"requestTime"`       // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart        float64 `json:"proxyStart"`        // Started resolving proxy.
	ProxyEnd          float64 `json:"proxyEnd"`          // Finished resolving proxy.
	DnsStart          float64 `json:"dnsStart"`          // Started DNS address resolve.
	DnsEnd            float64 `json:"dnsEnd"`            // Finished DNS address resolve.
	ConnectStart      float64 `json:"connectStart"`      // Started connecting to the remote host.
	ConnectEnd        float64 `json:"connectEnd"`        // Connected to the remote host.
	SslStart          float64 `json:"sslStart"`          // Started SSL handshake.
	SslEnd            float64 `json:"sslEnd"`            // Finished SSL handshake.
	WorkerStart       float64 `json:"workerStart"`       // Started running ServiceWorker.
	WorkerReady       float64 `json:"workerReady"`       // Finished Starting ServiceWorker.
	SendStart         float64 `json:"sendStart"`         // Started sending request.
	SendEnd           float64 `json:"sendEnd"`           // Finished sending request.
	PushStart         float64 `json:"pushStart"`         // Time the server started pushing request.
	PushEnd           float64 `json:"pushEnd"`           // Time the server finished pushing request.
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// HTTP request data.
type NetworkRequest struct {
	Url              string                 `json:"url"`                        // Request URL.
	Method           string                 `json:"method"`                     // HTTP request method.
	Headers          map[string]interface{} `json:"headers"`                    // HTTP request headers.
	PostData         string                 `json:"postData,omitempty"`         // HTTP POST request data.
	MixedContentType string                 `json:"mixedContentType,omitempty"` // The mixed content type of the request. enum values: blockable, optionally-blockable, none
	InitialPriority  string                 `json:"initialPriority"`            // Priority of the resource request at the time request is sent. enum values: VeryLow, Low, Medium, High, VeryHigh
	ReferrerPolicy   string                 `json:"referrerPolicy"`             // The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload    bool                   `json:"isLinkPreload,omitempty"`    // Whether is loaded via link preload.
}

// Details of a signed certificate timestamp (SCT).
type NetworkSignedCertificateTimestamp struct {
	Status             string  `json:"status"`             // Validation status.
	Origin             string  `json:"origin"`             // Origin.
	LogDescription     string  `json:"logDescription"`     // Log name / description.
	LogId              string  `json:"logId"`              // Log ID.
	Timestamp          float64 `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string  `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string  `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string  `json:"signatureData"`      // Signature data.
}

// Security details about a request.
type NetworkSecurityDetails struct {
	Protocol                       string                               `json:"protocol"`                       // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                    string                               `json:"keyExchange"`                    // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup               string                               `json:"keyExchangeGroup,omitempty"`     // (EC)DH group used by the connection, if applicable.
	Cipher                         string                               `json:"cipher"`                         // Cipher name.
	Mac                            string                               `json:"mac,omitempty"`                  // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId                  int                                  `json:"certificateId"`                  // Certificate ID value.
	SubjectName                    string                               `json:"subjectName"`                    // Certificate subject name.
	SanList                        []string                             `json:"sanList"`                        // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                         string                               `json:"issuer"`                         // Name of the issuing CA.
	ValidFrom                      float64                              `json:"validFrom"`                      // Certificate valid from date.
	ValidTo                        float64                              `json:"validTo"`                        // Certificate valid to (expiration) date
	SignedCertificateTimestampList []*NetworkSignedCertificateTimestamp `json:"signedCertificateTimestampList"` // List of signed certificate timestamps (SCTs).
}

// HTTP response data.
type NetworkResponse struct {
	Url                string                  `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             float64                 `json:"status"`                       // HTTP response status code.
	StatusText         string                  `json:"statusText"`                   // HTTP response status text.
	Headers            map[string]interface{}  `json:"headers"`                      // HTTP response headers.
	HeadersText        string                  `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string                  `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     map[string]interface{}  `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText string                  `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool                    `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionId       float64                 `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    string                  `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         int                     `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      bool                    `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  bool                    `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64                 `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *NetworkResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           string                  `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      string                  `json:"securityState"`                // Security state of the request resource. enum values: unknown, neutral, insecure, warning, secure, info
	SecurityDetails    *NetworkSecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocket request data.
type NetworkWebSocketRequest struct {
	Headers map[string]interface{} `json:"headers"` // HTTP request headers.
}

// WebSocket response data.
type NetworkWebSocketResponse struct {
	Status             float64                `json:"status"`                       // HTTP response status code.
	StatusText         string                 `json:"statusText"`                   // HTTP response status text.
	Headers            map[string]interface{} `json:"headers"`                      // HTTP response headers.
	HeadersText        string                 `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     map[string]interface{} `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText string                 `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocket frame data.
type NetworkWebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// Information about the cached resource.
type NetworkCachedResource struct {
	Url      string           `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     string           `json:"type"`               // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
	Response *NetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize float64          `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type NetworkInitiator struct {
	Type       string             `json:"type"`                 // Type of this initiator.
	Stack      *RuntimeStackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	Url        string             `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module).
	LineNumber float64            `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}

// Cookie object
type NetworkCookie struct {
	Name     string  `json:"name"`               // Cookie name.
	Value    string  `json:"value"`              // Cookie value.
	Domain   string  `json:"domain"`             // Cookie domain.
	Path     string  `json:"path"`               // Cookie path.
	Expires  float64 `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int     `json:"size"`               // Cookie size.
	HttpOnly bool    `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool    `json:"secure"`             // True if cookie is secure.
	Session  bool    `json:"session"`            // True in case of session cookie.
	SameSite string  `json:"sameSite,omitempty"` // Cookie SameSite type. enum values: Strict, Lax
}

// Authorization challenge for HTTP status code 401 or 407.
type NetworkAuthChallenge struct {
	Source string `json:"source,omitempty"` // Source of the authentication challenge.
	Origin string `json:"origin"`           // Origin of the challenger.
	Scheme string `json:"scheme"`           // The authentication scheme used, such as basic or digest
	Realm  string `json:"realm"`            // The realm of the challenge. May be empty.
}

// Response to an AuthChallenge.
type NetworkAuthChallengeResponse struct {
	Response string `json:"response"`           // The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// Fired when resource loading priority is changed
type NetworkResourceChangedPriorityEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId   string  `json:"requestId"`   // Request identifier.
		NewPriority string  `json:"newPriority"` // New priority enum values: VeryLow, Low, Medium, High, VeryHigh
		Timestamp   float64 `json:"timestamp"`   // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when page is about to send HTTP request.
type NetworkRequestWillBeSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId        string            `json:"requestId"`                  // Request identifier.
		LoaderId         string            `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched form worker.
		DocumentURL      string            `json:"documentURL"`                // URL of the document this request is loaded for.
		Request          *NetworkRequest   `json:"request"`                    // Request data.
		Timestamp        float64           `json:"timestamp"`                  // Timestamp.
		WallTime         float64           `json:"wallTime"`                   // Timestamp.
		Initiator        *NetworkInitiator `json:"initiator"`                  // Request initiator.
		RedirectResponse *NetworkResponse  `json:"redirectResponse,omitempty"` // Redirect response data.
		Type             string            `json:"type,omitempty"`             // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		FrameId          string            `json:"frameId,omitempty"`          // Frame identifier.
	} `json:"Params,omitempty"`
}

// Fired if request ended up loading from cache.
type NetworkRequestServedFromCacheEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string `json:"requestId"` // Request identifier.
	} `json:"Params,omitempty"`
}

// Fired when HTTP response is available.
type NetworkResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string           `json:"requestId"`         // Request identifier.
		LoaderId  string           `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched form worker.
		Timestamp float64          `json:"timestamp"`         // Timestamp.
		Type      string           `json:"type"`              // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		Response  *NetworkResponse `json:"response"`          // Response data.
		FrameId   string           `json:"frameId,omitempty"` // Frame identifier.
	} `json:"Params,omitempty"`
}

// Fired when data chunk was received over the network.
type NetworkDataReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId         string  `json:"requestId"`         // Request identifier.
		Timestamp         float64 `json:"timestamp"`         // Timestamp.
		DataLength        int     `json:"dataLength"`        // Data chunk length.
		EncodedDataLength int     `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
	} `json:"Params,omitempty"`
}

// Fired when HTTP request has finished loading.
type NetworkLoadingFinishedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId         string  `json:"requestId"`         // Request identifier.
		Timestamp         float64 `json:"timestamp"`         // Timestamp.
		EncodedDataLength float64 `json:"encodedDataLength"` // Total number of bytes received for this request.
	} `json:"Params,omitempty"`
}

// Fired when HTTP request has failed to load.
type NetworkLoadingFailedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId     string  `json:"requestId"`               // Request identifier.
		Timestamp     float64 `json:"timestamp"`               // Timestamp.
		Type          string  `json:"type"`                    // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		ErrorText     string  `json:"errorText"`               // User friendly error message.
		Canceled      bool    `json:"canceled,omitempty"`      // True if loading was canceled.
		BlockedReason string  `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any. enum values: csp, mixed-content, origin, inspector, subresource-filter, other
	} `json:"Params,omitempty"`
}

// Fired when WebSocket is about to initiate handshake.
type NetworkWebSocketWillSendHandshakeRequestEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                   `json:"requestId"` // Request identifier.
		Timestamp float64                  `json:"timestamp"` // Timestamp.
		WallTime  float64                  `json:"wallTime"`  // UTC Timestamp.
		Request   *NetworkWebSocketRequest `json:"request"`   // WebSocket request data.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket handshake response becomes available.
type NetworkWebSocketHandshakeResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                    `json:"requestId"` // Request identifier.
		Timestamp float64                   `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketResponse `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired upon WebSocket creation.
type NetworkWebSocketCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string            `json:"requestId"`           // Request identifier.
		Url       string            `json:"url"`                 // WebSocket request URL.
		Initiator *NetworkInitiator `json:"initiator,omitempty"` // Request initiator.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket is closed.
type NetworkWebSocketClosedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string  `json:"requestId"` // Request identifier.
		Timestamp float64 `json:"timestamp"` // Timestamp.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame is received.
type NetworkWebSocketFrameReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame error occurs.
type NetworkWebSocketFrameErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId    string  `json:"requestId"`    // Request identifier.
		Timestamp    float64 `json:"timestamp"`    // Timestamp.
		ErrorMessage string  `json:"errorMessage"` // WebSocket frame error message.
	} `json:"Params,omitempty"`
}

// Fired when WebSocket frame is sent.
type NetworkWebSocketFrameSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
	} `json:"Params,omitempty"`
}

// Fired when EventSource message is received.
type NetworkEventSourceMessageReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string  `json:"requestId"` // Request identifier.
		Timestamp float64 `json:"timestamp"` // Timestamp.
		EventName string  `json:"eventName"` // Message type.
		EventId   string  `json:"eventId"`   // Message identifier.
		Data      string  `json:"data"`      // Message content.
	} `json:"Params,omitempty"`
}

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or mocked.
type NetworkRequestInterceptedEvent struct {
	Method string `json:"method"`
	Params struct {
		InterceptionId     string                 `json:"interceptionId"`               // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
		Request            *NetworkRequest        `json:"request"`                      //
		ResourceType       string                 `json:"resourceType"`                 // How the requested resource will be used. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, Other
		RedirectHeaders    map[string]interface{} `json:"redirectHeaders,omitempty"`    // HTTP response headers, only sent if a redirect was intercepted.
		RedirectStatusCode int                    `json:"redirectStatusCode,omitempty"` // HTTP response code, only sent if a redirect was intercepted.
		RedirectUrl        string                 `json:"redirectUrl,omitempty"`        // Redirect location, only sent if a redirect was intercepted.
		AuthChallenge      *NetworkAuthChallenge  `json:"authChallenge,omitempty"`      // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	} `json:"Params,omitempty"`
}

type Network struct {
	target gcdmessage.ChromeTargeter
}

func NewNetwork(target gcdmessage.ChromeTargeter) *Network {
	c := &Network{target: target}
	return c
}

type NetworkEnableParams struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
}

// EnableWithParams - Enables network tracking, network events will now be delivered to the client.
func (c *Network) EnableWithParams(v *NetworkEnableParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enable", Params: v})
}

// Enable - Enables network tracking, network events will now be delivered to the client.
// maxTotalBufferSize - Buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxResourceBufferSize - Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
func (c *Network) Enable(maxTotalBufferSize int, maxResourceBufferSize int) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEnableParams
	v.MaxTotalBufferSize = maxTotalBufferSize
	v.MaxResourceBufferSize = maxResourceBufferSize
	return c.EnableWithParams(&v)
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *Network) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.disable"})
}

type NetworkSetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
}

// SetUserAgentOverrideWithParams - Allows overriding user agent with the given string.
func (c *Network) SetUserAgentOverrideWithParams(v *NetworkSetUserAgentOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setUserAgentOverride", Params: v})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
func (c *Network) SetUserAgentOverride(userAgent string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetUserAgentOverrideParams
	v.UserAgent = userAgent
	return c.SetUserAgentOverrideWithParams(&v)
}

type NetworkSetExtraHTTPHeadersParams struct {
	// Map with extra HTTP headers.
	Headers map[string]interface{} `json:"headers"`
}

// SetExtraHTTPHeadersWithParams - Specifies whether to always send extra HTTP headers with the requests from this page.
func (c *Network) SetExtraHTTPHeadersWithParams(v *NetworkSetExtraHTTPHeadersParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setExtraHTTPHeaders", Params: v})
}

// SetExtraHTTPHeaders - Specifies whether to always send extra HTTP headers with the requests from this page.
// headers - Map with extra HTTP headers.
func (c *Network) SetExtraHTTPHeaders(headers map[string]interface{}) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetExtraHTTPHeadersParams
	v.Headers = headers
	return c.SetExtraHTTPHeadersWithParams(&v)
}

type NetworkGetResponseBodyParams struct {
	// Identifier of the network request to get content for.
	RequestId string `json:"requestId"`
}

// GetResponseBodyWithParams - Returns content served for the given request.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyWithParams(v *NetworkGetResponseBodyParams) (string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBody", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Body          string
			Base64Encoded bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.Body, chromeData.Result.Base64Encoded, nil
}

// GetResponseBody - Returns content served for the given request.
// requestId - Identifier of the network request to get content for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBody(requestId string) (string, bool, error) {
	var v NetworkGetResponseBodyParams
	v.RequestId = requestId
	return c.GetResponseBodyWithParams(&v)
}

type NetworkSetBlockedURLsParams struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls []string `json:"urls"`
}

// SetBlockedURLsWithParams - Blocks URLs from loading.
func (c *Network) SetBlockedURLsWithParams(v *NetworkSetBlockedURLsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBlockedURLs", Params: v})
}

// SetBlockedURLs - Blocks URLs from loading.
// urls - URL patterns to block. Wildcards ('*') are allowed.
func (c *Network) SetBlockedURLs(urls []string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetBlockedURLsParams
	v.Urls = urls
	return c.SetBlockedURLsWithParams(&v)
}

type NetworkReplayXHRParams struct {
	// Identifier of XHR to replay.
	RequestId string `json:"requestId"`
}

// ReplayXHRWithParams - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
func (c *Network) ReplayXHRWithParams(v *NetworkReplayXHRParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.replayXHR", Params: v})
}

// ReplayXHR - This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
// requestId - Identifier of XHR to replay.
func (c *Network) ReplayXHR(requestId string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkReplayXHRParams
	v.RequestId = requestId
	return c.ReplayXHRWithParams(&v)
}

// CanClearBrowserCache - Tells whether clearing browser cache is supported.
// Returns -  result - True if browser cache can be cleared.
func (c *Network) CanClearBrowserCache() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCache"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cache.
func (c *Network) ClearBrowserCache() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCache"})
}

// CanClearBrowserCookies - Tells whether clearing browser cookies is supported.
// Returns -  result - True if browser cookies can be cleared.
func (c *Network) CanClearBrowserCookies() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canClearBrowserCookies"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

// Clears browser cookies.
func (c *Network) ClearBrowserCookies() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCookies"})
}

type NetworkGetCookiesParams struct {
	// The list of URLs for which applicable cookies will be fetched
	Urls []string `json:"urls,omitempty"`
}

// GetCookiesWithParams - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookiesWithParams(v *NetworkGetCookiesParams) ([]*NetworkCookie, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCookies", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Cookies, nil
}

// GetCookies - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// urls - The list of URLs for which applicable cookies will be fetched
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookies(urls []string) ([]*NetworkCookie, error) {
	var v NetworkGetCookiesParams
	v.Urls = urls
	return c.GetCookiesWithParams(&v)
}

// GetAllCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetAllCookies() ([]*NetworkCookie, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getAllCookies"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Cookies []*NetworkCookie
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Cookies, nil
}

type NetworkDeleteCookieParams struct {
	// Name of the cookie to remove.
	CookieName string `json:"cookieName"`
	// URL to match cooke domain and path.
	Url string `json:"url"`
}

// DeleteCookieWithParams - Deletes browser cookie with given name, domain and path.
func (c *Network) DeleteCookieWithParams(v *NetworkDeleteCookieParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.deleteCookie", Params: v})
}

// DeleteCookie - Deletes browser cookie with given name, domain and path.
// cookieName - Name of the cookie to remove.
// url - URL to match cooke domain and path.
func (c *Network) DeleteCookie(cookieName string, url string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkDeleteCookieParams
	v.CookieName = cookieName
	v.Url = url
	return c.DeleteCookieWithParams(&v)
}

type NetworkSetCookieParams struct {
	// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Url string `json:"url"`
	// The name of the cookie.
	Name string `json:"name"`
	// The value of the cookie.
	Value string `json:"value"`
	// If omitted, the cookie becomes a host-only cookie.
	Domain string `json:"domain,omitempty"`
	// Defaults to the path portion of the url parameter.
	Path string `json:"path,omitempty"`
	// Defaults ot false.
	Secure bool `json:"secure,omitempty"`
	// Defaults to false.
	HttpOnly bool `json:"httpOnly,omitempty"`
	// Defaults to browser default behavior. enum values: Strict, Lax
	SameSite string `json:"sameSite,omitempty"`
	// If omitted, the cookie becomes a session cookie.
	ExpirationDate float64 `json:"expirationDate,omitempty"`
}

// SetCookieWithParams - Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
// Returns -  success - True if successfully set cookie.
func (c *Network) SetCookieWithParams(v *NetworkSetCookieParams) (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCookie", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// SetCookie - Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
// url - The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
// name - The name of the cookie.
// value - The value of the cookie.
// domain - If omitted, the cookie becomes a host-only cookie.
// path - Defaults to the path portion of the url parameter.
// secure - Defaults ot false.
// httpOnly - Defaults to false.
// sameSite - Defaults to browser default behavior. enum values: Strict, Lax
// expirationDate - If omitted, the cookie becomes a session cookie.
// Returns -  success - True if successfully set cookie.
func (c *Network) SetCookie(url string, name string, value string, domain string, path string, secure bool, httpOnly bool, sameSite string, expirationDate float64) (bool, error) {
	var v NetworkSetCookieParams
	v.Url = url
	v.Name = name
	v.Value = value
	v.Domain = domain
	v.Path = path
	v.Secure = secure
	v.HttpOnly = httpOnly
	v.SameSite = sameSite
	v.ExpirationDate = expirationDate
	return c.SetCookieWithParams(&v)
}

// CanEmulateNetworkConditions - Tells whether emulation of network conditions is supported.
// Returns -  result - True if emulation of network conditions is supported.
func (c *Network) CanEmulateNetworkConditions() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.canEmulateNetworkConditions"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Result, nil
}

type NetworkEmulateNetworkConditionsParams struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`
	// Additional latency (ms).
	Latency float64 `json:"latency"`
	// Maximal aggregated download throughput.
	DownloadThroughput float64 `json:"downloadThroughput"`
	// Maximal aggregated upload throughput.
	UploadThroughput float64 `json:"uploadThroughput"`
	// Connection type if known. enum values: none, cellular2g, cellular3g, cellular4g, bluetooth, ethernet, wifi, wimax, other
	ConnectionType string `json:"connectionType,omitempty"`
}

// EmulateNetworkConditionsWithParams - Activates emulation of network conditions.
func (c *Network) EmulateNetworkConditionsWithParams(v *NetworkEmulateNetworkConditionsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.emulateNetworkConditions", Params: v})
}

// EmulateNetworkConditions - Activates emulation of network conditions.
// offline - True to emulate internet disconnection.
// latency - Additional latency (ms).
// downloadThroughput - Maximal aggregated download throughput.
// uploadThroughput - Maximal aggregated upload throughput.
// connectionType - Connection type if known. enum values: none, cellular2g, cellular3g, cellular4g, bluetooth, ethernet, wifi, wimax, other
func (c *Network) EmulateNetworkConditions(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64, connectionType string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEmulateNetworkConditionsParams
	v.Offline = offline
	v.Latency = latency
	v.DownloadThroughput = downloadThroughput
	v.UploadThroughput = uploadThroughput
	v.ConnectionType = connectionType
	return c.EmulateNetworkConditionsWithParams(&v)
}

type NetworkSetCacheDisabledParams struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// SetCacheDisabledWithParams - Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
func (c *Network) SetCacheDisabledWithParams(v *NetworkSetCacheDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCacheDisabled", Params: v})
}

// SetCacheDisabled - Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *Network) SetCacheDisabled(cacheDisabled bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetCacheDisabledParams
	v.CacheDisabled = cacheDisabled
	return c.SetCacheDisabledWithParams(&v)
}

type NetworkSetBypassServiceWorkerParams struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

// SetBypassServiceWorkerWithParams - Toggles ignoring of service worker for each request.
func (c *Network) SetBypassServiceWorkerWithParams(v *NetworkSetBypassServiceWorkerParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setBypassServiceWorker", Params: v})
}

// SetBypassServiceWorker - Toggles ignoring of service worker for each request.
// bypass - Bypass service worker and load from network.
func (c *Network) SetBypassServiceWorker(bypass bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetBypassServiceWorkerParams
	v.Bypass = bypass
	return c.SetBypassServiceWorkerWithParams(&v)
}

type NetworkSetDataSizeLimitsForTestParams struct {
	// Maximum total buffer size.
	MaxTotalSize int `json:"maxTotalSize"`
	// Maximum per-resource size.
	MaxResourceSize int `json:"maxResourceSize"`
}

// SetDataSizeLimitsForTestWithParams - For testing.
func (c *Network) SetDataSizeLimitsForTestWithParams(v *NetworkSetDataSizeLimitsForTestParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setDataSizeLimitsForTest", Params: v})
}

// SetDataSizeLimitsForTest - For testing.
// maxTotalSize - Maximum total buffer size.
// maxResourceSize - Maximum per-resource size.
func (c *Network) SetDataSizeLimitsForTest(maxTotalSize int, maxResourceSize int) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetDataSizeLimitsForTestParams
	v.MaxTotalSize = maxTotalSize
	v.MaxResourceSize = maxResourceSize
	return c.SetDataSizeLimitsForTestWithParams(&v)
}

type NetworkGetCertificateParams struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

// GetCertificateWithParams - Returns the DER-encoded certificate.
// Returns -  tableNames -
func (c *Network) GetCertificateWithParams(v *NetworkGetCertificateParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getCertificate", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TableNames []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.TableNames, nil
}

// GetCertificate - Returns the DER-encoded certificate.
// origin - Origin to get certificate for.
// Returns -  tableNames -
func (c *Network) GetCertificate(origin string) ([]string, error) {
	var v NetworkGetCertificateParams
	v.Origin = origin
	return c.GetCertificateWithParams(&v)
}

type NetworkSetRequestInterceptionEnabledParams struct {
	// Whether or not HTTP requests should be intercepted and Network.requestIntercepted events sent.
	Enabled bool `json:"enabled"`
}

// SetRequestInterceptionEnabledWithParams -
func (c *Network) SetRequestInterceptionEnabledWithParams(v *NetworkSetRequestInterceptionEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setRequestInterceptionEnabled", Params: v})
}

// SetRequestInterceptionEnabled -
// enabled - Whether or not HTTP requests should be intercepted and Network.requestIntercepted events sent.
func (c *Network) SetRequestInterceptionEnabled(enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetRequestInterceptionEnabledParams
	v.Enabled = enabled
	return c.SetRequestInterceptionEnabledWithParams(&v)
}

type NetworkContinueInterceptedRequestParams struct {
	//
	InterceptionId string `json:"interceptionId"`
	// If set this causes the request to fail with the given reason. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable
	ErrorReason string `json:"errorReason,omitempty"`
	// If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
	RawResponse string `json:"rawResponse,omitempty"`
	// If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
	Url string `json:"url,omitempty"`
	// If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
	Method string `json:"method,omitempty"`
	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData string `json:"postData,omitempty"`
	// If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse *NetworkAuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

// ContinueInterceptedRequestWithParams - Response to Network.requestIntercepted which either modifies the request to continue with any modifications, or blocks it, or completes it with the provided response bytes. If a network fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted event will be sent with the same InterceptionId.
func (c *Network) ContinueInterceptedRequestWithParams(v *NetworkContinueInterceptedRequestParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.continueInterceptedRequest", Params: v})
}

// ContinueInterceptedRequest - Response to Network.requestIntercepted which either modifies the request to continue with any modifications, or blocks it, or completes it with the provided response bytes. If a network fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted event will be sent with the same InterceptionId.
// interceptionId -
// errorReason - If set this causes the request to fail with the given reason. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable
// rawResponse - If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
// url - If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
// method - If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
// postData - If set this allows postData to be set. Must not be set in response to an authChallenge.
// headers - If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
// authChallengeResponse - Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
func (c *Network) ContinueInterceptedRequest(interceptionId string, errorReason string, rawResponse string, url string, method string, postData string, headers map[string]interface{}, authChallengeResponse *NetworkAuthChallengeResponse) (*gcdmessage.ChromeResponse, error) {
	var v NetworkContinueInterceptedRequestParams
	v.InterceptionId = interceptionId
	v.ErrorReason = errorReason
	v.RawResponse = rawResponse
	v.Url = url
	v.Method = method
	v.PostData = postData
	v.Headers = headers
	v.AuthChallengeResponse = authChallengeResponse
	return c.ContinueInterceptedRequestWithParams(&v)
}
