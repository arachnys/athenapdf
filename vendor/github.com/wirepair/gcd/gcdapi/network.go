// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Network functionality.
// API Version: 1.3

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
	HasPostData      bool                   `json:"hasPostData,omitempty"`      // True when the request has POST data. Note that postData might still be omitted when this flag is true when the data is too long.
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
	Protocol                          string                               `json:"protocol"`                          // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                       string                               `json:"keyExchange"`                       // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup                  string                               `json:"keyExchangeGroup,omitempty"`        // (EC)DH group used by the connection, if applicable.
	Cipher                            string                               `json:"cipher"`                            // Cipher name.
	Mac                               string                               `json:"mac,omitempty"`                     // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId                     int                                  `json:"certificateId"`                     // Certificate ID value.
	SubjectName                       string                               `json:"subjectName"`                       // Certificate subject name.
	SanList                           []string                             `json:"sanList"`                           // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                            string                               `json:"issuer"`                            // Name of the issuing CA.
	ValidFrom                         float64                              `json:"validFrom"`                         // Certificate valid from date.
	ValidTo                           float64                              `json:"validTo"`                           // Certificate valid to (expiration) date
	SignedCertificateTimestampList    []*NetworkSignedCertificateTimestamp `json:"signedCertificateTimestampList"`    // List of signed certificate timestamps (SCTs).
	CertificateTransparencyCompliance string                               `json:"certificateTransparencyCompliance"` // Whether the request complied with Certificate Transparency policy enum values: unknown, not-compliant, compliant
}

// HTTP response data.
type NetworkResponse struct {
	Url                string                  `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             int                     `json:"status"`                       // HTTP response status code.
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
	SecurityState      string                  `json:"securityState"`                // Security state of the request resource. enum values: unknown, neutral, insecure, secure, info
	SecurityDetails    *NetworkSecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocket request data.
type NetworkWebSocketRequest struct {
	Headers map[string]interface{} `json:"headers"` // HTTP request headers.
}

// WebSocket response data.
type NetworkWebSocketResponse struct {
	Status             int                    `json:"status"`                       // HTTP response status code.
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
	Type     string           `json:"type"`               // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
	Response *NetworkResponse `json:"response,omitempty"` // Cached response data.
	BodySize float64          `json:"bodySize"`           // Cached response body size.
}

// Information about the request initiator.
type NetworkInitiator struct {
	Type       string             `json:"type"`                 // Type of this initiator.
	Stack      *RuntimeStackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	Url        string             `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module) or for SignedExchange type.
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

// Cookie parameter object
type NetworkCookieParam struct {
	Name     string  `json:"name"`               // Cookie name.
	Value    string  `json:"value"`              // Cookie value.
	Url      string  `json:"url,omitempty"`      // The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain   string  `json:"domain,omitempty"`   // Cookie domain.
	Path     string  `json:"path,omitempty"`     // Cookie path.
	Secure   bool    `json:"secure,omitempty"`   // True if cookie is secure.
	HttpOnly bool    `json:"httpOnly,omitempty"` // True if cookie is http-only.
	SameSite string  `json:"sameSite,omitempty"` // Cookie SameSite type. enum values: Strict, Lax
	Expires  float64 `json:"expires,omitempty"`  // Cookie expiration date, session cookie if not set
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

// Request pattern for interception.
type NetworkRequestPattern struct {
	UrlPattern        string `json:"urlPattern,omitempty"`        // Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType      string `json:"resourceType,omitempty"`      // If set, only requests for matching resource types will be intercepted. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
	InterceptionStage string `json:"interceptionStage,omitempty"` // Stage at wich to begin intercepting requests. Default is Request. enum values: Request, HeadersReceived
}

// Information about a signed exchange signature. https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#rfc.section.3.1
type NetworkSignedExchangeSignature struct {
	Label        string   `json:"label"`                  // Signed exchange signature label.
	Signature    string   `json:"signature"`              // The hex string of signed exchange signature.
	Integrity    string   `json:"integrity"`              // Signed exchange signature integrity.
	CertUrl      string   `json:"certUrl,omitempty"`      // Signed exchange signature cert Url.
	CertSha256   string   `json:"certSha256,omitempty"`   // The hex string of signed exchange signature cert sha256.
	ValidityUrl  string   `json:"validityUrl"`            // Signed exchange signature validity Url.
	Date         int      `json:"date"`                   // Signed exchange signature date.
	Expires      int      `json:"expires"`                // Signed exchange signature expires.
	Certificates []string `json:"certificates,omitempty"` // The encoded certificates.
}

// Information about a signed exchange header. https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cbor-representation
type NetworkSignedExchangeHeader struct {
	RequestUrl      string                            `json:"requestUrl"`      // Signed exchange request URL.
	RequestMethod   string                            `json:"requestMethod"`   // Signed exchange request method.
	ResponseCode    int                               `json:"responseCode"`    // Signed exchange response code.
	ResponseHeaders map[string]interface{}            `json:"responseHeaders"` // Signed exchange response headers.
	Signatures      []*NetworkSignedExchangeSignature `json:"signatures"`      // Signed exchange response signature.
}

// Information about a signed exchange response.
type NetworkSignedExchangeError struct {
	Message        string `json:"message"`                  // Error message.
	SignatureIndex int    `json:"signatureIndex,omitempty"` // The index of the signature which caused the error.
	ErrorField     string `json:"errorField,omitempty"`     // The field which caused the error. enum values: signatureSig, signatureIntegrity, signatureCertUrl, signatureCertSha256, signatureValidityUrl, signatureTimestamps
}

// Information about a signed exchange response.
type NetworkSignedExchangeInfo struct {
	OuterResponse   *NetworkResponse              `json:"outerResponse"`             // The outer response of signed HTTP exchange which was received from network.
	Header          *NetworkSignedExchangeHeader  `json:"header,omitempty"`          // Information about the signed exchange header.
	SecurityDetails *NetworkSecurityDetails       `json:"securityDetails,omitempty"` // Security details for the signed exchange header.
	Errors          []*NetworkSignedExchangeError `json:"errors,omitempty"`          // Errors occurred while handling the signed exchagne.
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

// Fired when HTTP request has failed to load.
type NetworkLoadingFailedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId     string  `json:"requestId"`               // Request identifier.
		Timestamp     float64 `json:"timestamp"`               // Timestamp.
		Type          string  `json:"type"`                    // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
		ErrorText     string  `json:"errorText"`               // User friendly error message.
		Canceled      bool    `json:"canceled,omitempty"`      // True if loading was canceled.
		BlockedReason string  `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any. enum values: other, csp, mixed-content, origin, inspector, subresource-filter, content-type
	} `json:"Params,omitempty"`
}

// Fired when HTTP request has finished loading.
type NetworkLoadingFinishedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId                string  `json:"requestId"`                          // Request identifier.
		Timestamp                float64 `json:"timestamp"`                          // Timestamp.
		EncodedDataLength        float64 `json:"encodedDataLength"`                  // Total number of bytes received for this request.
		BlockedCrossSiteDocument bool    `json:"blockedCrossSiteDocument,omitempty"` // Set when response was blocked due to being cross-site document response.
	} `json:"Params,omitempty"`
}

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or mocked.
type NetworkRequestInterceptedEvent struct {
	Method string `json:"method"`
	Params struct {
		InterceptionId      string                 `json:"interceptionId"`                // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
		Request             *NetworkRequest        `json:"request"`                       //
		FrameId             string                 `json:"frameId"`                       // The id of the frame that initiated the request.
		ResourceType        string                 `json:"resourceType"`                  // How the requested resource will be used. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
		IsNavigationRequest bool                   `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
		IsDownload          bool                   `json:"isDownload,omitempty"`          // Set if the request is a navigation that will result in a download. Only present after response is received from the server (i.e. HeadersReceived stage).
		RedirectUrl         string                 `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
		AuthChallenge       *NetworkAuthChallenge  `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
		ResponseErrorReason string                 `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
		ResponseStatusCode  int                    `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
		ResponseHeaders     map[string]interface{} `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
	} `json:"Params,omitempty"`
}

// Fired if request ended up loading from cache.
type NetworkRequestServedFromCacheEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string `json:"requestId"` // Request identifier.
	} `json:"Params,omitempty"`
}

// Fired when page is about to send HTTP request.
type NetworkRequestWillBeSentEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId        string            `json:"requestId"`                  // Request identifier.
		LoaderId         string            `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
		DocumentURL      string            `json:"documentURL"`                // URL of the document this request is loaded for.
		Request          *NetworkRequest   `json:"request"`                    // Request data.
		Timestamp        float64           `json:"timestamp"`                  // Timestamp.
		WallTime         float64           `json:"wallTime"`                   // Timestamp.
		Initiator        *NetworkInitiator `json:"initiator"`                  // Request initiator.
		RedirectResponse *NetworkResponse  `json:"redirectResponse,omitempty"` // Redirect response data.
		Type             string            `json:"type,omitempty"`             // Type of this resource. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
		FrameId          string            `json:"frameId,omitempty"`          // Frame identifier.
		HasUserGesture   bool              `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
	} `json:"Params,omitempty"`
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

// Fired when a signed exchange was received over the network
type NetworkSignedExchangeReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                     `json:"requestId"` // Request identifier.
		Info      *NetworkSignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
	} `json:"Params,omitempty"`
}

// Fired when HTTP response is available.
type NetworkResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string           `json:"requestId"`         // Request identifier.
		LoaderId  string           `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
		Timestamp float64          `json:"timestamp"`         // Timestamp.
		Type      string           `json:"type"`              // Resource type. enum values: Document, Stylesheet, Image, Media, Font, Script, TextTrack, XHR, Fetch, EventSource, WebSocket, Manifest, SignedExchange, Other
		Response  *NetworkResponse `json:"response"`          // Response data.
		FrameId   string           `json:"frameId,omitempty"` // Frame identifier.
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

// Fired upon WebSocket creation.
type NetworkWebSocketCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string            `json:"requestId"`           // Request identifier.
		Url       string            `json:"url"`                 // WebSocket request URL.
		Initiator *NetworkInitiator `json:"initiator,omitempty"` // Request initiator.
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

// Fired when WebSocket frame is received.
type NetworkWebSocketFrameReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                 `json:"requestId"` // Request identifier.
		Timestamp float64                `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketFrame `json:"response"`  // WebSocket response data.
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

// Fired when WebSocket handshake response becomes available.
type NetworkWebSocketHandshakeResponseReceivedEvent struct {
	Method string `json:"method"`
	Params struct {
		RequestId string                    `json:"requestId"` // Request identifier.
		Timestamp float64                   `json:"timestamp"` // Timestamp.
		Response  *NetworkWebSocketResponse `json:"response"`  // WebSocket response data.
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

type Network struct {
	target gcdmessage.ChromeTargeter
}

func NewNetwork(target gcdmessage.ChromeTargeter) *Network {
	c := &Network{target: target}
	return c
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

// Clears browser cache.
func (c *Network) ClearBrowserCache() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCache"})
}

// Clears browser cookies.
func (c *Network) ClearBrowserCookies() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.clearBrowserCookies"})
}

type NetworkContinueInterceptedRequestParams struct {
	//
	InterceptionId string `json:"interceptionId"`
	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
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
// errorReason - If set this causes the request to fail with the given reason. Passing `Aborted` for requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in response to an authChallenge. enum values: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed, ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed, NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient, BlockedByResponse
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

type NetworkDeleteCookiesParams struct {
	// Name of the cookies to remove.
	Name string `json:"name"`
	// If specified, deletes all the cookies with the given name where domain and path match provided URL.
	Url string `json:"url,omitempty"`
	// If specified, deletes only cookies with the exact domain.
	Domain string `json:"domain,omitempty"`
	// If specified, deletes only cookies with the exact path.
	Path string `json:"path,omitempty"`
}

// DeleteCookiesWithParams - Deletes browser cookies with matching name and url or domain/path pair.
func (c *Network) DeleteCookiesWithParams(v *NetworkDeleteCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.deleteCookies", Params: v})
}

// DeleteCookies - Deletes browser cookies with matching name and url or domain/path pair.
// name - Name of the cookies to remove.
// url - If specified, deletes all the cookies with the given name where domain and path match provided URL.
// domain - If specified, deletes only cookies with the exact domain.
// path - If specified, deletes only cookies with the exact path.
func (c *Network) DeleteCookies(name string, url string, domain string, path string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkDeleteCookiesParams
	v.Name = name
	v.Url = url
	v.Domain = domain
	v.Path = path
	return c.DeleteCookiesWithParams(&v)
}

// Disables network tracking, prevents network events from being sent to the client.
func (c *Network) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.disable"})
}

type NetworkEmulateNetworkConditionsParams struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`
	// Minimum latency from request sent to response headers received (ms).
	Latency float64 `json:"latency"`
	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput float64 `json:"downloadThroughput"`
	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
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
// latency - Minimum latency from request sent to response headers received (ms).
// downloadThroughput - Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
// uploadThroughput - Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
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

type NetworkEnableParams struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize int `json:"maxPostDataSize,omitempty"`
}

// EnableWithParams - Enables network tracking, network events will now be delivered to the client.
func (c *Network) EnableWithParams(v *NetworkEnableParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.enable", Params: v})
}

// Enable - Enables network tracking, network events will now be delivered to the client.
// maxTotalBufferSize - Buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxResourceBufferSize - Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
// maxPostDataSize - Longest post body size (in bytes) that would be included in requestWillBeSent notification
func (c *Network) Enable(maxTotalBufferSize int, maxResourceBufferSize int, maxPostDataSize int) (*gcdmessage.ChromeResponse, error) {
	var v NetworkEnableParams
	v.MaxTotalBufferSize = maxTotalBufferSize
	v.MaxResourceBufferSize = maxResourceBufferSize
	v.MaxPostDataSize = maxPostDataSize
	return c.EnableWithParams(&v)
}

// GetAllCookies - Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the `cookies` field.
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

type NetworkGetCookiesParams struct {
	// The list of URLs for which applicable cookies will be fetched
	Urls []string `json:"urls,omitempty"`
}

// GetCookiesWithParams - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the `cookies` field.
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

// GetCookies - Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the `cookies` field.
// urls - The list of URLs for which applicable cookies will be fetched
// Returns -  cookies - Array of cookie objects.
func (c *Network) GetCookies(urls []string) ([]*NetworkCookie, error) {
	var v NetworkGetCookiesParams
	v.Urls = urls
	return c.GetCookiesWithParams(&v)
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

type NetworkGetRequestPostDataParams struct {
	// Identifier of the network request to get content for.
	RequestId string `json:"requestId"`
}

// GetRequestPostDataWithParams - Returns post data sent with the request. Returns an error when no data was sent with the request.
// Returns -  postData - Base64-encoded request body.
func (c *Network) GetRequestPostDataWithParams(v *NetworkGetRequestPostDataParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getRequestPostData", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			PostData string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.PostData, nil
}

// GetRequestPostData - Returns post data sent with the request. Returns an error when no data was sent with the request.
// requestId - Identifier of the network request to get content for.
// Returns -  postData - Base64-encoded request body.
func (c *Network) GetRequestPostData(requestId string) (string, error) {
	var v NetworkGetRequestPostDataParams
	v.RequestId = requestId
	return c.GetRequestPostDataWithParams(&v)
}

type NetworkGetResponseBodyForInterceptionParams struct {
	// Identifier for the intercepted request to get body for.
	InterceptionId string `json:"interceptionId"`
}

// GetResponseBodyForInterceptionWithParams - Returns content served for the given currently intercepted request.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyForInterceptionWithParams(v *NetworkGetResponseBodyForInterceptionParams) (string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.getResponseBodyForInterception", Params: v})
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

// GetResponseBodyForInterception - Returns content served for the given currently intercepted request.
// interceptionId - Identifier for the intercepted request to get body for.
// Returns -  body - Response body. base64Encoded - True, if content was sent as base64.
func (c *Network) GetResponseBodyForInterception(interceptionId string) (string, bool, error) {
	var v NetworkGetResponseBodyForInterceptionParams
	v.InterceptionId = interceptionId
	return c.GetResponseBodyForInterceptionWithParams(&v)
}

type NetworkTakeResponseBodyForInterceptionAsStreamParams struct {
	//
	InterceptionId string `json:"interceptionId"`
}

// TakeResponseBodyForInterceptionAsStreamWithParams - Returns a handle to the stream representing the response body. Note that after this command, the intercepted request can't be continued as is -- you either need to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified.
// Returns -  stream -
func (c *Network) TakeResponseBodyForInterceptionAsStreamWithParams(v *NetworkTakeResponseBodyForInterceptionAsStreamParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.takeResponseBodyForInterceptionAsStream", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Stream string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Stream, nil
}

// TakeResponseBodyForInterceptionAsStream - Returns a handle to the stream representing the response body. Note that after this command, the intercepted request can't be continued as is -- you either need to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified.
// interceptionId -
// Returns -  stream -
func (c *Network) TakeResponseBodyForInterceptionAsStream(interceptionId string) (string, error) {
	var v NetworkTakeResponseBodyForInterceptionAsStreamParams
	v.InterceptionId = interceptionId
	return c.TakeResponseBodyForInterceptionAsStreamWithParams(&v)
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

type NetworkSearchInResponseBodyParams struct {
	// Identifier of the network response to search.
	RequestId string `json:"requestId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInResponseBodyWithParams - Searches for given string in response content.
// Returns -  result - List of search matches.
func (c *Network) SearchInResponseBodyWithParams(v *NetworkSearchInResponseBodyParams) ([]*DebuggerSearchMatch, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.searchInResponseBody", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
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

	return chromeData.Result.Result, nil
}

// SearchInResponseBody - Searches for given string in response content.
// requestId - Identifier of the network response to search.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Network) SearchInResponseBody(requestId string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v NetworkSearchInResponseBodyParams
	v.RequestId = requestId
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInResponseBodyWithParams(&v)
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

type NetworkSetCacheDisabledParams struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// SetCacheDisabledWithParams - Toggles ignoring cache for each request. If `true`, cache will not be used.
func (c *Network) SetCacheDisabledWithParams(v *NetworkSetCacheDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCacheDisabled", Params: v})
}

// SetCacheDisabled - Toggles ignoring cache for each request. If `true`, cache will not be used.
// cacheDisabled - Cache disabled state.
func (c *Network) SetCacheDisabled(cacheDisabled bool) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetCacheDisabledParams
	v.CacheDisabled = cacheDisabled
	return c.SetCacheDisabledWithParams(&v)
}

type NetworkSetCookieParams struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
	// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Url string `json:"url,omitempty"`
	// Cookie domain.
	Domain string `json:"domain,omitempty"`
	// Cookie path.
	Path string `json:"path,omitempty"`
	// True if cookie is secure.
	Secure bool `json:"secure,omitempty"`
	// True if cookie is http-only.
	HttpOnly bool `json:"httpOnly,omitempty"`
	// Cookie SameSite type. enum values: Strict, Lax
	SameSite string `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires float64 `json:"expires,omitempty"`
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
// name - Cookie name.
// value - Cookie value.
// url - The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
// domain - Cookie domain.
// path - Cookie path.
// secure - True if cookie is secure.
// httpOnly - True if cookie is http-only.
// sameSite - Cookie SameSite type. enum values: Strict, Lax
// expires - Cookie expiration date, session cookie if not set
// Returns -  success - True if successfully set cookie.
func (c *Network) SetCookie(name string, value string, url string, domain string, path string, secure bool, httpOnly bool, sameSite string, expires float64) (bool, error) {
	var v NetworkSetCookieParams
	v.Name = name
	v.Value = value
	v.Url = url
	v.Domain = domain
	v.Path = path
	v.Secure = secure
	v.HttpOnly = httpOnly
	v.SameSite = sameSite
	v.Expires = expires
	return c.SetCookieWithParams(&v)
}

type NetworkSetCookiesParams struct {
	// Cookies to be set.
	Cookies []*NetworkCookieParam `json:"cookies"`
}

// SetCookiesWithParams - Sets given cookies.
func (c *Network) SetCookiesWithParams(v *NetworkSetCookiesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setCookies", Params: v})
}

// SetCookies - Sets given cookies.
// cookies - Cookies to be set.
func (c *Network) SetCookies(cookies []*NetworkCookieParam) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetCookiesParams
	v.Cookies = cookies
	return c.SetCookiesWithParams(&v)
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

type NetworkSetRequestInterceptionParams struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
	Patterns []*NetworkRequestPattern `json:"patterns"`
}

// SetRequestInterceptionWithParams - Sets the requests to intercept that match a the provided patterns and optionally resource types.
func (c *Network) SetRequestInterceptionWithParams(v *NetworkSetRequestInterceptionParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setRequestInterception", Params: v})
}

// SetRequestInterception - Sets the requests to intercept that match a the provided patterns and optionally resource types.
// patterns - Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
func (c *Network) SetRequestInterception(patterns []*NetworkRequestPattern) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetRequestInterceptionParams
	v.Patterns = patterns
	return c.SetRequestInterceptionWithParams(&v)
}

type NetworkSetUserAgentOverrideParams struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
	// Browser langugage to emulate.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	// The platform navigator.platform should return.
	Platform string `json:"platform,omitempty"`
}

// SetUserAgentOverrideWithParams - Allows overriding user agent with the given string.
func (c *Network) SetUserAgentOverrideWithParams(v *NetworkSetUserAgentOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Network.setUserAgentOverride", Params: v})
}

// SetUserAgentOverride - Allows overriding user agent with the given string.
// userAgent - User agent to use.
// acceptLanguage - Browser langugage to emulate.
// platform - The platform navigator.platform should return.
func (c *Network) SetUserAgentOverride(userAgent string, acceptLanguage string, platform string) (*gcdmessage.ChromeResponse, error) {
	var v NetworkSetUserAgentOverrideParams
	v.UserAgent = userAgent
	v.AcceptLanguage = acceptLanguage
	v.Platform = platform
	return c.SetUserAgentOverrideWithParams(&v)
}
