// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Security functionality.
// API Version: 1.2

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// An explanation of an factor contributing to the security state.
type SecuritySecurityStateExplanation struct {
	SecurityState    string `json:"securityState"`    // Security state representing the severity of the factor being explained. enum values: unknown, neutral, insecure, warning, secure, info
	Summary          string `json:"summary"`          // Short phrase describing the type of factor.
	Description      string `json:"description"`      // Full text explanation of the factor.
	HasCertificate   bool   `json:"hasCertificate"`   // True if the page has a certificate.
	MixedContentType string `json:"mixedContentType"` // The type of mixed content described by the explanation. enum values: blockable, optionally-blockable, none
}

// Information about insecure content on the page.
type SecurityInsecureContentStatus struct {
	RanMixedContent                bool   `json:"ranMixedContent"`                // True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	DisplayedMixedContent          bool   `json:"displayedMixedContent"`          // True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	ContainedMixedForm             bool   `json:"containedMixedForm"`             // True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	RanContentWithCertErrors       bool   `json:"ranContentWithCertErrors"`       // True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool   `json:"displayedContentWithCertErrors"` // True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	RanInsecureContentStyle        string `json:"ranInsecureContentStyle"`        // Security state representing a page that ran insecure content. enum values: unknown, neutral, insecure, warning, secure, info
	DisplayedInsecureContentStyle  string `json:"displayedInsecureContentStyle"`  // Security state representing a page that displayed insecure content. enum values: unknown, neutral, insecure, warning, secure, info
}

// The security state of the page changed.
type SecuritySecurityStateChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		SecurityState         string                              `json:"securityState"`         // Security state. enum values: unknown, neutral, insecure, warning, secure, info
		SchemeIsCryptographic bool                                `json:"schemeIsCryptographic"` // True if the page was loaded over cryptographic transport such as HTTPS.
		Explanations          []*SecuritySecurityStateExplanation `json:"explanations"`          // List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
		InsecureContentStatus *SecurityInsecureContentStatus      `json:"insecureContentStatus"` // Information about insecure content on the page.
		Summary               string                              `json:"summary,omitempty"`     // Overrides user-visible description of the state.
	} `json:"Params,omitempty"`
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the handleCertificateError command. Note: this event does not fire if the certificate error has been allowed internally.
type SecurityCertificateErrorEvent struct {
	Method string `json:"method"`
	Params struct {
		EventId    int    `json:"eventId"`    // The ID of the event.
		ErrorType  string `json:"errorType"`  // The type of the error.
		RequestURL string `json:"requestURL"` // The url that was requested.
	} `json:"Params,omitempty"`
}

type Security struct {
	target gcdmessage.ChromeTargeter
}

func NewSecurity(target gcdmessage.ChromeTargeter) *Security {
	c := &Security{target: target}
	return c
}

// Enables tracking security state changes.
func (c *Security) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.enable"})
}

// Disables tracking security state changes.
func (c *Security) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.disable"})
}

// Displays native dialog with the certificate details.
func (c *Security) ShowCertificateViewer() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.showCertificateViewer"})
}

type SecurityHandleCertificateErrorParams struct {
	// The ID of the event.
	EventId int `json:"eventId"`
	// The action to take on the certificate error. enum values: continue, cancel
	Action string `json:"action"`
}

// HandleCertificateErrorWithParams - Handles a certificate error that fired a certificateError event.
func (c *Security) HandleCertificateErrorWithParams(v *SecurityHandleCertificateErrorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.handleCertificateError", Params: v})
}

// HandleCertificateError - Handles a certificate error that fired a certificateError event.
// eventId - The ID of the event.
// action - The action to take on the certificate error. enum values: continue, cancel
func (c *Security) HandleCertificateError(eventId int, action string) (*gcdmessage.ChromeResponse, error) {
	var v SecurityHandleCertificateErrorParams
	v.EventId = eventId
	v.Action = action
	return c.HandleCertificateErrorWithParams(&v)
}

type SecuritySetOverrideCertificateErrorsParams struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// SetOverrideCertificateErrorsWithParams - Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
func (c *Security) SetOverrideCertificateErrorsWithParams(v *SecuritySetOverrideCertificateErrorsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Security.setOverrideCertificateErrors", Params: v})
}

// SetOverrideCertificateErrors - Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
// override - If true, certificate errors will be overridden.
func (c *Security) SetOverrideCertificateErrors(override bool) (*gcdmessage.ChromeResponse, error) {
	var v SecuritySetOverrideCertificateErrorsParams
	v.Override = override
	return c.SetOverrideCertificateErrorsWithParams(&v)
}
